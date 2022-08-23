package repository

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap"
	"github.com/wakataw/moku/config"
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

type ldapRepository struct {
	ldapConfig  *config.Ldap
	ldapMapping *config.LdapAttributeMapping
	Conn        *ldap.Conn
	DB          *gorm.DB
}

var (
	LdapProfileNotFound = errors.New("ldap profile not found")
)

func (r *ldapRepository) Search(uid string) (*ldap.Entry, error) {
	searchRequest := ldap.SearchRequest{
		BaseDN:       r.ldapConfig.BaseDN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       fmt.Sprintf("(&(objectClass=user)(samaccountname=%v))", uid),
		Attributes:   []string{},
		Controls:     nil,
	}

	result, err := r.Conn.Search(&searchRequest)

	if err != nil {
		return nil, err
	}

	if len(result.Entries) != 1 {
		return nil, LdapProfileNotFound
	}

	return result.Entries[0], nil
}

func (r *ldapRepository) Authenticate(uid string, pwd string) (*entity.User, error) {
	entry, err := r.Search(uid)

	if err != nil {
		return &entity.User{}, err
	}

	err = r.Conn.Bind(entry.GetAttributeValue("dn"), pwd)

	if err != nil {
		return &entity.User{}, err
	}

	// get user from db
	var user entity.User
	user.AccountType = "ldap"

	for _, v := range entry.Attributes {
		switch v.Name {
		case r.ldapMapping.Username:
			user.Username = v.Values[0]
		case r.ldapMapping.Email:
			user.Email = v.Values[0]
		case r.ldapMapping.FullName:
			user.FullName = v.Values[0]
		case r.ldapMapping.Position:
			user.Position = v.Values[0]
		case r.ldapMapping.Department:
			user.Department = v.Values[0]
		case r.ldapMapping.Office:
			user.Office = v.Values[0]
		case r.ldapMapping.Title:
			user.Title = v.Values[0]
		case r.ldapMapping.IDNumber:
			user.IDNumber = v.Values[0]
		}
	}

	return &user, nil
}

func NewLdapRepository(ldapConfig *config.Ldap, ldapMapping *config.LdapAttributeMapping) (LdapRepository, error) {
	conn, err := ldap.Dial(ldapConfig.Network, ldapConfig.Host)

	if err != nil {
		return nil, err
	}

	err = conn.Bind(ldapConfig.BindDN, ldapConfig.BindPwd)

	if err != nil {
		return nil, err
	}

	return &ldapRepository{
		ldapConfig:  ldapConfig,
		ldapMapping: ldapMapping,
		Conn:        conn,
	}, err
}

package service

import (
	"errors"
	"fmt"
	"github.com/go-ldap/ldap"
	"github.com/golang-jwt/jwt"
	"github.com/wakataw/moku/config"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type authService struct {
	userRepo     repository.UserRepository
	tokenManager pkg.TokenManager
	ldapConfig   config.Ldap
	ldapMapping  config.LdapAttributeMapping
}

func (a *authService) LocalLogin(request model.LoginRequest) (*entity.User, error) {
	user, exist := a.userRepo.FindByUsername(request.Username)

	if user.AccountType == "ldap" {
		return &user, ErrNotLocalUser
	}

	if !exist {
		return &entity.User{}, ErrUserNotExists
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return &entity.User{}, ErrWrongUsernamePassword
	}

	return &user, nil
}

func (a *authService) LdapLogin(request model.LoginRequest) (*entity.User, error) {
	conn, err := ldap.Dial(a.ldapConfig.Network, a.ldapConfig.Host)

	if err != nil {
		return &entity.User{}, ErrLdapConnection
	}

	defer conn.Close()

	err = conn.Bind(a.ldapConfig.BindDN, a.ldapConfig.BindPwd)

	if err != nil {
		return &entity.User{}, ErrLdapBind
	}

	searchRequest := ldap.SearchRequest{
		BaseDN:       a.ldapConfig.BaseDN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       fmt.Sprintf("(&(objectClass=user)(samaccountname=%v))", request.Username),
		Attributes:   []string{},
		Controls:     nil,
	}

	result, err := conn.Search(&searchRequest)

	if err != nil || len(result.Entries) != 1 {
		return &entity.User{}, ErrLdapEmptyResult
	}

	err = conn.Bind(result.Entries[0].DN, request.Password)

	if err != nil {
		return &entity.User{}, ErrWrongUsernamePassword
	}

	user := &entity.User{
		AccountType: "ldap",
		Office:      "",
		Title:       "",
	}

	for _, v := range result.Entries[0].Attributes {
		switch v.Name {
		case a.ldapMapping.Username:
			user.Username = v.Values[0]
		case a.ldapMapping.Email:
			user.Email = v.Values[0]
		case a.ldapMapping.FullName:
			user.FullName = v.Values[0]
		case a.ldapMapping.Position:
			user.Position = v.Values[0]
		case a.ldapMapping.Department:
			user.Department = v.Values[0]
		case a.ldapMapping.Office:
			user.Office = v.Values[0]
		case a.ldapMapping.Title:
			user.Title = v.Values[0]
		}
	}

	err = a.userRepo.Insert(user)

	if err != nil {
		return &entity.User{}, err
	}

	return user, nil

}

func (a *authService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	// first try local login
	user, err := a.LocalLogin(request)

	if err == ErrNotLocalUser || err == ErrUserNotExists {
		err = nil
		user, err = a.LdapLogin(request)
	}

	if err != nil {
		log.Println(err.Error())
		return &model.LoginResponse{}, err
	}

	// construct role slice
	var roles []string
	if user.IsAdmin {
		roles = append(roles, "admin")
	}
	if user.IsManager {
		roles = append(roles, "manager")
	}
	if user.IsTeacher {
		roles = append(roles, "teacher")
	}

	token, err := a.tokenManager.GenerateToken(user.ID, roles)

	if err != nil {
		return &model.LoginResponse{}, errors.New("user does not exist")
	}

	response := model.LoginResponse{
		AcessToken:   token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return &response, nil

}

func (a *authService) RefreshToken(request model.RefreshTokenRequest) (*model.LoginResponse, error) {
	var response model.LoginResponse

	// vlaidate token
	token, err := a.tokenManager.Validate(request.RefreshToken)

	if err != nil {
		return &response, err
	}
	claims := token.Claims.(jwt.MapClaims)

	if claims["typ"] != "refresh" {
		return &response, errors.New("provided token is not a valid refresh token")
	}

	// TODO: Implement get roles from repository
	var roles []string

	userId, err := strconv.Atoi(claims["sub"].(string))

	if err != nil {
		return &response, err
	}

	if userId == 1 {
		roles = []string{"admin", "student"}
	} else {
		roles = []string{"student"}
	}

	newToken, err := a.tokenManager.GenerateToken(userId, roles)

	if err != nil {
		return &response, err
	}

	response.AcessToken = newToken.AccessToken
	response.RefreshToken = newToken.RefreshToken

	return &response, nil
}

func NewAuthService(userRepo *repository.UserRepository, tokenManager *pkg.TokenManager, ldapConfig *config.Ldap, ldapMapping *config.LdapAttributeMapping) AuthService {
	return &authService{
		userRepo:     *userRepo,
		tokenManager: *tokenManager,
		ldapConfig:   *ldapConfig,
		ldapMapping:  *ldapMapping,
	}
}

package repository

import "github.com/wakataw/moku/entity"

type LdapRepository interface {
	Authenticate(uid string, pwd string) (*entity.User, error)
}

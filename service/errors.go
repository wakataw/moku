package service

import "errors"

var (
	ErrNotLocalUser          = errors.New("account is not local user")
	ErrLdapConnection        = errors.New("cannot dial ldap server")
	ErrLdapBind              = errors.New("cannot bind to ldap")
	ErrLdapEmptyResult       = errors.New("user is not found in active directory")
	ErrWrongUsernamePassword = errors.New("wrong username/password")
	ErrUserNotExists         = errors.New("user not exists")
	ErrObjectDoesntExists    = errors.New("object doesnt exists")
)

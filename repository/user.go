package repository

import "github.com/wakataw/moku/entity"

type UserRepository interface {
	Insert(user *entity.User) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
	FindById(userId int) entity.User
	FindByUsername(username string) (entity.User, bool)
	FindRoles(userId int) []string
}

package repository

import "github.com/wakataw/moku/entity"

type UserRepository interface {
	Insert(user entity.User) error
	FindById(userId int) entity.User
}

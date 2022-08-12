package repository

import (
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

type userRepositoryImpl struct {
	DB *gorm.DB
}

func (u *userRepositoryImpl) Insert(user entity.User) error {
	result := u.DB.Create(&user)

	return result.Error
}

func (u *userRepositoryImpl) FindById(userId int) entity.User {
	var user entity.User

	u.DB.First(&user, userId)

	return user
}

package repository

import (
	"errors"
	"fmt"
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

func (u *userRepositoryImpl) SetRoles(user *entity.User, roles ...*entity.Role) (*entity.User, error) {
	err := u.DB.Model(&user).Association("Roles").Replace(roles)

	return user, err
}

func (u *userRepositoryImpl) All(lastCursor int, limit int, query string, ascending bool) (users *[]entity.User, err error) {
	tx := u.DB.Where("full_name like ?", fmt.Sprintf("%%%v%%", query))

	// pagination
	if lastCursor > 0 {
		if ascending {
			tx.Where("id > ?", lastCursor)
		} else {
			tx.Where("id < ?", lastCursor)
		}
	}

	// order
	if ascending {
		tx.Order("id asc")
	} else {
		tx.Order("id desc")
	}

	// add limit
	tx.Limit(limit)

	err = tx.Find(&users).Error

	if err != nil {
		return &[]entity.User{}, err
	}

	return users, nil
}

func (u *userRepositoryImpl) Update(user *entity.User) error {
	result := u.DB.Omit("users.created_at").Where("id = ?", user.ID).Updates(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

func (u *userRepositoryImpl) Delete(user *entity.User) error {
	// hard delete user since username is unique
	err := u.DB.Unscoped().Delete(&user).Error

	return err
}

func (u *userRepositoryImpl) FindRoles(userId int) []string {
	var user entity.User
	result := u.DB.Select("id").Preload("Roles").First(&user, userId)

	if result.RowsAffected != 1 {
		return []string{}
	}

	var roles []string

	for _, v := range user.Roles {
		roles = append(roles, v.Name)
	}

	return roles
}

func (u *userRepositoryImpl) FindByUsername(username string) (user entity.User, exists bool) {
	result := u.DB.Where("username = ?", username).Preload("Roles").First(&user)

	return user, result.RowsAffected == 1
}

func (u *userRepositoryImpl) Insert(user *entity.User) (*entity.User, error) {
	result := u.DB.Where("username = ?", user.Username).Find(&user)

	if result.RowsAffected > 0 {
		return user, errors.New("user already exists")
	}

	result = u.DB.Create(&user)

	return user, result.Error
}

func (u *userRepositoryImpl) FindById(userId int) entity.User {
	var user entity.User

	u.DB.Preload("Roles").First(&user, userId)

	return user
}

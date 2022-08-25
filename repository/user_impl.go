package repository

import (
	"fmt"
	"github.com/wakataw/moku/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

type userRepositoryImpl struct {
	DB *gorm.DB
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
	result := u.DB.Save(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userRepositoryImpl) Delete(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepositoryImpl) FindRoles(userId int) []string {
	var user entity.User
	result := u.DB.Select("is_admin", "is_manager", "is_teacher").First(&user, userId)

	if result.RowsAffected != 1 {
		return []string{}
	}

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

	return roles
}

func (u *userRepositoryImpl) FindByUsername(username string) (user entity.User, exists bool) {
	result := u.DB.Where("username = ?", username).First(&user)

	return user, result.RowsAffected == 1
}

func (u *userRepositoryImpl) Insert(user *entity.User) error {
	result := u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.AssignmentColumns([]string{"position", "department", "office", "title"}),
	}).Create(&user)

	return result.Error
}

func (u *userRepositoryImpl) FindById(userId int) entity.User {
	var user entity.User

	u.DB.First(&user, userId)

	return user
}

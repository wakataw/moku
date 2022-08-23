package repository

import (
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

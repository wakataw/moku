package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string `gorm:"type:varchar(32);index;unique"`
	Password    string `gorm:"type:varchar(255)"`
	AccountType string `gorm:"type:varchar(16);default:'local'"`
	Email       string `gorm:";index;unique"`
	IDNumber    string `gorm:"type:varchar(18);unique"`
	FullName    string `gorm:"index"`
	Position    string `gorm:"index"`
	Department  string
	Office      string
	Title       string
	IsActive    bool
	IsAdmin     bool
	IsTeacher   bool
	IsManager   bool
}

type UserRoles struct {
	ID        int
	IsAdmin   bool
	IsTeacher bool
	IsManager bool
}

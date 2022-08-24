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
	Roles       []Role `gorm:"many2many:users_roles"`
}

type Role struct {
	*gorm.Model
	ID          int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"unique;type:varchar(32)"`
	Permissions []Permission `gorm:"many2many:roles_permissions"`
}

type Permission struct {
	*gorm.Model
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"unique;type:varchar(100)"`
	Roles []Role `gorm:"many2many:roles_permissions"`
}

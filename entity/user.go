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
	Section     string
	Division    string
	Office      string
	IsActive    bool
	IsAdmin     bool
	Roles       []Role `gorm:"many2many:users_roles;"`
}

type Role struct {
	gorm.Model
	ID          uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string        `gorm:"type:varchar(120)"`
	Users       []*User       `gorm:"many2many:users_roles"`
	Permissions []*Permission `gorm:"many2many:roles_permissions"`
}

type Permission struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Object string
	Create bool
	Read   bool
	Update bool
	Delete bool
	Roles  []Role `gorm:"many2many:roles_permissions"`
}

package entity

import (
	"gorm.io/gorm"
	"time"
)

type Program struct {
	*gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(255);index" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Slug        string    `gorm:"type:varchar(128);index;unique" json:"slug"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Show        bool      `json:"show"`
	Public      bool      `json:"public"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	DeletedBy   int       `json:"deleted_by"`
}

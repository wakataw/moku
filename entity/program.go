package entity

import (
	"gorm.io/gorm"
)

type Program struct {
	*gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(255);index" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Slug        string `gorm:"type:varchar(128);index;unique;<-:create" json:"slug"`
	Start       *int   `json:"start"`
	End         *int   `json:"end"`
	Show        bool   `json:"show"`
	Public      bool   `json:"public"`
	CreatedBy   *int   `json:"created_by" gorm:"<-:create"`
	CreatedUser User   `gorm:"foreignKey:CreatedBy"`
	UpdatedBy   *int   `json:"updated_by"`
	UpdatedUser User   `gorm:"foreignKey:UpdatedBy"`
}

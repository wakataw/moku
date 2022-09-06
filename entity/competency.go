package entity

import "gorm.io/gorm"

type Competency struct {
	gorm.Model
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       int    `gorm:"index" json:"name"`
	LegalBasis string `gorm:"index" json:"legal_basis"`
	Expired    bool   `json:"expired"`
}

package entity

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int `json:"user_id"`
	User      User
	ProgramID int `json:"program_id"`
	Program   Program
}

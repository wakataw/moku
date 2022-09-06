package entity

type Enrollment struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	ProgramID int `json:"program_id"`
	Program   Program
}

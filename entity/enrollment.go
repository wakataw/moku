package entity

type Enrollment struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int `json:"user_id"`
	User      User
	ProgramID int `json:"program_id"`
	Program   Program
}

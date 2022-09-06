package entity

import (
	"gorm.io/gorm"
	"time"
)

type Quiz struct {
	gorm.Model
	ID             int           `gorm:"primaryKey;autoIncrement"`
	Name           string        `gorm:"varchar(100);index" json:"name"`
	Description    string        `gorm:"text" json:"description"`
	Slug           string        `gorm:"varchar(128);unique;index" json:"slug"`
	Start          time.Time     `json:"start"`
	End            time.Time     `json:"end"`
	Duration       time.Duration `json:"duration"`
	TotalQuestions int           `json:"total_questions"`
	MaxAttempts    int           `json:"max_attempts"`
	IsAdaptive     bool          `json:"is_adaptive"`
	ShowScore      bool          `json:"show_score"`
	ProgramID      int           `json:"program_id"`
	Program        Program
}

type QuestionPackage struct {
	gorm.Model
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(100)" json:"name"`
	QuizID int    `json:"quiz_id"`
	Quiz   Quiz
}

type Question struct {
	gorm.Model
	ID           int `gorm:"primaryKey;autoIncrement" json:"id"`
	Level        int `json:"level"`
	CompetencyID int `json:"competency_id"`
	Competency   Competency
	Text         string `gorm:"type:text" json:"text"`
	Type         int    `json:"type"`
	Choices      []QuestionMultipleChoices
}

type QuestionMultipleChoices struct {
	gorm.Model
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Choice     string `gorm:"type:text" json:"choice"`
	Score      int    `json:"score"`
	Question   Question
	QuestionID int `json:"question_id"`
}

type Attempt struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	User   User
	UserID int `json:"user_id"`
}

type Grade struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:"id"`
	AttemptID int `json:"attempt_id"`
	Attempt   Attempt
	Grades    float64 `json:"grades"`
}

type Submission struct {
	ID         int `gorm:"primaryKey;autoIncrement" json:"id"`
	AttemptID  int `json:"attempt_id"`
	Attempt    Attempt
	QuestionID int `json:"question_id"`
	Question   Question
	AnswerID   int                     `json:"answer_id"`
	Answer     QuestionMultipleChoices `gorm:"foreignKey:AnswerID"`
	AnswerText string                  `gorm:"type:text" json:"answer_text"`
	Status     bool                    `json:"status"`
}

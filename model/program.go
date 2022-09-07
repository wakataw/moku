package model

import "time"

type CreateProgramRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Show        bool      `json:"show"`
	Public      bool      `json:"public"`
}

type GetProgramResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Show        bool      `json:"show"`
	Public      bool      `json:"public"`
	CreatedBy   int       `json:"created_by"`
}

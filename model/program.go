package model

type CreateProgramRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Start       *int   `json:"start" time_format:"2006-01-02 15:04"`
	End         *int   `json:"end" time_format:"2006-01-02 15:04"`
	Show        bool   `json:"show"`
	Public      bool   `json:"public"`
	CreatedBy   int    `json:"-"`
	UpdatedBy   int    `json:"-"`
}

type GetProgramResponse struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Start       *int                   `json:"start"`
	End         *int                   `json:"end"`
	Show        bool                   `json:"show"`
	Public      bool                   `json:"public"`
	CreatedUser *GetUserResponseSimple `json:"created_by,omitempty"`
	UpdatedUser *GetUserResponseSimple `json:"updated_by,omitempty"`
}

type UpdateProgramRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start       *int   `json:"start"`
	End         *int   `json:"end"`
	Show        bool   `json:"show"`
	Public      bool   `json:"public"`
	UpdatedBy   int    `json:"-"`
}

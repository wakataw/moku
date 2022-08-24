package model

type RequestParameter struct {
	LastCursor *int   `json:"last_cursor" form:"last_cursor" binding:"required"`
	Limit      int    `json:"limit" binding:"required"`
	Query      string `json:"query" binding:"required"`
	Ascending  bool   `json:"ascending" binding:"required"`
}

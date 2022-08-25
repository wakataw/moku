package model

type RequestParameter struct {
	LastCursor *int   `json:"last_cursor" form:"last_cursor" binding:"required"`
	Limit      int    `json:"limit" form:"limit" binding:"required"`
	Query      string `json:"query" form:"query"`
	Ascending  bool   `json:"ascending" form:"ascending"`
}

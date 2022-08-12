package model

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IDNumber string `json:"id_number"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
	Section  string `json:"section"`
	Division string `json:"division"`
	Office   string `json:"office"`
}

type CreateUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IDNumber string `json:"id_number"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
	Section  string `json:"section"`
	Division string `json:"division"`
	Office   string `json:"office"`
}

type GetUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IDNumber string `json:"id_number"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
	Section  string `json:"section"`
	Division string `json:"division"`
	Office   string `json:"office"`
}

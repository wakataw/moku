package model

type CreateUserRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IDNumber   string `json:"id_number"`
	FullName   string `json:"full_name"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Office     string `json:"office"`
	Title      string `json:"title"`
	IsAdmin    bool   `json:"is_admin"`
	IsTeacher  bool   `json:"is_teacher"`
	IsManager  bool   `json:"is_manager"`
}

type CreateUserResponse struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	IDNumber   string `json:"id_number"`
	FullName   string `json:"full_name"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Office     string `json:"office"`
	Title      string `json:"title"`
	IsAdmin    bool   `json:"is_admin"`
	IsTeacher  bool   `json:"is_teacher"`
	IsManager  bool   `json:"is_manager"`
}

type GetUserResponse struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	IDNumber   string `json:"id_number"`
	FullName   string `json:"full_name"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Office     string `json:"office"`
	Title      string `json:"title"`
	IsAdmin    bool   `json:"is_admin"`
	IsTeacher  bool   `json:"is_teacher"`
	IsManager  bool   `json:"is_manager"`
}

type SetUserRoleRequest struct {
	UserId int                     `json:"user_id"`
	Roles  []GetRoleSimpleResponse `json:"roles"`
}

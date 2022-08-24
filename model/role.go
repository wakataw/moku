package model

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type GetRoleByIdRequest struct {
	ID int `json:"id"`
}

type GetRoleByNameRequest struct {
	Name string `json:"name"`
}

type GetRoleResponse struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Permissions []GetPermissionResponse `json:"permissions"`
}

type GetRoleSimpleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateRoleRequest struct {
	ID          int                       `json:"id"`
	Name        string                    `json:"name"`
	Permissions []UpdatePermissionRequest `json:"permissions"`
}

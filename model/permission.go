package model

type GetPermissionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdatePermissionRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

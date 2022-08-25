package service

import "github.com/wakataw/moku/model"

type PermissionService interface {
	All(request *model.RequestParameter) (responses *[]model.GetPermissionResponse, err error)
	Create(request *model.CreatePermissionRequest) (response *model.GetPermissionResponse, err error)
	Update(request *model.UpdatePermissionRequest) (response *model.GetPermissionResponse, err error)
	Delete(permId int) (err error)
	GetPermissionById(permId int) *model.GetPermissionResponse
	GetPermissionByName(permId string) (*model.GetPermissionResponse, error)
}

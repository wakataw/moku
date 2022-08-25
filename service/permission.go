package service

import "github.com/wakataw/moku/model"

type PermissionService interface {
	All(request *model.RequestParameter) (responses *[]model.GetPermissionResponse, err error)
	Create(request *model.CreatePermissionRequest) (response *model.GetPermissionResponse, err error)
}

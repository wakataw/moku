package service

import "github.com/wakataw/moku/model"

type PermissionService interface {
	Create(request *model.CreatePermissionRequest) (response *model.GetPermissionResponse, err error)
}

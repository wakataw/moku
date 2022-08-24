package service

import "github.com/wakataw/moku/model"

type RoleService interface {
	Create(request *model.CreateRoleRequest) (response *model.GetRoleResponse, err error)
	Update(request *model.UpdateRoleRequest) (response *model.GetRoleResponse, err error)
	Delete(roleId int) (err error)
	GetRoleById(roleId int) *model.GetRoleResponse
	GetRoleByName(roleName string) (*model.GetRoleResponse, error)
	All(request *model.RequestParameter) (responses *[]model.GetRoleSimpleResponse, err error)
}

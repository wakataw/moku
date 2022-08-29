package service

import "github.com/wakataw/moku/model"

type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	GetById(userId int) (response model.GetUserResponse, exists bool)
	CreateAdmin(request *model.CreateUserRequest) error
	Update(request *model.UpdateUserRequest) (response *model.CreateUserResponse, err error)
	Delete(userId int) error
	All(request *model.RequestParameter) (responses *[]model.GetUserResponse, err error)
	SetRole(request *model.SetUserRoleRequest) error
}

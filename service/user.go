package service

import "github.com/wakataw/moku/model"

type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	GetById(userId int) (response model.GetUserResponse, exists bool)
	CreateAdmin(request *model.CreateUserRequest) error
}

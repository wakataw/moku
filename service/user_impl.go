package service

import (
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Respository repository.UserRepository
}

func (u *userService) Create(request model.CreateUserRequest) (response model.CreateUserResponse, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.CreateUserResponse{}, err
	}

	user := entity.User{
		Username:    request.Username,
		Password:    string(passwordHash),
		AccountType: "local",
		Email:       request.Email,
		IDNumber:    request.IDNumber,
		FullName:    request.FullName,
		Position:    request.Position,
		Section:     request.Section,
		Office:      request.Office,
		Title:       request.Title,
	}

	err = u.Respository.Insert(user)

	if err != nil {
		return model.CreateUserResponse{}, err
	}

	response = model.CreateUserResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IDNumber: user.IDNumber,
		FullName: user.FullName,
		Position: user.Position,
		Section:  user.Section,
		Office:   user.Office,
		Title:    user.Title,
	}

	return response, nil

}

func (u *userService) GetById(userId int) (response model.GetUserResponse, exists bool) {
	user := u.Respository.FindById(userId)

	response = model.GetUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		IDNumber:  user.IDNumber,
		FullName:  user.FullName,
		Position:  user.Position,
		Section:   user.Section,
		Office:    user.Office,
		Title:     user.Title,
		IsManager: user.IsManager,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}
	return response, user.Model != nil
}

func (u *userService) CreateAdmin(request *model.CreateUserRequest) (err error) {
	request.FullName = "Administrator User"
	_, err = u.Create(*request)

	return
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userService{
		Respository: *userRepository,
	}
}

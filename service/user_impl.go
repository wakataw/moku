package service

import (
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repository.UserRepository
}

func (u *userService) All(request *model.RequestParameter) (responses *[]model.GetUserResponse, err error) {
	results, err := u.repository.All(
		*request.LastCursor,
		request.Limit,
		request.Query,
		request.Ascending,
	)

	if err != nil {
		return &[]model.GetUserResponse{}, err
	}

	var usersResp []model.GetUserResponse

	for _, v := range *results {
		usersResp = append(usersResp, model.GetUserResponse{
			Id:         v.ID,
			Username:   v.Username,
			Email:      v.Email,
			IDNumber:   v.IDNumber,
			FullName:   v.FullName,
			Position:   v.Position,
			Department: v.Department,
			Office:     v.Office,
			Title:      v.Title,
			IsAdmin:    v.IsAdmin,
			IsTeacher:  v.IsTeacher,
			IsManager:  v.IsManager,
		})
	}

	return &usersResp, nil
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
		Department:  request.Department,
		Office:      request.Office,
		Title:       request.Title,
	}

	err = u.repository.Insert(&user)

	if err != nil {
		return model.CreateUserResponse{}, err
	}

	response = model.CreateUserResponse{
		Id:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		IDNumber:   user.IDNumber,
		FullName:   user.FullName,
		Position:   user.Position,
		Department: user.Department,
		Office:     user.Office,
		Title:      user.Title,
	}

	return response, nil

}

func (u *userService) GetById(userId int) (response model.GetUserResponse, exists bool) {
	user := u.repository.FindById(userId)

	response = model.GetUserResponse{
		Id:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		IDNumber:   user.IDNumber,
		FullName:   user.FullName,
		Position:   user.Position,
		Department: user.Department,
		Office:     user.Office,
		Title:      user.Title,
		IsManager:  user.IsManager,
		IsAdmin:    user.IsAdmin,
		IsTeacher:  user.IsTeacher,
	}
	return response, user.Model != nil
}

func (u *userService) CreateAdmin(request *model.CreateUserRequest) (err error) {

	_, exists := u.repository.FindByUsername("admin")

	if exists {
		return nil
	}

	request.FullName = "Administrator User"
	_, err = u.Create(*request)

	return
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userService{
		repository: *userRepository,
	}
}

package service

import (
	"errors"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type authService struct {
	userRepo     repository.UserRepository
	tokenManager pkg.TokenManager
}

func (a *authService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	log.Println(request)
	user, exist := a.userRepo.FindByUsername(request.Username)
	log.Println(user)

	if !exist {
		return &model.LoginResponse{}, errors.New("User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return &model.LoginResponse{}, errors.New("wrong username or password")
	}

	// TODO: Implement get roles from repository
	var roles []string
	if user.ID == 1 {
		roles = []string{"admin", "student"}
	} else {
		roles = []string{"student"}
	}

	token, err := a.tokenManager.GenerateToken(user.ID, roles)

	if err != nil {
		return &model.LoginResponse{}, errors.New("User does not exist")
	}

	response := model.LoginResponse{
		AcessToken:   token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return &response, nil

}

func (a *authService) RefreshToken(request model.RefreshTokenRequest) (*model.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(userRepo *repository.UserRepository, tokenManager *pkg.TokenManager) AuthService {
	return &authService{
		userRepo:     *userRepo,
		tokenManager: *tokenManager,
	}
}

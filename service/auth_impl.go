package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type authService struct {
	userRepo     repository.UserRepository
	tokenManager pkg.TokenManager
}

func (a *authService) LocalLogin(request model.LoginRequest) (*entity.User, error) {
	user, exist := a.userRepo.FindByUsername(request.Username)

	if !exist {
		return &entity.User{}, errors.New("User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return &entity.User{}, errors.New("wrong username or password")
	}

	return &user, nil
}

func (a *authService) LdapLogin(request model.LoginRequest) (*model.LoginResponse, error) {
	return &model.LoginResponse{}, nil
}

func (a *authService) Login(request model.LoginRequest) (*model.LoginResponse, error) {
	// first try local login
	user, err := a.LocalLogin(request)

	if err != nil {
		return &model.LoginResponse{}, err
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
	var response model.LoginResponse

	// vlaidate token
	token, err := a.tokenManager.Validate(request.RefreshToken)

	if err != nil {
		return &response, err
	}
	claims := token.Claims.(jwt.MapClaims)

	if claims["typ"] != "refresh" {
		return &response, errors.New("provided token is not a valid refresh token")
	}

	// TODO: Implement get roles from repository
	var roles []string

	userId, err := strconv.Atoi(claims["sub"].(string))

	if err != nil {
		return &response, err
	}

	if userId == 1 {
		roles = []string{"admin", "student"}
	} else {
		roles = []string{"student"}
	}

	newToken, err := a.tokenManager.GenerateToken(userId, roles)

	if err != nil {
		return &response, err
	}

	response.AcessToken = newToken.AccessToken
	response.RefreshToken = newToken.RefreshToken

	return &response, nil
}

func NewAuthService(userRepo *repository.UserRepository, tokenManager *pkg.TokenManager) AuthService {
	return &authService{
		userRepo:     *userRepo,
		tokenManager: *tokenManager,
	}
}

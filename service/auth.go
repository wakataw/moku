package service

import "github.com/wakataw/moku/model"

type AuthService interface {
	Login(request model.LoginRequest) (*model.LoginResponse, error)
	RefreshToken(request model.RefreshTokenRequest) (*model.LoginResponse, error)
}

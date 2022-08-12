package model

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

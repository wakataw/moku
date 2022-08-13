package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/service"
	"net/http"
)

type authController struct {
	Service service.AuthService
}

func NewAuthController(service *service.AuthService) *authController {
	return &authController{
		Service: *service,
	}
}

func (ctl *authController) Route(r *gin.RouterGroup) {
	r.POST("/login", ctl.Login)
	r.POST("/refresh_token", ctl.RefreshToken)
}

func (ctl *authController) Login(c *gin.Context) {
	var request model.LoginRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ctl.Service.Login(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (ctl *authController) RefreshToken(c *gin.Context) {
	var request model.RefreshTokenRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ctl.Service.RefreshToken(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

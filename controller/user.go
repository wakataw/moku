package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/service"
	"net/http"
	"strconv"
)

type userController struct {
	service      service.UserService
	TokenManager pkg.TokenManager
}

func NewUserController(userService *service.UserService) *userController {
	return &userController{
		service: *userService,
	}
}

func (ctl *userController) Route(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/", ctl.Create)
		users.GET("/:id", ctl.GetById)
		users.GET("/", ctl.All)
		users.DELETE("/:id", ctl.Delete)
		users.POST("/:id/roles", ctl.SetRoles)
		users.PUT("/:id", ctl.Update)
	}
}

func (ctl *userController) Create(c *gin.Context) {
	var request model.CreateUserRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user object is not valid",
			"error":   err.Error(),
		})
		return
	}

	response, err := ctl.service.Create(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
		"data":    response,
	})

}

func (ctl *userController) GetById(c *gin.Context) {
	var statusCode int
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id is not valid",
		})
		return
	}

	user, exists := ctl.service.GetById(userId)

	if !exists {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, gin.H{
		"data": user,
	})

}

func (ctl *userController) All(c *gin.Context) {
	var request model.RequestParameter

	if err := c.BindQuery(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := ctl.service.All(&request)

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

func (ctl *userController) Delete(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ctl.service.Delete(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (ctl *userController) SetRoles(c *gin.Context) {
	var requests model.SetUserRoleRequest

	if err := c.BindJSON(&requests); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := ctl.service.SetRole(&requests)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (ctl *userController) Update(c *gin.Context) {
	var request model.UpdateUserRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := ctl.service.Update(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    response,
	})
}

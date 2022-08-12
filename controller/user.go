package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/service"
	"net/http"
	"strconv"
)

type userController struct {
	Service service.UserService
}

func NewUserController(userService *service.UserService) *userController {
	return &userController{
		Service: *userService,
	}
}

func (ctl *userController) Route(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/", ctl.Create)
		users.GET("/:id", ctl.GetById)
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

	response, err := ctl.Service.Create(request)

	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)

		switch mysqlErr.Number {
		case 1062:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Combination of username/email/ID Number already exists",
			})
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "General database error",
			})
		}

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

	user, exists := ctl.Service.GetById(userId)

	if !exists {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, gin.H{
		"data": user,
	})

}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type roleController struct {
	service service.RoleService
}

func NewRoleController(roleService *service.RoleService) *roleController {
	return &roleController{
		service: *roleService,
	}
}

func (ctl *roleController) Route(r *gin.RouterGroup) {
	roles := r.Group("/roles")
	{
		roles.POST("/", ctl.Create)
		roles.PUT("/", ctl.Update)
		roles.DELETE("/:id", ctl.Delete)
		roles.GET("/", ctl.GetAll)
		roles.GET("/:name", ctl.GetByName)
	}
}

func (ctl *roleController) Create(c *gin.Context) {
	var request model.CreateRoleRequest
	err := c.BindJSON(&request)

	// sanitize white space
	request.Name = strings.TrimSpace(request.Name)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ctl.service.Create(&request)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (ctl *roleController) Update(c *gin.Context) {
	var request model.UpdateRoleRequest
	err := c.BindJSON(&request)

	// sanitize white space
	request.Name = strings.TrimSpace(request.Name)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ctl.service.Update(&request)

	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("role with name `%v` already exists", request.Name),
			})
		default:
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "general server error",
			})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (ctl *roleController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "role id is not valid",
		})
		return
	}

	err = ctl.service.Delete(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (ctl *roleController) GetAll(c *gin.Context) {
	var request model.RequestParameter

	if err := c.BindQuery(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := ctl.service.All(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (ctl *roleController) GetByName(c *gin.Context) {
	name := c.Param("name")

	response, err := ctl.service.GetRoleByName(name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    response,
		"message": "success",
	})
}

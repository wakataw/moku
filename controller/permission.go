package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/service"
	"net/http"
)

type permissionController struct {
	service service.PermissionService
}

func NewPermissionController(permissionService *service.PermissionService) *permissionController {
	return &permissionController{service: *permissionService}
}

func (ctl *permissionController) Route(r *gin.RouterGroup) {
	permissions := r.Group("/permissions")
	{
		permissions.POST("/", ctl.Create)
		permissions.GET("/", ctl.All)
	}
}

func (ctl *permissionController) Create(c *gin.Context) {
	var request model.CreatePermissionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	response, err := ctl.service.Create(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    response,
	})
}

func (ctl *permissionController) All(c *gin.Context) {
	var request model.RequestParameter

	if err := c.BindQuery(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := ctl.service.All(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}

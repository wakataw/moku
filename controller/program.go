package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/model"
	"github.com/wakataw/moku/service"
	"net/http"
	"strconv"
)

type programController struct {
	service service.ProgramService
}

func NewProgramController(service *service.ProgramService) *programController {
	return &programController{service: *service}
}

func (ctl *programController) Route(r *gin.RouterGroup) {
	program := r.Group("/programs")
	{
		program.POST("/", ctl.Create)
		program.GET("/", ctl.All)
		program.GET("/:id", ctl.GetById)
		program.DELETE("/:id", ctl.Delete)
		program.PUT("/:id", ctl.Update)
	}
}

func (ctl *programController) Create(c *gin.Context) {
	var request model.CreateProgramRequest

	userId := c.GetInt("userId")
	request.CreatedBy = userId

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	response, err := ctl.service.Create(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data":    response,
	})
}

func (ctl *programController) All(c *gin.Context) {
	var request model.RequestParameter

	if err := c.BindQuery(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	response, err := ctl.service.All(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data":    response,
	})

}

func (ctl *programController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	response, err := ctl.service.GetProgramById(id)

	if err != nil {
		if errors.Is(err, service.ErrObjectDoesntExists) {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  false,
				"data":    nil,
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data":    response,
	})
}

func (ctl *programController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	err = ctl.service.Delete(id)

	if err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  false,
				"data":    nil,
			})
			return
		}
	}

	c.Status(http.StatusOK)
}

func (ctl *programController) Update(c *gin.Context) {
	var request model.UpdateProgramRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	response, err := ctl.service.Update(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data":    response,
	})

}

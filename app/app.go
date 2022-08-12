package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/config"
	"github.com/wakataw/moku/controller"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/middleware"
	"github.com/wakataw/moku/pkg"
	"github.com/wakataw/moku/repository"
	"github.com/wakataw/moku/service"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Permission{})

	return err
}

func Run(configDir string) {
	/*
		Configuration
	*/
	cfg, err := config.NewConfig(configDir)

	if err != nil {
		log.Fatal(err)
		return
	}

	/*
		Database
	*/
	// init database
	db := config.NewDB(&cfg.Mysql)
	// migrate
	err = migrate(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	/*
		initiate token manager
	*/

	tokenManager := pkg.TokenManager{
		Secret:     cfg.Auth.Secret,
		AccessTTL:  cfg.Auth.AccessTokenTTL,
		RefreshTTL: cfg.Auth.RefreshTokenTTL,
	}

	/*
		user repo and service
	*/
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	/*
		auth object
	*/
	authService := service.NewAuthService(&userRepository, &tokenManager)
	authController := controller.NewAuthController(&authService)

	/*
		Gin Server
	*/
	r := gin.Default()

	// index
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})

	/*
		auth router
	*/
	auth := r.Group("/auth")
	authController.Route(auth)

	/*
		api router
	*/
	api := r.Group("/api")
	api.Use(middleware.AuthRequiredMiddleware(&tokenManager), middleware.AdminRequiredMiddleware())
	{
		userController.Route(api)
	}

	r.Run(":8088")
}

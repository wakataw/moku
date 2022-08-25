package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wakataw/moku/config"
	"github.com/wakataw/moku/controller"
	"github.com/wakataw/moku/entity"
	"github.com/wakataw/moku/middleware"
	"github.com/wakataw/moku/model"
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

func createAdmin(userService service.UserService, adminCfg *config.DefaultAdmin) error {
	userRequest := &model.CreateUserRequest{
		Username:   adminCfg.Username,
		Password:   adminCfg.Password,
		Email:      adminCfg.Email,
		IDNumber:   "0",
		FullName:   "Administrator User",
		Position:   "-",
		Department: "-",
		Office:     "-",
		Title:      "-",
		IsAdmin:    true,
		IsTeacher:  true,
		IsManager:  true,
	}
	err := userService.CreateAdmin(userRequest)

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
		initiate JWT token manager
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
		role repo and service
	*/
	roleRepository := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(&roleRepository)
	roleController := controller.NewRoleController(&roleService)

	/*
		permission repo and service
	*/
	permissionRepository := repository.NewPermissionRepository(db)
	permissionService := service.NewPermissionService(&permissionRepository)
	permissionController := controller.NewPermissionController(&permissionService)

	// generate admin
	err = createAdmin(userService, &cfg.DefaultAdmin)

	if err != nil {
		log.Println(err.Error())
	}
	/*
		auth object
	*/
	authService := service.NewAuthService(&userRepository, &tokenManager, &cfg.Ldap, &cfg.LdapAttributeMapping)
	authController := controller.NewAuthController(&authService)

	/*
		Gin Server
	*/

	gin.SetMode(cfg.Mode)

	r := gin.Default()

	// index
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})

	/*
		V1 router
	*/
	v1 := r.Group("/v1")
	{

		auth := v1.Group("/")
		{
			authController.Route(auth)
		}

		adminRoute := v1.Group("/")
		adminRoute.Use(middleware.AuthRequiredMiddleware(&tokenManager), middleware.AdminRequiredMiddleware())
		{
			userController.Route(adminRoute)
			roleController.Route(adminRoute)
			permissionController.Route(adminRoute)
		}
	}
	r.Run(":8088")
}

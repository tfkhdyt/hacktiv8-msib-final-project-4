package handler

import (
	"hacktiv8-msib-final-project-4/database"
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-4/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-4/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	db := database.GetPostgresInstance()

	userRepo := userpg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := httphandler.NewUserHandler(userService)

	categoryRepo := categorypg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := httphandler.NewCategoryHandler(categoryService)

	authService := service.NewAuthService(userRepo)

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PATCH("/users/topup", authService.Authentication(), userHandler.TopUp)

	r.POST("/categories", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.CreateCategory)

	log.Fatalln(r.Run(":" + port))
}

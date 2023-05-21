package handler

import (
	"hacktiv8-msib-final-project-4/injector"
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

	userHandler := injector.InitializeUserHandler()
	categoryHandler := injector.InitializeCategoryHandler()
	authService := injector.InitializeAuthService()

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PATCH("/users/topup", authService.Authentication(), userHandler.TopUp)

	r.Use(authService.Authentication(), authService.AdminAuthorization())
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories", categoryHandler.GetAllCategories)
	r.PATCH("/categories/:categoryID", categoryHandler.UpdateCategory)

	log.Fatalln(r.Run(":" + port))
}

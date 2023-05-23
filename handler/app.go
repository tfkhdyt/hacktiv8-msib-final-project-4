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
	productHandler := injector.InitializeProductHandler()
	authService := injector.InitializeAuthService()

	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PATCH("/users/topup", authService.Authentication(), userHandler.TopUp)

	r.POST("/categories", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.CreateCategory)
	r.GET("/categories", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.GetAllCategories)
	r.PATCH("/categories/:categoryID", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.UpdateCategory)
	r.DELETE("/categories/:categoryID", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.DeleteCategory)

	r.POST("/products", authService.Authentication(), authService.AdminAuthorization(), productHandler.CreateProduct)
	r.GET("/products", authService.Authentication(), productHandler.GetAllProducts)
	r.PUT("/products/:productID", authService.Authentication(), authService.AdminAuthorization(), productHandler.UpdateProduct)

	log.Fatalln(r.Run(":" + port))
}

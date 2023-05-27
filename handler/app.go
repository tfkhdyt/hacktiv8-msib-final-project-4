package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/injector"
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
	transactionHandler := injector.InitializeTransactionHistoryHandler()
	authService := injector.InitializeAuthService()

	// users routes
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.PATCH("/users/topup", authService.Authentication(), userHandler.TopUp)

	// categories routes
	r.POST(
		"/categories",
		authService.Authentication(),
		authService.AdminAuthorization(),
		categoryHandler.CreateCategory,
	)
	r.GET(
		"/categories",
		authService.Authentication(),
		authService.AdminAuthorization(),
		categoryHandler.GetAllCategories,
	)
	r.PATCH(
		"/categories/:categoryID",
		authService.Authentication(),
		authService.AdminAuthorization(),
		categoryHandler.UpdateCategory,
	)
	r.DELETE(
		"/categories/:categoryID",
		authService.Authentication(),
		authService.AdminAuthorization(),
		categoryHandler.DeleteCategory,
	)

	// products routes
	r.POST(
		"/products",
		authService.Authentication(),
		authService.AdminAuthorization(),
		productHandler.CreateProduct,
	)
	r.GET("/products", authService.Authentication(), productHandler.GetAllProducts)
	r.PUT(
		"/products/:productID",
		authService.Authentication(),
		authService.AdminAuthorization(),
		productHandler.UpdateProduct,
	)
	r.DELETE(
		"/products/:productID",
		authService.Authentication(),
		authService.AdminAuthorization(),
		productHandler.DeleteProduct,
	)

	// transaction histories routes
	r.POST("/transactions", authService.Authentication(), transactionHandler.CreateTransaction)
	r.GET(
		"/transactions/my-transactions",
		authService.Authentication(),
		transactionHandler.GetTransactionsByUserID,
	)
	r.GET(
		"/transactions/all-transactions",
		authService.Authentication(),
		authService.AdminAuthorization(),
		transactionHandler.GetAllTransactions,
	)

	log.Fatalln(r.Run(":" + port))
}

package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/injector"
	"hacktiv8-msib-final-project-4/router"
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

	router.NewUserRouter(r, userHandler, authService).Route()
	router.NewCategoryRouter(r, categoryHandler, authService).Route()
	router.NewProductRouter(r, productHandler, authService).Route()

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

package handler

import (
	"fmt"
	"hacktiv8-msib-final-project-4/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.GetPostgresInstance()

	fmt.Println(db)

	r := gin.Default()

	log.Fatalln(r.Run(":" + port))
}

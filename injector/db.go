package injector

import (
	"hacktiv8-msib-final-project-4/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	db = database.GetPostgresInstance()
}

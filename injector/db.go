package injector

import (
	"gorm.io/gorm"

	"hacktiv8-msib-final-project-4/database"
)

var db *gorm.DB

func initDB() {
	db = database.GetPostgresInstance()
}

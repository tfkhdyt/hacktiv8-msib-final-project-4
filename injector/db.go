package injector

import (
	"hacktiv8-msib-final-project-4/database"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = database.GetPostgresInstance()
}

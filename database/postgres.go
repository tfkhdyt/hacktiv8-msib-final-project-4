package database

import (
	"hacktiv8-msib-final-project-4/config"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"log"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	errs.CheckErr(err)

	errs.CheckErr(db.AutoMigrate())

	log.Println("Connected to DB!")
}

func GetPostgresInstance() *gorm.DB {
	return db
}

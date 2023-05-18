package userpg

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/userrepository"
	"log"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) userrepository.UserRepository {
	return &userPG{db}
}

func (u *userPG) Register(user *entity.User) (*entity.User, errs.MessageErr) {
	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

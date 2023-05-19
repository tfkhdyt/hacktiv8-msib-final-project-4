package userpg

import (
	"fmt"
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

func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}

func (u *userPG) GetUserByID(id uint) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with id %d is not found", id))
	}

	return &user, nil
}

func (u *userPG) TopUp(id uint, balance uint) (*entity.User, errs.MessageErr) {
	user, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	user.Balance += balance

	if err := u.db.Save(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to do topup")
	}

	return user, nil
}

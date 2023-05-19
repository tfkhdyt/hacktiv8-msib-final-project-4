package userrepository

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
)

type UserRepository interface {
	Register(user *entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	GetUserByID(id uint) (*entity.User, errs.MessageErr)
	TopUp(id uint, balance uint) (*entity.User, errs.MessageErr)
}

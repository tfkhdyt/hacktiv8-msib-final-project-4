package userrepository

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
)

type UserRepository interface {
	Register(user *entity.User) (*entity.User, errs.MessageErr)
}

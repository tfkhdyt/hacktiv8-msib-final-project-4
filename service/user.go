package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/userrepository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u *userService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	user := payload.ToEntity()

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		ID:        registeredUser.ID,
		FullName:  registeredUser.FullName,
		Email:     registeredUser.Email,
		Password:  registeredUser.Password,
		Balance:   registeredUser.Balance,
		CreatedAt: registeredUser.CreatedAt,
	}

	return response, nil
}

package service

import (
	"fmt"
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/userrepository"
)

type UserService interface {
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	TopUp(id uint, payload *dto.TopUpRequest) (*dto.TopUpResponse, errs.MessageErr)
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

func (u *userService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &dto.LoginResponse{Token: token}

	return response, nil
}

func (u *userService) TopUp(id uint, payload *dto.TopUpRequest) (*dto.TopUpResponse, errs.MessageErr) {
	result, err := u.userRepo.TopUp(id, payload.Balance)
	if err != nil {
		return nil, err
	}

	response := &dto.TopUpResponse{
		Message: fmt.Sprintf("Your balance has been successfully updated to Rp %d", result.Balance),
	}

	return response, nil
}

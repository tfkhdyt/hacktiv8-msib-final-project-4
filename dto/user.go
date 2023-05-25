package dto

import (
	"hacktiv8-msib-final-project-4/entity"
	"time"
)

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email"     binding:"email,required"`
	Password string `json:"password"  binding:"required,min=6"`
}

func (r *RegisterRequest) ToEntity() *entity.User {
	return &entity.User{
		FullName: r.FullName,
		Email:    r.Email,
		Password: r.Password,
		Role:     "customer",
		Balance:  0,
	}
}

type RegisterResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

type TopUpRequest struct {
	Balance uint `json:"balance" binding:"required,min=0,max=100000000"`
}

type TopUpResponse struct {
	Message string `json:"message"`
}

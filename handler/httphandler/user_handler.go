package httphandler

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	var reqBody dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	registeredUser, err := u.userService.Register(&reqBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var reqBody dto.LoginRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	token, err := u.userService.Login(&reqBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

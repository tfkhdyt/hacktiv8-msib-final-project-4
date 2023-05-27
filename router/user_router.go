package router

import (
	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/service"
)

type UserRouter struct {
	r           *gin.Engine
	userHandler *httphandler.UserHandler
	authService service.AuthService
}

func NewUserRouter(
	r *gin.Engine,
	userHandler *httphandler.UserHandler,
	authService service.AuthService,
) *UserRouter {
	return &UserRouter{r, userHandler, authService}
}

func (u *UserRouter) Route() {
	u.r.POST("/users/register", u.userHandler.Register)
	u.r.POST("/users/login", u.userHandler.Login)
	u.r.PATCH("/users/topup", u.authService.Authentication(), u.userHandler.TopUp)
}

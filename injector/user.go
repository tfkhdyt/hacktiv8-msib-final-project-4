package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/userrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	userRepo    userrepository.UserRepository
	userService service.UserService
	userHandler *httphandler.UserHandler
)

func initUser() {
	userRepo = userpg.NewUserPG(db)
	userService = service.NewUserService(userRepo)
	userHandler = httphandler.NewUserHandler(userService)
}

func InitializeUserHandler() *httphandler.UserHandler {
	return userHandler
}

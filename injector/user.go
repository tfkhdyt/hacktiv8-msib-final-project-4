package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/userrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	UserRepo    userrepository.UserRepository
	UserService service.UserService
	UserHandler *httphandler.UserHandler
)

func init() {
	UserRepo = userpg.NewUserPG(DB)
	UserService = service.NewUserService(UserRepo)
	UserHandler = httphandler.NewUserHandler(UserService)
}

func InitializeUserHandler() *httphandler.UserHandler {
	return UserHandler
}

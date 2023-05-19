package injector

import "hacktiv8-msib-final-project-4/service"

var AuthService service.AuthService

func init() {
	AuthService = service.NewAuthService(UserRepo)
}

func InitializeAuthService() service.AuthService {
	return AuthService
}

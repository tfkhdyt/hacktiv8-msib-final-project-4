package injector

import "hacktiv8-msib-final-project-4/service"

var authService service.AuthService

func initAuth() {
	authService = service.NewAuthService(userRepo)
}

func InitializeAuthService() service.AuthService {
	return authService
}

package injector

import (
	"hacktiv8-msib-final-project-4/database"
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-4/repository/userrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository/userpg"
	"hacktiv8-msib-final-project-4/service"

	"gorm.io/gorm"
)

var (
	db *gorm.DB

	userRepo    userrepository.UserRepository
	userService service.UserService
	userHandler *httphandler.UserHandler

	authService service.AuthService

	categoryRepo    categoryrepository.CategoryRepository
	categoryService service.CategoryService
	categoryHandler *httphandler.CategoryHandler
)

func init() {
	db = database.GetPostgresInstance()

	userRepo = userpg.NewUserPG(db)
	userService = service.NewUserService(userRepo)
	userHandler = httphandler.NewUserHandler(userService)

	authService = service.NewAuthService(userRepo)

	categoryRepo = categorypg.NewCategoryPG(db)
	categoryService = service.NewCategoryService(categoryRepo)
	categoryHandler = httphandler.NewCategoryHandler(categoryService)
}

func InitializeUserHandler() *httphandler.UserHandler {
	return userHandler
}

func InitializeAuthService() service.AuthService {
	return authService
}

func InitializeCategoryHandler() *httphandler.CategoryHandler {
	return categoryHandler
}

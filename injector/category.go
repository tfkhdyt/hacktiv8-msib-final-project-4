package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	categoryRepo    categoryrepository.CategoryRepository
	categoryService service.CategoryService
	categoryHandler *httphandler.CategoryHandler
)

func initCategory() {
	categoryRepo = categorypg.NewCategoryPG(db)
	categoryService = service.NewCategoryService(categoryRepo)
	categoryHandler = httphandler.NewCategoryHandler(categoryService)
}

func InitializeCategoryHandler() *httphandler.CategoryHandler {
	return categoryHandler
}

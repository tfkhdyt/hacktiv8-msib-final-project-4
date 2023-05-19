package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	CategoryRepo    categoryrepository.CategoryRepository
	CategoryService service.CategoryService
	CategoryHandler *httphandler.CategoryHandler
)

func init() {
	CategoryRepo = categorypg.NewCategoryPG(DB)
	CategoryService = service.NewCategoryService(CategoryRepo)
	CategoryHandler = httphandler.NewCategoryHandler(CategoryService)
}

func InitializeCategoryHandler() *httphandler.CategoryHandler {
	return CategoryHandler
}

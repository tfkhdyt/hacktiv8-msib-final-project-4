package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/categoryrepository/categorypg"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository/productpg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	categoryRepo    categoryrepository.CategoryRepository
	categoryService service.CategoryService
	categoryHandler *httphandler.CategoryHandler

	productRepo    productrepository.ProductRepository
	productService service.ProductService
	productHandler *httphandler.ProductHandler
)

func initCategoryAndProduct() {
	categoryRepo = categorypg.NewCategoryPG(db)
	productRepo = productpg.NewProductPG(db)

	categoryService = service.NewCategoryService(categoryRepo, productRepo)
	productService = service.NewProductService(productRepo, categoryRepo)

	categoryHandler = httphandler.NewCategoryHandler(categoryService)
	productHandler = httphandler.NewProductHandler(productService)
}

func InitializeCategoryHandler() *httphandler.CategoryHandler {
	return categoryHandler
}

func InitializeProductHandler() *httphandler.ProductHandler {
	return productHandler
}

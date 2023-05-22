package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository/productpg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	productRepo    productrepository.ProductRepository
	productService service.ProductService
	productHandler *httphandler.ProductHandler
)

func initProduct() {
	productRepo = productpg.NewProductPG(db)
	productService = service.NewProductService(productRepo, categoryRepo)
	productHandler = httphandler.NewProductHandler(productService)
}

func InitializeProductHandler() *httphandler.ProductHandler {
	return productHandler
}

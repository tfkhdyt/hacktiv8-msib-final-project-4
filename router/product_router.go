package router

import (
	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/service"
)

type ProductRouter struct {
	r              *gin.Engine
	productHandler *httphandler.ProductHandler
	authService    service.AuthService
}

func NewProductRouter(
	r *gin.Engine,
	productHandler *httphandler.ProductHandler,
	authService service.AuthService,
) *ProductRouter {
	return &ProductRouter{r, productHandler, authService}
}

func (p *ProductRouter) Route() {
	p.r.POST(
		"/products",
		p.authService.Authentication(),
		p.authService.AdminAuthorization(),
		p.productHandler.CreateProduct,
	)
	p.r.GET("/products", p.authService.Authentication(), p.productHandler.GetAllProducts)
	p.r.PUT(
		"/products/:productID",
		p.authService.Authentication(),
		p.authService.AdminAuthorization(),
		p.productHandler.UpdateProduct,
	)
	p.r.DELETE(
		"/products/:productID",
		p.authService.Authentication(),
		p.authService.AdminAuthorization(),
		p.productHandler.DeleteProduct,
	)
}

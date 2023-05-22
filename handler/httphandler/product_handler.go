package httphandler

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var reqBody dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	createdProduct, err := p.productService.CreateProduct(&reqBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (p *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

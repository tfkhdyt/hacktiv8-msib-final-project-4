package httphandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
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

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("productID")
	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		validationError := errs.NewBadRequest("Product id should be in unsigned integer")
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	var reqBody dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	updatedProduct, errUpdate := p.productService.UpdateProduct(uint(productIDUint), &reqBody)
	if errUpdate != nil {
		ctx.JSON(errUpdate.StatusCode(), errUpdate)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("productID")
	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		validationError := errs.NewBadRequest("Product id should be in unsigned integer")
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	response, errDelete := p.productService.DeleteProduct(uint(productIDUint))
	if errDelete != nil {
		ctx.JSON(errDelete.StatusCode(), errDelete)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

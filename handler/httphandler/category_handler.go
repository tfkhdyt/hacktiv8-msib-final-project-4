package httphandler

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (c *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var reqBody dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	createdCategory, err := c.categoryService.CreateCategory(&reqBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdCategory)
}
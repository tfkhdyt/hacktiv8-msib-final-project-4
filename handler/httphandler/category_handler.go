package httphandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
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

func (c *CategoryHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	categoryID := ctx.Param("categoryID")
	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 32)
	if err != nil {
		validationError := errs.NewBadRequest("Category id should be in unsigned integer")
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	var reqBody dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		validationError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	updatedCategory, updateErr := c.categoryService.UpdateCategory(uint(categoryIDUint), &reqBody)
	if updateErr != nil {
		ctx.JSON(updateErr.StatusCode(), updateErr)
		return
	}

	ctx.JSON(http.StatusOK, updatedCategory)
}

func (c *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("categoryID")
	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 32)
	if err != nil {
		validationError := errs.NewBadRequest("Category id should be in unsigned integer")
		ctx.JSON(validationError.StatusCode(), validationError)
		return
	}

	response, deleteErr := c.categoryService.DeleteCategory(uint(categoryIDUint))
	if deleteErr != nil {
		ctx.JSON(deleteErr.StatusCode(), deleteErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

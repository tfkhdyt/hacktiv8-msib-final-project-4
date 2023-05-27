package router

import (
	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/service"
)

type CategoryRouter struct {
	r               *gin.Engine
	categoryHandler *httphandler.CategoryHandler
	authService     service.AuthService
}

func NewCategoryRouter(
	r *gin.Engine,
	categoryHandler *httphandler.CategoryHandler,
	authService service.AuthService,
) *CategoryRouter {
	return &CategoryRouter{r, categoryHandler, authService}
}

func (c *CategoryRouter) Route() {
	c.r.POST("/categories",
		c.authService.Authentication(),
		c.authService.AdminAuthorization(),
		c.categoryHandler.CreateCategory,
	)
	c.r.GET(
		"/categories",
		c.authService.Authentication(),
		c.authService.AdminAuthorization(),
		c.categoryHandler.GetAllCategories,
	)
	c.r.PATCH(
		"/categories/:categoryID",
		c.authService.Authentication(),
		c.authService.AdminAuthorization(),
		c.categoryHandler.UpdateCategory,
	)
	c.r.DELETE(
		"/categories/:categoryID",
		c.authService.Authentication(),
		c.authService.AdminAuthorization(),
		c.categoryHandler.DeleteCategory,
	)
}

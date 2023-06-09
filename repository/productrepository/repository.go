package productrepository

import (
	"gorm.io/gorm"

	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, errs.MessageErr)
	GetAllProducts() ([]entity.Product, errs.MessageErr)
	GetProductByID(id uint) (*entity.Product, errs.MessageErr)
	GetAllProductsByCategoryID(categoryID uint) ([]entity.Product, errs.MessageErr)
	UpdateProduct(
		oldProduct *entity.Product,
		newProduct *entity.Product,
	) (*entity.Product, errs.MessageErr)
	DeleteProduct(product *entity.Product) errs.MessageErr
	DecrementStock(id uint, quantity uint, tx *gorm.DB) errs.MessageErr
}

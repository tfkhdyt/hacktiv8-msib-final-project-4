package productpg

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"log"

	"gorm.io/gorm"
)

type productPG struct {
	db *gorm.DB
}

func NewProductPG(db *gorm.DB) productrepository.ProductRepository {
	return &productPG{db}
}

func (p *productPG) CreateProduct(product *entity.Product) (*entity.Product, errs.MessageErr) {
	if err := p.db.Create(product).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new product")
	}

	return product, nil
}

func (p *productPG) GetAllProducts() ([]entity.Product, errs.MessageErr) {
	var products []entity.Product
	if err := p.db.Find(&products).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to get all products")
	}

	return products, nil
}

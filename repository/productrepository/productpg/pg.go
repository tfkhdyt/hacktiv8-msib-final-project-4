package productpg

import (
	"fmt"
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

func (p *productPG) GetAllProductsByCategoryID(categoryID uint) ([]entity.Product, errs.MessageErr) {
	var products []entity.Product
	if err := p.db.Find(&products, "category_id = ?", categoryID).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to get all products")
	}

	return products, nil
}

func (p *productPG) GetProductByID(id uint) (*entity.Product, errs.MessageErr) {
	var product *entity.Product
	if err := p.db.First(&product, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Product with id %d is not found", id))
	}

	return product, nil
}

func (p *productPG) UpdateProduct(oldProduct *entity.Product, newProduct *entity.Product) (*entity.Product, errs.MessageErr) {
	if err := p.db.Model(oldProduct).Updates(newProduct).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update product with id %d", oldProduct.ID))
	}

	return oldProduct, nil
}

func (p *productPG) DeleteProduct(product *entity.Product) errs.MessageErr {
	if err := p.db.Delete(product).Error; err != nil {
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete product with id %d", product.ID))
	}

	return nil
}

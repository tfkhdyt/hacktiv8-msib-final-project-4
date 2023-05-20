package categorypg

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"log"

	"gorm.io/gorm"
)

type categoryPG struct {
	db *gorm.DB
}

func NewCategoryPG(db *gorm.DB) categoryrepository.CategoryRepository {
	return &categoryPG{db}
}

func (c *categoryPG) CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr) {
	if err := c.db.Create(category).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new category")
	}

	return category, nil
}

func (c *categoryPG) GetAllCategories() ([]entity.Category, errs.MessageErr) {
	var categories []entity.Category
	if err := c.db.Find(&categories).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new category")
	}

	return categories, nil
}

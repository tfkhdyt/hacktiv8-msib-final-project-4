package categoryrepository

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(category *entity.Category) (*entity.Category, errs.MessageErr)
	GetAllCategories() ([]entity.Category, errs.MessageErr)
}

package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
)

type CategoryService interface {
	CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
}

func NewCategoryService(categoryRepo categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (c *categoryService) CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr) {
	category := payload.ToEntity()

	createdCategory, err := c.categoryRepo.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateCategoryResponse{
		ID:                createdCategory.ID,
		Type:              createdCategory.Type,
		SoldProductAmount: createdCategory.SoldProductAmount,
		CreatedAt:         createdCategory.CreatedAt,
	}

	return response, nil
}

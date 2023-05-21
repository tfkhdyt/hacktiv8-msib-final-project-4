package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
)

type CategoryService interface {
	CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr)
	GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr)
	UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
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

func (c *categoryService) GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr) {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllCategoriesResponse{}
	for _, category := range categories {
		response = append(response, dto.GetAllCategoriesResponse{
			ID:                category.ID,
			Type:              category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
			UpdatedAt:         category.UpdatedAt,
		})
	}

	return response, nil
}

func (c *categoryService) UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr) {
	oldCategory, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	newCategory := payload.ToEntity()

	updatedCategory, updateErr := c.categoryRepo.UpdateCategory(oldCategory, newCategory)
	if updateErr != nil {
		return nil, updateErr
	}

	response := &dto.UpdateCategoryResponse{
		ID:                updatedCategory.ID,
		Type:              updatedCategory.Type,
		SoldProductAmount: updatedCategory.SoldProductAmount,
		UpdatedAt:         updatedCategory.UpdatedAt,
	}

	return response, nil
}

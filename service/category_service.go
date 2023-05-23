package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository"
)

type CategoryService interface {
	CreateCategory(payload *dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, errs.MessageErr)
	GetAllCategories() ([]dto.GetAllCategoriesResponse, errs.MessageErr)
	GetCategoryByID(id uint) (*dto.GetCategoryByIDResponse, errs.MessageErr)
	UpdateCategory(id uint, payload *dto.UpdateCategoryRequest) (*dto.UpdateCategoryResponse, errs.MessageErr)
	DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo categoryrepository.CategoryRepository
	productRepo  productrepository.ProductRepository
}

func NewCategoryService(categoryRepo categoryrepository.CategoryRepository, productRepo productrepository.ProductRepository) CategoryService {
	return &categoryService{categoryRepo, productRepo}
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
		products, err := c.productRepo.GetAllProductsByCategoryID(category.ID)
		if err != nil {
			return nil, err
		}

		productsData := []dto.ProductData{}
		for _, product := range products {
			productsData = append(productsData, dto.ProductData{
				ID:        product.ID,
				Title:     product.Title,
				Price:     product.Price,
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			})
		}

		response = append(response, dto.GetAllCategoriesResponse{
			ID:                category.ID,
			Type:              category.Type,
			SoldProductAmount: category.SoldProductAmount,
			CreatedAt:         category.CreatedAt,
			UpdatedAt:         category.UpdatedAt,
			Products:          productsData,
		})
	}

	return response, nil
}

func (c *categoryService) GetCategoryByID(id uint) (*dto.GetCategoryByIDResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	response := &dto.GetCategoryByIDResponse{
		ID:                category.ID,
		Type:              category.Type,
		SoldProductAmount: category.SoldProductAmount,
		CreatedAt:         category.CreatedAt,
		UpdatedAt:         category.UpdatedAt,
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

func (c *categoryService) DeleteCategory(id uint) (*dto.DeleteCategoryResponse, errs.MessageErr) {
	category, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(category); err != nil {
		return nil, err
	}

	response := &dto.DeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}

	return response, nil
}

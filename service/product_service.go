package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository"
)

type ProductService interface {
	CreateProduct(payload *dto.CreateProductRequest) (*dto.CreateProductResponse, errs.MessageErr)
}

type productService struct {
	productRepo  productrepository.ProductRepository
	categoryRepo categoryrepository.CategoryRepository
}

func NewProductService(productRepo productrepository.ProductRepository, categoryRepo categoryrepository.CategoryRepository) ProductService {
	return &productService{productRepo, categoryRepo}
}

func (p *productService) CreateProduct(payload *dto.CreateProductRequest) (*dto.CreateProductResponse, errs.MessageErr) {
	product := payload.ToEntity()

	_, checkCategoryErr := p.categoryRepo.GetCategoryByID(product.CategoryID)
	if checkCategoryErr != nil {
		return nil, checkCategoryErr
	}

	createdProduct, err := p.productRepo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateProductResponse{
		ID:         createdProduct.ID,
		Title:      createdProduct.Title,
		Price:      createdProduct.Price,
		Stock:      createdProduct.Stock,
		CategoryID: createdProduct.CategoryID,
		CreatedAt:  createdProduct.CreatedAt,
	}

	return response, nil
}

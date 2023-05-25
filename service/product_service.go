package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository"

	"github.com/leekchan/accounting"
)

var (
	lc = accounting.LocaleInfo["IDR"]
	ac = accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
		Thousand:  lc.ThouSep,
		Decimal:   lc.DecSep,
	}
)

type ProductService interface {
	CreateProduct(payload *dto.CreateProductRequest) (*dto.CreateProductResponse, errs.MessageErr)
	GetAllProducts() ([]dto.GetAllProductsResponse, errs.MessageErr)
	UpdateProduct(
		id uint,
		payload *dto.UpdateProductRequest,
	) (*dto.UpdateProductResponse, errs.MessageErr)
	DeleteProduct(id uint) (*dto.DeleteProductResponse, errs.MessageErr)
}

type productService struct {
	productRepo  productrepository.ProductRepository
	categoryRepo categoryrepository.CategoryRepository
}

func NewProductService(
	productRepo productrepository.ProductRepository,
	categoryRepo categoryrepository.CategoryRepository,
) ProductService {
	return &productService{productRepo, categoryRepo}
}

func (p *productService) CreateProduct(
	payload *dto.CreateProductRequest,
) (*dto.CreateProductResponse, errs.MessageErr) {
	product := payload.ToEntity()

	if _, checkCategoryErr := p.categoryRepo.GetCategoryByID(product.CategoryID); checkCategoryErr != nil {
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

func (p *productService) GetAllProducts() ([]dto.GetAllProductsResponse, errs.MessageErr) {
	products, err := p.productRepo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllProductsResponse{}
	for _, product := range products {
		response = append(response, dto.GetAllProductsResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		})
	}

	return response, nil
}

func (p *productService) UpdateProduct(
	id uint,
	payload *dto.UpdateProductRequest,
) (*dto.UpdateProductResponse, errs.MessageErr) {
	product := payload.ToEntity()

	if _, checkCategoryErr := p.categoryRepo.GetCategoryByID(product.CategoryID); checkCategoryErr != nil {
		return nil, checkCategoryErr
	}

	oldProduct, err := p.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	updatedProduct, errUpdate := p.productRepo.UpdateProduct(oldProduct, product)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := &dto.UpdateProductResponse{
		Product: dto.ProductDataWithCategoryID{
			ID:         updatedProduct.ID,
			Title:      updatedProduct.Title,
			Price:      ac.FormatMoney(updatedProduct.Price),
			Stock:      updatedProduct.Stock,
			CategoryID: updatedProduct.CategoryID,
			CreatedAt:  updatedProduct.CreatedAt,
			UpdatedAt:  updatedProduct.UpdatedAt,
		},
	}

	return response, nil
}

func (p *productService) DeleteProduct(id uint) (*dto.DeleteProductResponse, errs.MessageErr) {
	product, err := p.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	if err := p.productRepo.DeleteProduct(product); err != nil {
		return nil, err
	}

	response := &dto.DeleteProductResponse{
		Message: "Product has been successfully deleted",
	}

	return response, nil
}

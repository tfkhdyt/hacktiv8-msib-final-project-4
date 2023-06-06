package dto

import (
	"time"

	"hacktiv8-msib-final-project-4/entity"
)

type ProductData struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	Stock     uint      `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Title      string `json:"title"       binding:"required"`
	Price      uint   `json:"price"       binding:"required,max=50000000,min=0"`
	Stock      uint   `json:"stock"       binding:"required,min=5"`
	CategoryID uint   `json:"category_Id"`
}

func (p *CreateProductRequest) ToEntity() *entity.Product {
	return &entity.Product{
		Title:      p.Title,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
	}
}

type CreateProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetAllProductsResponse CreateProductResponse

type UpdateProductRequest CreateProductRequest

func (p *UpdateProductRequest) ToEntity() *entity.Product {
	return &entity.Product{
		Title:      p.Title,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
	}
}

type UpdateProductResponse struct {
	Product ProductDataWithCategoryID `json:"product"`
}

type ProductDataWithCategoryID struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      string    `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductDataWithCategoryIDAndIntegerPrice struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DeleteProductResponse struct {
	Message string `json:"message"`
}

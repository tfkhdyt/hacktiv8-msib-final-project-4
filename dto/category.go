package dto

import (
	"hacktiv8-msib-final-project-4/entity"
	"time"
)

type CreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

func (c *CreateCategoryRequest) ToEntity() *entity.Category {
	return &entity.Category{
		Type:              c.Type,
		SoldProductAmount: 0,
	}
}

type CreateCategoryResponse struct {
	ID                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type GetAllCategoriesResponse struct {
	ID                uint          `json:"id"`
	Type              string        `json:"type"`
	SoldProductAmount uint          `json:"sold_product_amount"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Products          []ProductData `json:"products"`
}

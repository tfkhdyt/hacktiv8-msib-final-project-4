package dto

import "hacktiv8-msib-final-project-4/entity"

type CreateTransactionRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"   binding:"required,min=1"`
}

func (t *CreateTransactionRequest) ToEntity() *entity.TransactionHistory {
	return &entity.TransactionHistory{
		ProductID: t.ProductID,
		Quantity:  t.Quantity,
	}
}

type CreateTransactionResponse struct {
	Message         string          `json:"message"`
	TransactionBill TransactionBill `json:"transaction_bill"`
}

type TransactionBill struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

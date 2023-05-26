package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository"
)

type TransactionHistoryService interface {
	CreateTransaction(
		user *entity.User,
		payload *dto.CreateTransactionRequest,
	) (*dto.CreateTransactionResponse, errs.MessageErr)
}

type transactionHistoryService struct {
	transactionRepo transactionhistoryrepository.TransactionHistoryRepository
	productRepo     productrepository.ProductRepository
}

func NewTransactionHistoryService(
	transactionRepo transactionhistoryrepository.TransactionHistoryRepository,
	productRepo productrepository.ProductRepository,
) TransactionHistoryService {
	return &transactionHistoryService{transactionRepo, productRepo}
}

func (t *transactionHistoryService) CreateTransaction(
	user *entity.User,
	payload *dto.CreateTransactionRequest,
) (*dto.CreateTransactionResponse, errs.MessageErr) {
	transaction := payload.ToEntity()

	product, err := t.productRepo.GetProductByID(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	transaction.TotalPrice = product.Price * transaction.Quantity

	if err := product.CheckStock(transaction.Quantity); err != nil {
		return nil, err
	}

	if err := user.CheckBalance(transaction.TotalPrice); err != nil {
		return nil, err
	}

	createdTransaction, err := t.transactionRepo.CreateTransaction(user, product, transaction)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateTransactionResponse{
		Message: "You have successfully purchased the product",
		TransactionBill: dto.TransactionBill{
			TotalPrice:   createdTransaction.TotalPrice,
			Quantity:     createdTransaction.Quantity,
			ProductTitle: product.Title,
		},
	}

	return response, nil
}

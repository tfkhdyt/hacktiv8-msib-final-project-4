package service

import (
	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository"
)

type TransactionHistoryService interface {
	CreateTransaction(
		user *entity.User,
		payload *dto.CreateTransactionRequest,
	) (*dto.CreateTransactionResponse, errs.MessageErr)

	GetTransactionsByUserID(userID uint) ([]dto.GetTransactionsByUserIDResponse, errs.MessageErr)

	GetUserTransactions() ([]dto.GetUserTransactionsResponse, errs.MessageErr)
}

type transactionHistoryService struct {
	transactionRepo transactionhistoryrepository.TransactionHistoryRepository
	productRepo     productrepository.ProductRepository
	userRepo        userrepository.UserRepository
}

func NewTransactionHistoryService(
	transactionRepo transactionhistoryrepository.TransactionHistoryRepository,
	productRepo productrepository.ProductRepository,
	userRepo userrepository.UserRepository,
) TransactionHistoryService {
	return &transactionHistoryService{transactionRepo, productRepo, userRepo}
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

func (t *transactionHistoryService) GetTransactionsByUserID(userID uint) ([]dto.GetTransactionsByUserIDResponse, errs.MessageErr) {
	transactions, err := t.transactionRepo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := []dto.GetTransactionsByUserIDResponse{}
	for _, transaction := range transactions {
		product, err := t.productRepo.GetProductByID(transaction.ProductID)
		if err != nil {
			return nil, err
		}

		response = append(response, dto.GetTransactionsByUserIDResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: dto.ProductDataWithCategoryIDAndIntegerPrice{
				ID:         product.ID,
				Title:      product.Title,
				Price:      product.Price,
				Stock:      product.Stock,
				CategoryID: product.CategoryID,
				CreatedAt:  product.CreatedAt,
				UpdatedAt:  product.UpdatedAt,
			},
		})
	}

	return response, nil
}

func (t *transactionHistoryService) GetUserTransactions() ([]dto.GetUserTransactionsResponse, errs.MessageErr) {
	transactions, err := t.transactionRepo.GetUserTransactions()
	if err != nil {
		return nil, err
	}

	response := []dto.GetUserTransactionsResponse{}
	for _, transaction := range transactions {
		product, err := t.productRepo.GetProductByID(transaction.ProductID)
		if err != nil {
			return nil, err
		}

		user, errGetUser := t.userRepo.GetUserByID(transaction.UserID)
		if errGetUser != nil {
			return nil, errGetUser
		}

		response = append(response, dto.GetUserTransactionsResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: dto.ProductDataWithCategoryIDAndIntegerPrice{
				ID:         product.ID,
				Title:      product.Title,
				Price:      product.Price,
				Stock:      product.Stock,
				CategoryID: product.CategoryID,
				CreatedAt:  product.CreatedAt,
				UpdatedAt:  product.UpdatedAt,
			},
			User: dto.UserData{
				ID:        user.ID,
				Email:     user.Email,
				FullName:  user.FullName,
				Balance:   user.Balance,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		})
	}

	return response, nil
}

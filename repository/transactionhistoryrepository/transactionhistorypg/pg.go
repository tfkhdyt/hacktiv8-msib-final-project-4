package transactionhistorypg

import (
	"fmt"
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository"
	"log"

	"gorm.io/gorm"
)

type transactionHistoryPG struct {
	db           *gorm.DB
	productRepo  productrepository.ProductRepository
	userRepo     userrepository.UserRepository
	categoryRepo categoryrepository.CategoryRepository
}

func NewTransactionHistoryPG(
	db *gorm.DB,
	productRepo productrepository.ProductRepository,
	userRepo userrepository.UserRepository,
	categoryRepo categoryrepository.CategoryRepository,
) transactionhistoryrepository.TransactionHistoryRepository {
	return &transactionHistoryPG{db, productRepo, userRepo, categoryRepo}
}

func (t *transactionHistoryPG) CreateTransaction(user *entity.User, product *entity.Product, transaction *entity.TransactionHistory) (*entity.TransactionHistory, errs.MessageErr) {
	transaction.TotalPrice = product.Price * transaction.Quantity

	if product.Stock < transaction.Quantity {
		return nil, errs.NewBadRequest(fmt.Sprintf("Insufficient product stock. There are only %d items left in stock", product.Stock))
	}

	if user.Balance < transaction.TotalPrice {
		return nil, errs.NewBadRequest("Your balance is not sufficient")
	}

	tx := t.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := t.productRepo.DecrementStock(product.ID, transaction.Quantity, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := t.userRepo.DecrementBalance(user.ID, transaction.TotalPrice, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := t.categoryRepo.IncrementSoldProductAmount(product.CategoryID, transaction.Quantity, tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(user).Association("TransactionHistories").Append(transaction); err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new transaction")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to commit transaction")
	}

	return transaction, nil
}

package transactionhistorypg

import (
	"fmt"
	"log"

	"github.com/leekchan/accounting"
	"gorm.io/gorm"

	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/repository/categoryrepository"
	"hacktiv8-msib-final-project-4/repository/productrepository"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository"
	"hacktiv8-msib-final-project-4/repository/userrepository"
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

func (t *transactionHistoryPG) CreateTransaction(
	user *entity.User,
	product *entity.Product,
	transaction *entity.TransactionHistory,
) (*entity.TransactionHistory, errs.MessageErr) {
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

func (t *transactionHistoryPG) GetTransactionsByUserID(userID uint) ([]entity.TransactionHistory, errs.MessageErr) {
	var transactions []entity.TransactionHistory
	if err := t.db.Find(&transactions, "user_id = ?", userID).Error; err != nil {
		return nil, errs.NewInternalServerError(
			fmt.Sprintf(
				"Failed to get transaction histories of user with id %d",
				userID,
			),
		)
	}

	return transactions, nil
}

func (t *transactionHistoryPG) GetUserTransactions() ([]entity.TransactionHistory, errs.MessageErr) {
	var transactions []entity.TransactionHistory
	if err := t.db.Find(&transactions).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all transactions")
	}

	return transactions, nil
}

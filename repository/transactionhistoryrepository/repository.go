package transactionhistoryrepository

import (
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
)

type TransactionHistoryRepository interface {
	CreateTransaction(user *entity.User, product *entity.Product, transaction *entity.TransactionHistory) (*entity.TransactionHistory, errs.MessageErr)
}

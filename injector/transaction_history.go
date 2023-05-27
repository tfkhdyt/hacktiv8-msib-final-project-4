package injector

import (
	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository"
	"hacktiv8-msib-final-project-4/repository/transactionhistoryrepository/transactionhistorypg"
	"hacktiv8-msib-final-project-4/service"
)

var (
	transactionRepo    transactionhistoryrepository.TransactionHistoryRepository
	transactionService service.TransactionHistoryService
	transactionHandler *httphandler.TransactionHistoryHandler
)

func initTransactionHistory() {
	transactionRepo = transactionhistorypg.NewTransactionHistoryPG(
		db,
		productRepo,
		userRepo,
		categoryRepo,
	)
	transactionService = service.NewTransactionHistoryService(transactionRepo, productRepo, userRepo)
	transactionHandler = httphandler.NewTransactionHistoryHandler(transactionService)
}

func InitializeTransactionHistoryHandler() *httphandler.TransactionHistoryHandler {
	return transactionHandler
}

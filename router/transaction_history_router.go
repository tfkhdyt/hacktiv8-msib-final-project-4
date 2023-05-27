package router

import (
	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/handler/httphandler"
	"hacktiv8-msib-final-project-4/service"
)

type TransactionRouter struct {
	r                  *gin.Engine
	transactionHandler *httphandler.TransactionHistoryHandler
	authService        service.AuthService
}

func NewTransactionRouter(
	r *gin.Engine,
	transactionHandler *httphandler.TransactionHistoryHandler,
	authService service.AuthService,
) *TransactionRouter {
	return &TransactionRouter{r, transactionHandler, authService}
}

func (t *TransactionRouter) Route() {
	t.r.POST(
		"/transactions",
		t.authService.Authentication(),
		t.transactionHandler.CreateTransaction,
	)
	t.r.GET(
		"/transactions/my-transactions",
		t.authService.Authentication(),
		t.transactionHandler.GetTransactionsByUserID,
	)
	t.r.GET(
		"/transactions/all-transactions",
		t.authService.Authentication(),
		t.authService.AdminAuthorization(),
		t.transactionHandler.GetAllTransactions,
	)
}

package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hacktiv8-msib-final-project-4/dto"
	"hacktiv8-msib-final-project-4/entity"
	"hacktiv8-msib-final-project-4/pkg/errs"
	"hacktiv8-msib-final-project-4/service"
)

type TransactionHistoryHandler struct {
	transactionService service.TransactionHistoryService
}

func NewTransactionHistoryHandler(
	transactionService service.TransactionHistoryService,
) *TransactionHistoryHandler {
	return &TransactionHistoryHandler{transactionService}
}

func (t *TransactionHistoryHandler) CreateTransaction(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var reqBody dto.CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		errValidation := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(errValidation.StatusCode(), errValidation)
		return
	}

	createdTransaction, err := t.transactionService.CreateTransaction(userData, &reqBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdTransaction)
}

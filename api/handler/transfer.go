package handler

import (
	"fmt"
	"net/http"
	"simple-bank-v2/entity"
	"simple-bank-v2/usecase/transfer"

	"github.com/gin-gonic/gin"
)

type transferRequest struct{
	FromAccountID 	int64 	`json:"from_account_id" binding:"required,min=1"`
	ToAccountID 	int64 	`json:"to_account_id" binding:"required,min=1"`
	Amount 			int64 	`json:"amount" binding:"required,gt=0"`
}

func createTransfer(service transfer.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req transferRequest
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(req)
		result, err := service.TransferTx(entity.TransferTxParams{
			FromAccountID: req.FromAccountID,
			ToAccountID: req.ToAccountID,
			Amount: req.Amount,
		})
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func MakeTransferHandlers(group *gin.RouterGroup, service transfer.UseCase) {
	group.POST("/", createTransfer(service))
}
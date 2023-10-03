package handler

import (
	"fmt"
	"net/http"
	"simple-bank-v2/usecase/account"

	"github.com/gin-gonic/gin"
)

type CreateAccountParams struct {
	UserID int64 `form:"user_id" binding:"required"`
	Bank   string `form:"bank"`
}

func createAccount(service account.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req CreateAccountParams
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(req)
		account, err := service.CreateAccount(req.UserID, req.Bank)
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, account)
	}
}

type GetAccountParams struct{
	ID int64 `uri:"id" binding:"required"`
}

func getAccount(service account.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req GetAccountParams
		if err := c.ShouldBindUri(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		account, err := service.GetAccount(req.ID)
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, account)
	}
}

type UpdateAccountParams struct{
	ID int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func updateAccount(service account.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req UpdateAccountParams
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
		}
		// TODO: ...
		c.JSON(http.StatusOK,nil)
	}
}

func MakeAccountHandlers(group *gin.RouterGroup, service account.UseCase) {
	group.POST("/", createAccount(service))
	group.GET("/:id", getAccount(service))
	group.PUT("/:id", updateAccount(service))

}
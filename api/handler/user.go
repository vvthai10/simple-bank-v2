package handler

import (
	"fmt"
	"net/http"
	"simple-bank-v2/usecase/user"

	"github.com/gin-gonic/gin"
)

func listUsers(service user.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		users, err := service.ListUsers()
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, users)
	}
}

type CreateUserParams struct{
	Email string `form:"email" binding:"required"`
	Name string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func createUser(service user.UseCase) func(c *gin.Context) {
	return func(c *gin.Context){
		var req CreateUserParams
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(req)
		user, err := service.CreateUser(req.Email, req.Name, req.Password)
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

type GetUserParams struct{
	ID int64 `uri:"id" binding:"required"`
}

func getUser(service user.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req GetUserParams
		if err := c.ShouldBindUri(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		user, err := service.GetUser(req.ID)
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

type DeleteUserParams struct{
	ID int64 `uri:"id" binding:"required"`
}

func deleteUser(service user.UseCase) func(c *gin.Context){
	return func(c *gin.Context){
		var req DeleteUserParams
		if err := c.ShouldBindUri(&req); err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		err := service.DeleteUser(req.ID)
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func MakeUserHandlers(group *gin.RouterGroup, service user.UseCase) {
	group.GET("/", listUsers(service))
	group.POST("/", createUser(service))
	group.GET("/:id", getUser(service))
	group.DELETE("/:id", deleteUser(service))
}

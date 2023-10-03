package main

import (
	"fmt"
	"log"
	"simple-bank-v2/api"
	"simple-bank-v2/api/handler"
	repository "simple-bank-v2/infrastructure/repository/gorm"
	"simple-bank-v2/usecase/account"
	"simple-bank-v2/usecase/transfer"
	"simple-bank-v2/usecase/user"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Main file")
	// TODO 1: Connect DB
	postgresDB, err := api.NewDatabaseByGORM()
	if err != nil{
		fmt.Println("Can't connect to db: ", err)
	}
	// TODO 2: Add DB to Repository
	userRepo := repository.NewUserDB(postgresDB)
	accountRepo := repository.NewAccountDB(postgresDB)
	transferRepo := repository.NewTransferDB(postgresDB)
	// TODO 3: Add repo to Service
	userService := user.NewService(userRepo)
	accountService := account.NewService(accountRepo)
	transferService := transfer.NewService(transferRepo)

	// TODO 4: Init server
	engine := gin.Default()

	// TODO 5: Add service to Controller
	v1 := engine.Group("/users")
	handler.MakeUserHandlers(v1, userService)
	v2 := engine.Group("/accounts")
	handler.MakeAccountHandlers(v2, accountService)
	v3 := engine.Group("/transfer")
	handler.MakeTransferHandlers(v3, transferService)

	if err := engine.Run(":5000"); err != nil {
		log.Fatalln(err)
	}
}
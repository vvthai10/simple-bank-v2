package account

import (
	"fmt"
	"simple-bank-v2/api"
	repository "simple-bank-v2/infrastructure/repository/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoAccount := repository.NewAccountDB(db)
	service := NewService(repoAccount)

	a, err := service.CreateAccount(5, "MB Bank")
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_GetAccount(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoAccount := repository.NewAccountDB(db)
	service := NewService(repoAccount)

	a, err := service.GetAccount(2)
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_UpdateAccount(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoAccount := repository.NewAccountDB(db)
	service := NewService(repoAccount)

	account1, err := service.CreateAccount(1, "TBank")
	assert.Nil(t, err)

	account1.Balance = 500
	account2, err := service.UpdateAccount(account1)
	assert.Nil(t, err)
	fmt.Println(account2)
}

func Test_DeleteUser(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoAccount := repository.NewAccountDB(db)
	service := NewService(repoAccount)

	err = service.DeleteAccount(2)
	assert.Nil(t, err)
}
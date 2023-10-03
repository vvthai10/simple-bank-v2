package user

import (
	"fmt"
	"simple-bank-v2/api"
	repository "simple-bank-v2/infrastructure/repository/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoUser := repository.NewUserDB(db)
	service := NewService(repoUser)

	u, err := service.CreateUser("tung5@gmail.com", "tung", "pass")
	assert.Nil(t, err)
	fmt.Println(u)
}

func Test_GetUser(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoUser := repository.NewUserDB(db)
	service := NewService(repoUser)

	user1, err := service.CreateUser("tung@gmail.com", "tung", "pass")
	assert.Nil(t, err)

	user2, err := service.GetUser(user1.ID)
	fmt.Println(user2)
	assert.Nil(t, err)
	assert.Equal(t, user1.ID, user2.ID)
}

func Test_GetAllUsers(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoUser := repository.NewUserDB(db)
	service := NewService(repoUser)

	listUsers, err := service.ListUsers()
	assert.Nil(t, err)
	for _, user := range listUsers{
		fmt.Println(user)
	}
}

func Test_UpdateUser(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoUser := repository.NewUserDB(db)
	service := NewService(repoUser)

	user1, err := service.CreateUser("tung_TT@gmail.com", "tung", "pass")
	assert.Nil(t, err)

	user1.Name = "tung TT"

	user2, err := service.UpdateUser(user1)
	assert.Nil(t, err)
	fmt.Println(user2)
}

func Test_DeleteUser(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoUser := repository.NewUserDB(db)
	service := NewService(repoUser)

	err = service.DeleteUser(9)
	assert.Nil(t, err)
}
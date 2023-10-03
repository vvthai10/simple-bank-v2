package entry

import (
	"fmt"
	"simple-bank-v2/api"
	repository "simple-bank-v2/infrastructure/repository/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateEntry(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoEntry := repository.NewEntryDB(db)
	service := NewService(repoEntry)

	a, err := service.CreateEntry(2, 10)
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_GetEntry(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoEntry := repository.NewEntryDB(db)
	service := NewService(repoEntry)

	a, err := service.GetEntry(2)
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_ListEntry(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoEntry := repository.NewEntryDB(db)
	service := NewService(repoEntry)

	listEntries, err := service.ListEntry()
	assert.Nil(t, err)
	fmt.Println(listEntries)
}
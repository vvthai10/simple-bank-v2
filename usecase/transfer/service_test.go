package transfer

import (
	"fmt"
	"simple-bank-v2/api"
	"simple-bank-v2/entity"
	repository "simple-bank-v2/infrastructure/repository/gorm"
	"simple-bank-v2/usecase/account"
	"simple-bank-v2/usecase/entry"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateTransfer(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoTransfer := repository.NewTransferDB(db)
	service := NewService(repoTransfer)

	a, err := service.CreateTransfer(2, 2, 10)
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_GetTransfer(t *testing.T) {
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoTransfer := repository.NewTransferDB(db)
	service := NewService(repoTransfer)

	a, err := service.GetTransfer(2)
	assert.Nil(t, err)
	fmt.Println(a)
}

func Test_ListTransfer(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoTransfer := repository.NewTransferDB(db)
	service := NewService(repoTransfer)

	listEntries, err := service.ListTransfer()
	assert.Nil(t, err)
	fmt.Println(listEntries)
}

func Test_TransferTx(t *testing.T){
	db, err := api.NewDatabaseByGORM()
	assert.Nil(t, err)

	repoAccount := repository.NewAccountDB(db)
	serviceAccount := account.NewService(repoAccount)
	repoEntry := repository.NewEntryDB(db)
	serviceEntry := entry.NewService(repoEntry)
	repoTransfer := repository.NewTransferDB(db)
	service := NewService(repoTransfer)

	account1, err := serviceAccount.GetAccount(2)
	assert.Nil(t, err)
	assert.NotNil(t, account1)

	account2, err := serviceAccount.GetAccount(4)
	assert.Nil(t, err)
	assert.NotNil(t, account2)

	fmt.Println(">> before:", account1.Balance, account2.Balance)


	n := 5
	amount := int64(10)
	errs := make(chan error)
	results := make(chan entity.TransferTxResult)
	existed := make(map[int]bool)


	for i :=0; i < n; i++ {
		go func(){
			result, err := service.TransferTx(entity.TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: int64(amount),
			})

			errs <- err
			results  <- result
		}()
	}

	for i := 0; i < n; i++{
		err := <- errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t,transfer)

		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = service.GetTransfer(transfer.ID)
		require.NoError(t, err)

		// Check entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = serviceEntry.GetEntry(fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = serviceEntry.GetEntry(toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)
		
		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check balances
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
	
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balances
	updateAccount1, err := serviceAccount.GetAccount( account1.ID)
	require.NoError(t, err)

	updateAccount2, err := serviceAccount.GetAccount( account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updateAccount1.Balance, updateAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updateAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updateAccount2.Balance)
}
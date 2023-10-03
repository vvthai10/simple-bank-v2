package entity

import "time"

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct{
	Transfer *Transfer `json:"transfer"`
	FromAccount *Account `json:"from_account"`
	ToAccount *Account `json:"to_account"`
	FromEntry *Entry `json:"from_entry"`
	ToEntry *Entry `json:"to_entry"`
}

func NewTransfer(FromAccountID int64, ToAccountID int64, Amount int64) *Transfer{
	return &Transfer{
		FromAccountID: FromAccountID,
		ToAccountID: ToAccountID,
		Amount: Amount,
	}
}

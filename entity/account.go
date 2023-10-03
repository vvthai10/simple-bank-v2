package entity

import "time"

type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Balance   int64     `json:"balance"`
	Bank      string    `json:"bank"`
	CreatedAt time.Time `json:"create_at"`
}

func NewAccount(userID int64, bank string) (*Account){
	return &Account{
		UserID: userID,
		Balance: 0,
		Bank: bank,
	}
}
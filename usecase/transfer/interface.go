package transfer

import "simple-bank-v2/entity"

type Reader interface {
	Get(id int64) (*entity.Transfer, error)
	List() ([]*entity.Transfer, error)
}

type Writer interface{
	Create(e *entity.Transfer) (*entity.Transfer, error)
	TransferTx(params entity.TransferTxParams) (entity.TransferTxResult, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface{
	GetTransfer(id int64) (*entity.Transfer, error)
	ListTransfer() ([]*entity.Transfer, error)
	CreateTransfer(FromAccountID int64, ToAccountID int64, Amount int64) (*entity.Transfer, error)

	TransferTx(params entity.TransferTxParams) (entity.TransferTxResult, error)
}
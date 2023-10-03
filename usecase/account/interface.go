package account

import "simple-bank-v2/entity"

type Reader interface {
	Get(id int64) (*entity.Account, error)
	GetForUpdate(id int64) (*entity.Account, error)
	// List() ([]*entity.Account, error)
}

type Writer interface{
	Create(e *entity.Account) (*entity.Account, error)
	Update(e *entity.Account) (*entity.Account, error)
	Delete(id int64) (error)

	AddAccountBalance(id int64, amount int64) (*entity.Account, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface{
	GetAccount(id int64) (*entity.Account, error)
	CreateAccount(userId int64, bank string) (*entity.Account, error)
	UpdateAccount(e *entity.Account) (*entity.Account, error)
	DeleteAccount(id int64) error
}
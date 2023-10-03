package entry

import "simple-bank-v2/entity"

type Reader interface {
	Get(id int64) (*entity.Entry, error)
	List() ([]*entity.Entry, error)
}

type Writer interface{
	Create(e *entity.Entry) (*entity.Entry, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface{
	GetEntry(id int64) (*entity.Entry, error)
	ListEntry() ([]*entity.Entry, error)
	CreateEntry(AccountID int64, Amount int64) (*entity.Entry, error)
}
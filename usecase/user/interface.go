package user

import "simple-bank-v2/entity"

type Reader interface {
	Get(id int64) (*entity.User, error)
	List() ([]*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (*entity.User, error)
	Update(e *entity.User) (*entity.User, error)
	Delete(id int64) (error)
}

type Repository interface{
	Reader
	Writer
}

type UseCase interface {
	GetUser(id int64) (*entity.User, error)
	ListUsers()([]*entity.User, error)
	CreateUser(email, name , password string)(*entity.User, error)
	UpdateUser(e *entity.User) (*entity.User, error)
	DeleteUser(id int64) (error)
}
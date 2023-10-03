package user

import (
	"simple-bank-v2/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// ==================== ===========================
func (s *Service) CreateUser(email, password, name string) (*entity.User, error) {
	e, err := entity.NewUser(email, password, name)
	if err != nil{
		return nil, err
	}
	
	return s.repo.Create(e)
}

func (s *Service) GetUser(id int64) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

func (s *Service) UpdateUser(e *entity.User) (*entity.User, error) {
	return s.repo.Update(e)
}
func (s *Service) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
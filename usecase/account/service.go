package account

import "simple-bank-v2/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// TODO Overwrite các use-case
func (s *Service) CreateAccount(userID int64, bank string) (*entity.Account, error) {
	e := entity.NewAccount(userID, bank)

	return s.repo.Create(e)
}
func (s *Service) GetAccount(id int64) (*entity.Account, error) {
	return s.repo.Get(id)
}
func (s *Service) UpdateAccount(e *entity.Account) (*entity.Account, error) {
	return s.repo.Update(e)
}
func (s *Service) DeleteAccount(id int64) error {
	return s.repo.Delete(id)
}
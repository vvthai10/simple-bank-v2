package entry

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
func (s *Service) CreateEntry(AccountID int64, Amount int64) (*entity.Entry, error) {
	e := entity.NewEntry(AccountID, Amount)

	return s.repo.Create(e)
}
func (s *Service) GetEntry(id int64) (*entity.Entry, error) {
	return s.repo.Get(id)
}
func (s *Service) ListEntry() ([]*entity.Entry, error) {
	return s.repo.List()
}
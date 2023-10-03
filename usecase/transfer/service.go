package transfer

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
func (s *Service) CreateTransfer(FromAccountID, ToAccountID, Amount int64) (*entity.Transfer, error) {
	e := entity.NewTransfer(FromAccountID, ToAccountID, Amount)

	return s.repo.Create(e)
}
func (s *Service) GetTransfer(id int64) (*entity.Transfer, error) {
	return s.repo.Get(id)
}
func (s *Service) ListTransfer() ([]*entity.Transfer, error) {
	return s.repo.List()
}


// TODO Create transactions
func (s *Service) TransferTx(params entity.TransferTxParams) (entity.TransferTxResult, error){
	return s.repo.TransferTx(params)
}
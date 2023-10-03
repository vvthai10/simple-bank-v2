package repository

import (
	"simple-bank-v2/entity"

	"gorm.io/gorm"
)

type TransferDB struct {
	db *gorm.DB
}

func NewTransferDB(db *gorm.DB) *TransferDB{
	return &TransferDB{
		db: db,
	}
}

// TODO Overwrite repository
func (r *TransferDB) Create(e *entity.Transfer) (*entity.Transfer, error){
	err := r.db.Create(&e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}
func (r *TransferDB) Get(id int64) (*entity.Transfer, error){
	Transfer := &entity.Transfer{}

	err := r.db.Where("id = ?", id).First(Transfer).Error
	if err != nil{
		return nil, err
	}
	return Transfer, nil
}

func (r *TransferDB) List() ([]*entity.Transfer, error){
	transfers := []*entity.Transfer{}

	err := r.db.Find(&transfers).Error
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (r *TransferDB) execTx(fn func(*gorm.DB) error) error{
	tx := r.db.Begin()
	err := fn(tx)
	if err != nil{
		tx.Rollback()
     	return err
	}
	return tx.Commit().Error
}

func (r *TransferDB) TransferTx(params entity.TransferTxParams) (entity.TransferTxResult, error){
	var result entity.TransferTxResult
	
	err := r.execTx(func(tx *gorm.DB) error{
		// TODO init each entity DB
		entryDB := NewEntryDB(tx)
		transferDB := NewTransferDB(tx)

		var err error

		result.Transfer, err = transferDB.Create(&entity.Transfer{
			FromAccountID: params.FromAccountID,
			ToAccountID: params.ToAccountID,
			Amount: params.Amount,
		})
		if err != nil{
			return err
		}

		result.FromEntry, err = entryDB.Create(&entity.Entry{
			AccountID: params.FromAccountID,
			Amount: -params.Amount,
		})
		if err != nil{
			return err
		}

		result.ToEntry, err = entryDB.Create(&entity.Entry{
			AccountID: params.ToAccountID,
			Amount: params.Amount,
		})
		if err != nil{
			return err
		}

		// TODO get account and update its balance
		if params.FromAccountID < params.ToAccountID{
			result.FromAccount, result.ToAccount, err = addMoney(tx, params.FromAccountID, -params.Amount, params.ToAccountID, params.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(tx, params.ToAccountID, params.Amount, params.FromAccountID, -params.Amount)
		}

		if err != nil{
			return err
		}

		return nil
	})

	return result, err
}

func addMoney(tx *gorm.DB,
			accountID1, amount1, accountID2, amount2 int64)(account1, account2 *entity.Account, err error){
	accountDB := NewAccountDB(tx)
	account1, err = accountDB.AddAccountBalance(accountID1, amount1)
	if err != nil {
		return 
	}	

	account2, err = accountDB.AddAccountBalance(accountID2, amount2)
	if err != nil {
		return 
	}	
	return account1, account2, err
}
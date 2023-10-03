package repository

import (
	"simple-bank-v2/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountDB struct {
	db *gorm.DB
}

func NewAccountDB(db *gorm.DB) *AccountDB{
	return &AccountDB{
		db: db,
	}
}

// TODO Overwrite repository
func (r *AccountDB) AddAccountBalance(id int64, amount int64) (*entity.Account, error){
	account, err := r.GetForUpdate(id)
	if err != nil{
		return nil, err
	}
	account.Balance += amount
	accountUpdate, err := r.Update(account)
	if err != nil{
		return nil, err
	}

	return accountUpdate, nil
}

func (r *AccountDB) Create(e *entity.Account) (*entity.Account, error){
	err := r.db.Create(&e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}
func (r *AccountDB) Get(id int64) (*entity.Account, error){
	account := &entity.Account{}

	err := r.db.Where("id = ?", id).First(account).Error
	if err != nil{
		return nil, err
	}
	return account, nil
}
func (r *AccountDB) GetForUpdate(id int64) (*entity.Account, error){
	account := &entity.Account{}

	err := r.db.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Where("id = ?", id).First(account).Error
	if err != nil{
		return nil, err
	}
	return account, nil
}
func (r *AccountDB) Update(e *entity.Account) (*entity.Account, error){
	err := r.db.Where("id = ?", e.ID).Updates(e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}
func (r *AccountDB) Delete(id int64) (error){
	err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil{
		return err
	}
	return nil
}
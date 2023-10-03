package repository

import (
	"simple-bank-v2/entity"

	"gorm.io/gorm"
)

type EntryDB struct {
	db *gorm.DB
}

func NewEntryDB(db *gorm.DB) *EntryDB{
	return &EntryDB{
		db: db,
	}
}

// TODO Overwrite repository
func (r *EntryDB) Create(e *entity.Entry) (*entity.Entry, error){
	err := r.db.Create(&e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}
func (r *EntryDB) Get(id int64) (*entity.Entry, error){
	Entry := &entity.Entry{}

	err := r.db.Where("id = ?", id).First(Entry).Error
	if err != nil{
		return nil, err
	}
	return Entry, nil
}

func (r *EntryDB) List() ([]*entity.Entry, error){
	entries := []*entity.Entry{}

	err := r.db.Find(&entries).Error
	if err != nil {
		return nil, err
	}

	return entries, nil
}
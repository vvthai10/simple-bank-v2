package repository

import (
	"simple-bank-v2/entity"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB{
	return &UserDB{
		db: db,
	}
}

// ================== ================
func (r *UserDB) Create(e *entity.User) (*entity.User, error){
	err := r.db.Create(&e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}

func (r *UserDB) Get(id int64) (*entity.User, error){
	user := &entity.User{}

	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserDB) List() ([]*entity.User, error){
	users := []*entity.User{}

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserDB) Update(e *entity.User) (*entity.User, error){
	err := r.db.Where("id = ?", e.ID).Updates(e).Error
	if err != nil{
		return nil, err
	}
	return e, nil
}
func (r *UserDB) Delete(id int64) (error){
	err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil{
		return err
	}
	return nil
}
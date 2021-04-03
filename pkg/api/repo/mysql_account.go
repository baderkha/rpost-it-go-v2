package repo

import (
	"gorm.io/gorm"
)

// MYSQLAccountRepo : mysql implementation of the account repository
type MYSQLAccountRepo struct {
	db *gorm.DB
}

func NewMYSQLAccountRepo(db *gorm.DB) *MYSQLAccountRepo {
	return &MYSQLAccountRepo{
		db: db,
	}
}

func (macc *MYSQLAccountRepo) FindById(id string) *Account {
	var account Account
	macc.db.First(&account, "id=?", id)
	return &account
}

func (macc *MYSQLAccountRepo) FindByLikeId(id string) *[]Account {
	var accounts []Account
	macc.db.Where("id LIKE ?", id).Find(&accounts)
	return &accounts
}

func (macc *MYSQLAccountRepo) Create(acc *Account) (*Account, error) {
	err := macc.db.Create(acc).Error
	return acc, err
}

func (macc *MYSQLAccountRepo) Update(acc *Account) (*Account, error) {
	err := macc.db.Model(acc).Updates(acc).Error
	return acc, err
}

func (macc *MYSQLAccountRepo) Delete(id string) error {
	return macc.db.Delete(&Account{}, "id=?", id).Error
}

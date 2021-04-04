package repo

import (
	"time"

	"github.com/jinzhu/copier"
)

// Account : entity that models a user
type Account struct {
	ID          string // short-handle-representation , username
	Name        string
	LastName    string
	Email       string
	Password    string
	DateOfBirth time.Time
}

// AccountView : entity that can be passed back to the front end
type AccountView struct {
	ID          string
	Email       string
	DateOfBirth time.Time
}

// use same table name as accounts
func (a Account) TableName() string {
	return "accounts"
}

type Accounts []Account

// use same table name as accounts
func (a AccountView) TableName() string {
	return "accounts"
}

// GenerateView : generate the client safe view object for an account
func (a *Account) GenerateView() *AccountView {
	var acc AccountView
	_ = copier.Copy(&acc, a)
	return &acc
}

// GenerateView : generate the client safe view object for many accounts
func (a *Accounts) GenerateView() *[]AccountView {
	var accs []AccountView
	_ = copier.Copy(&accs, a)
	return &accs
}

// IAccountRepo : Contract that ensures every repo must satisfy these transactions for account entity
type IAccountRepo interface {
	FindById(id string) *Account
	FindByLikeId(id string) *[]Account
	Create(acc *Account) (*Account, error)
	Update(acc *Account) (*Account, error)
	Delete(id string) error
}

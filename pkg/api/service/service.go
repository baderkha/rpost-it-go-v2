package service

import (
	"rpost-it-go/pkg/api/repo"
	"rpost-it-go/pkg/util"

	"gorm.io/gorm"
)

// IService : A facade to all the sub services
type IService interface {
	// GetAccountById : query for an account with a safe
	GetAccountById(idHandle string) (*repo.AccountView, error)
	// GetAccountByApproximateId : fuzzy search the id
	GetAccountByApproximateId(idHandle string) *[]repo.AccountView
	// CreateAccount : create new account safley and do validations on the user input , todo add roles generations aswell
	CreateAccount(acc *repo.Account) (*repo.AccountView, error)
	// UpdateAccount : update the account fields , via a patch
	UpdateAccount(acc *repo.Account) (*repo.AccountView, error)
	// DeleteAccount : Remove the account completely, this needs to account for removing all posts and comments
	DeleteAccount(id string) error
}

type Service struct {
	accountService Account
}

// new service instance
func New(db *gorm.DB) Service {
	return Service{
		accountService: Account{
			repo: repo.NewMYSQLAccountRepo(db),
			er: serviceErrorTemplate{
				model: "Account",
			},
			hasher: &util.Bcrypt{
				Rounds: 12,
			},
		},
	}
}

// GetAccountById : query for an account with a safe
func (s *Service) GetAccountById(idHandle string) (*repo.AccountView, error) {
	return s.accountService.GetByIdPublic(idHandle)
}

// GetAccountByApproximateId : fuzzy search the id
func (s *Service) GetAccountByApproximateId(idHandle string) *[]repo.AccountView {
	return s.accountService.GetByIdFuzzy(idHandle)
}

// CreateAccount : create new account safley and do validations on the user input , todo add roles generations aswell
func (s *Service) CreateAccount(acc *repo.Account) (*repo.AccountView, error) {
	return s.accountService.Create(acc)
}

// UpdateAccount : update the account fields , via a patch
func (s *Service) UpdateAccount(acc *repo.Account) (*repo.AccountView, error) {
	return s.accountService.Update(acc)
}

// DeleteAccount : Remove the account completely, this needs to account for removing all posts and comments
func (s *Service) DeleteAccount(id string) error {
	return s.accountService.Delete(id)
}

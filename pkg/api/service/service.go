package service

import (
	"rpost-it-go/pkg/api/repo"
	"rpost-it-go/pkg/util/crypto"

	"gorm.io/gorm"
)

// IService : A facade to all the sub services
type IService interface {
	// Accounts

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

	// Communities

	// GetCommunityById : fetch a community by an id
	GetCommunityById(idHandle string) (*repo.Community, error)
	//GetCommunityByApproixmateId : fuzzy search on the id
	GetCommunityByApproixmateId(idHandle string) *[]repo.Community
	//GetCommunitiesByAccountOwnerId : grab the communities by the id of an account owner , will show all the communities they own
	GetCommunitiesByAccountOwnerId(idHandle string) *[]repo.Community
	// CreateCommunity : creates a community
	CreateCommunity(com *repo.Community) (*repo.Community, error)
	// UpdateCommunity : Updates the details in a community
	UpdateCommunity(id string, com *repo.Community) (*repo.Community, error)
	// DeleteCommunity : Deletes the community and all the relavent posts for it
	DeleteCommunity(id string) error
}

type Service struct {
	acc Account
	com Community
}

// new service instance
func New(db *gorm.DB) Service {
	return Service{
		acc: Account{
			repo: repo.NewMYSQLAccountRepo(db),
			er: serviceErrorTemplate{
				model: "Account",
			},
			hasher: &crypto.Bcrypt{
				Rounds: 12,
			},
		},
		com: Community{
			repo: repo.NewMYSQLCommunityRepo(db),
			er: serviceErrorTemplate{
				model: "Community",
			},
		},
	}
}

// GetAccountById : query for an account with a safe
func (s *Service) GetAccountById(idHandle string) (*repo.AccountView, error) {
	return s.acc.GetByIdPublic(idHandle)
}

// GetAccountByApproximateId : fuzzy search the id
func (s *Service) GetAccountByApproximateId(idHandle string) *[]repo.AccountView {
	return s.acc.GetByIdFuzzy(idHandle)
}

// CreateAccount : create new account safley and do validations on the user input , todo add roles generations aswell
func (s *Service) CreateAccount(acc *repo.Account) (*repo.AccountView, error) {
	return s.acc.Create(acc)
}

// UpdateAccount : update the account fields , via a patch
func (s *Service) UpdateAccount(acc *repo.Account) (*repo.AccountView, error) {
	return s.acc.Update(acc)
}

// DeleteAccount : Remove the account completely, this needs to account for removing all posts and comments
func (s *Service) DeleteAccount(id string) error {
	return s.acc.Delete(id)
}

// GetCommunityById : fetch a community by an id
func (s *Service) GetCommunityById(idHandle string) (*repo.Community, error) {
	return s.com.GetById(idHandle)
}

//GetCommunityByApproixmateId : fuzzy search on the id
func (s *Service) GetCommunityByApproixmateId(idHandle string) *[]repo.Community {
	return s.com.GetByApproximateId(idHandle)
}

//GetCommunitiesByAccountOwnerId : grab the communities by the id of an account owner , will show all the communities they own
func (s *Service) GetCommunitiesByAccountOwnerId(idHandle string) *[]repo.Community {
	return s.com.GetByAccountOwnerId(idHandle)
}

// CreateCommunity : creates a community
func (s *Service) CreateCommunity(com *repo.Community) (*repo.Community, error) {
	if com.AccountOwnerId == "" {
		return nil, s.com.er.UserInputError("accountOwnerId", "A community Must be associated with an account")
	}
	if !s.acc.isAccountExist(com.AccountOwnerId) {
		return nil, s.com.er.NotFoundResourceReason("The account you're trying to associate with this community does not exist")
	}
	return s.com.Create(com)
}

// UpdateCommunity : Updates the details in a community
func (s *Service) UpdateCommunity(id string, com *repo.Community) (*repo.Community, error) {
	return s.com.Update(id, com)
}

// DeleteCommunity : Deletes the community and all the relavent posts for it
func (s *Service) DeleteCommunity(id string) error {
	return s.com.Delete(id)
}

package repo

// Community : models a community
type Community struct {
	ID             string // short hand id for the community
	Title          string
	Description    string
	AccountOwnerId string
	About          string
	AccountOwner   *AccountView `gorm:"foreignKey:AccountOwnerId"`
}

// ICommunityRepo : Contract that ensures every repo must satisfy these transactions for Community entity
type ICommunityRepo interface {
	FindById(id string) *Community
	FindByLikeIdInput(input string) *[]Community // does a fuzzy search for stuff that may match the id
	FindByAccountOwnerId(accountId string) *[]Community
	Create(com *Community) (*Community, error)
	Update(com *Community) (*Community, error)
	Delete(id string, accountId string) error
}

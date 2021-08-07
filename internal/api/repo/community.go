package repo

// Community : models a community
type Community struct {
	ID             string // short hand id for the community
	Title          string `gorm:"type:VARCHAR(50)"`
	Description    string `gorm:"type:VARCHAR(255)"`
	AccountOwnerId string `gorm:"type:VARCHAR(40)"`
	About          string `gorm:"type:VARCHAR(50)"`
	AccountOwner   *AccountView
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

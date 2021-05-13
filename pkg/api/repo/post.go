package repo

// Post : entity that models a post in a community
type Post struct {
	ID                 string
	Title              string
	Text               string
	PosterId           string
	Poster             *AccountView `gorm:"foreignKey:PosterId"`
	CommunityId        string
	CommunityPostOwner Community `gorm:"foreignKey:CommunityId"`
}

// IPostRepo : Contract that provides the calls a service can consume for the posts entity
type IPostRepo interface {
	// FindById Get 1 post by the primary key
	FindById(id string) *Post
	// FindPostsByAccountOwner Get Posts by the account id
	FindPostsByAccountOwner(posterId string) *[]Post
	// FindPostsByCommunityId Get the posts by a commiunity
	FindPostsByCommunityId(communityId string) *[]Post
	// FindPostsByCommunityIdAndAccountOwner sGet the posts by community and account owner
	FindPostsByCommunityIdAndAccountOwner(communityId string, posterId string) *[]Post
	//  Create Create a post
	Create(post *Post) (*Post, bool)
	// Update Update a post
	Update(post *Post) (*Post, bool)
	// Delete Delete a post
	Delete(id string) bool
}

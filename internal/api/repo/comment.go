package repo

// Comment : entity that models a comment
type Comment struct {
	ID             string `gorm:"type:VARCHAR(50)"`
	Text           string
	AccountOwnerId string `gorm:"type:VARCHAR(50)"`
	PostId         string `gorm:"type:VARCHAR(50)"`
}

// ICommentRepo : Comment repository that models the crud for a comment model , this is a contract
type ICommentRepo interface {
	// FindById : find a comment by a sepcific id
	FindById(id string) *Comment
	// FindPostsByAccountOwner : find comment by account owner
	FindCommentsByAccountOwnerAndPost(accountId string, posterId string) *[]Comment
	// Get comments by post id
	FindCommentsByPostId(postId string) *[]Comment
	// Create : Creates a comment
	Create(comment *Comment) (*Comment, bool)
	// Update : updates a comment
	Update(comment *Comment) (*Comment, bool)
	// DeleteByIdAndAccountId : Deletes a comment by id , this is a perma delete
	DeleteByIdAndAccountId(id string, accountId string) bool
}

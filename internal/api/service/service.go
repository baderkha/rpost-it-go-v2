package service

import (
	"rpost-it-go/internal/api/repo"
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

	// posts

	// GetPostById : fetch a specific post
	GetPostById(id string) (*repo.Post, error)
	// GetPosts : get many posts , by a bunch of filtering options
	GetPosts(req *PostRequest) (*[]repo.Post, error)
	// CreatePost : createa  new post for an account
	CreatePost(request *PostCreateRequest) (*repo.Post, error)
	// UpdatePost : Update a post by an account
	UpdatePost(request *PostUpdateRequest) (*repo.Post, error)
	// DeletePost : Delete a post by an account
	DeletePost(req *PostDeleteRequest) error

	// comments

	// GetCommentById : Fetch a specific comment
	GetCommentById(id string) (*repo.Comment, error)
	// GetCommentsByPostId : get comments relavent to the post
	GetCommentsByPostId(postId string) (*[]repo.Comment, error)
	// CreateComment : Create a comment
	CreateComment(cr *CreateCommentRequest) (*repo.Comment, error)
	// UpdateComment : Update a comment
	UpdateComment(ucr *UpdateCommentRequest) (*repo.Comment, error)
	// DeleteComment : Delete a specific comment
	DeleteComment(dcr *DeletecommentRequest) error

	// Session

	// Login : login and authenticate the client via adding session
	LoginSession(loginCreds *AccountLoginJSON) (*repo.Session, error)
	// Logout : invalidates the stored session , so we forget that user
	LogoutSession(sessionId string)
	// VerifySession : verifies session
	VerifySession(sessionId string) (*repo.Session, error)
}

type Service struct {
	acc     Account
	com     Community
	comment Comment
	post    Post
	session session
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
		post: Post{
			BaseService: BaseService{
				er: serviceErrorTemplate{
					model: "Post",
				},
			},
			repo: repo.NewMYSQLPostRepo(db),
		},
		comment: Comment{
			BaseService: BaseService{
				er: serviceErrorTemplate{
					model: "Comment",
				},
			},
			repo: repo.NewMYSQLCommentRepo(db),
		},
		session: session{
			repo: repo.NewMYSQLSessionRepo(db),
			BaseService: BaseService{
				er: serviceErrorTemplate{
					model: "Session",
				},
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

// GetPostById : fetch a specific post
func (s *Service) GetPostById(id string) (*repo.Post, error) {
	return s.post.GetPostById(id)
}

// GetPosts : get many posts , by a bunch of filtering options
func (s *Service) GetPosts(req *PostRequest) (*[]repo.Post, error) {
	return s.post.GetPosts(req)
}

// CreatePost : createa  new post for an account
func (s *Service) CreatePost(request *PostCreateRequest) (*repo.Post, error) {
	_, err := s.com.GetById(request.Record.CommunityId)
	if err != nil {
		return nil, err
	}
	return s.post.CreatePost(request)
}

// UpdatePost : Update a post by an account
func (s *Service) UpdatePost(request *PostUpdateRequest) (*repo.Post, error) {
	return s.post.UpdatePost(request)
}

// DeletePost : Delete a post by an account
func (s *Service) DeletePost(req *PostDeleteRequest) error {
	return s.post.DeletePost(req)
}

// GetCommentById : get a specific comment by id
func (s *Service) GetCommentById(id string) (*repo.Comment, error) {
	return s.comment.GetCommentById(id)
}

// GetCommentsByPostId : get comments relavent to the post
func (s *Service) GetCommentsByPostId(postId string) (*[]repo.Comment, error) {
	return s.comment.GetCommentsByPostId(postId)
}

// CreateComment : Create a comment
func (s *Service) CreateComment(cr *CreateCommentRequest) (*repo.Comment, error) {
	_, err := s.post.GetPostById(cr.PostId)
	if err != nil {
		return nil, err
	}
	return s.comment.CreateComment(cr)
}

// UpdateComment : Update a comment
func (s *Service) UpdateComment(ucr *UpdateCommentRequest) (*repo.Comment, error) {
	return s.comment.UpdateComment(ucr)
}

// DeleteComment : Delete a specific comment
func (s *Service) DeleteComment(dcr *DeletecommentRequest) error {
	return s.comment.DeleteComment(dcr)
}

// Login : login and authenticate the client via adding session
func (s *Service) LoginSession(loginCreds *AccountLoginJSON) (*repo.Session, error) {
	account, err := s.acc.Authenticate(loginCreds)
	if err != nil {
		return nil, err
	}
	ses, err := s.session.create(account.ID)
	if err != nil {
		return nil, err
	}
	return ses, nil
}

func (s *Service) LogoutSession(sessionId string) {
	_ = s.session.delete(sessionId)
}

func (s *Service) VerifySession(sessionId string) (*repo.Session, error) {
	return s.session.get(sessionId)
}

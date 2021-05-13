package service

import (
	"rpost-it-go/pkg/api/repo"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/copier"
)

// PostRequest : Request that models rquests for many posts
type PostRequest struct {
	AccountId   string `json:"accountId"`
	CommunityId string `json:"communityId"`
}

// PostUpdateRequest : Update request for posts
type PostUpdateRequest struct {
	AccountId string `json:"accountId"`
	PostId    string `json:"id"`
	Record    *PostUpdateJSON
}

// PostCreateRequest : Create request for posts
type PostCreateRequest struct {
	AccountId string `json:"accountId"`
	Record    *PostCreateJSON
}

// PostDeleteRequest : Delete request for posts
type PostDeleteRequest struct {
	AccountId string `json:"accountId"`
	PostId    string `json:"postId"`
}

// PostCreateJSON : post creation json structure for creating a record
type PostCreateJSON struct {
	Text        string
	Title       string
	CommunityId string
}

// PostUpdateJSON : update allowed for Posts
type PostUpdateJSON struct {
	Text  string
	Title string
}

// ConvertToPostModel : converts to the model
func (pcj *PostCreateJSON) ConvertToPostModel() *repo.Post {
	var postModel repo.Post
	copier.Copy(&postModel, pcj)
	return &postModel
}

// ConvertToPostModel : converts to the model
func (puj *PostUpdateJSON) ConvertToPostModel() *repo.Post {
	var postModel repo.Post
	copier.Copy(&postModel, puj)
	return &postModel
}

type Post struct {
	repo repo.IPostRepo
	BaseService
}

func (p *Post) GetPostById(id string) (*repo.Post, error) {
	post := p.repo.FindById(id)
	if post == nil || post.ID == "" {
		return nil, p.Error().NotFoundResourceGeneric()
	}
	return post, nil
}

func (p *Post) GetPosts(req *PostRequest) (*[]repo.Post, error) {
	if req.AccountId != "" && req.CommunityId != "" {
		return p.repo.FindPostsByCommunityIdAndAccountOwner(req.CommunityId, req.AccountId), nil
	} else if req.AccountId != "" && req.CommunityId == "" {
		return p.repo.FindPostsByAccountOwner(req.AccountId), nil
	} else if req.AccountId == "" && req.CommunityId != "" {
		return p.repo.FindPostsByCommunityId(req.CommunityId), nil
	}
	return nil, p.Error().CustomError(400, "Must provide either accountId or communityId or both", "")
}

func (p *Post) CreatePost(request *PostCreateRequest) (*repo.Post, error) {

	err := p.ValidateInput(request, []*validation.FieldRules{
		validation.Field(&request.Record.Text, validation.Required),
		validation.Field(&request.Record.Title, validation.Required),
		validation.Field(&request.Record.CommunityId, validation.Required),
		validation.Field(&request.AccountId, validation.Required),
	})

	if err != nil {
		return nil, err
	}

	modelRecord := request.Record.ConvertToPostModel()

	createdRecord, isCreated := p.repo.Create(modelRecord)
	if !isCreated {
		return nil, p.Error().InternalError()
	}

	return createdRecord, nil
}

func (p *Post) isPostOwnedByAccountId(id string, accountOwnerId string) bool {
	post := p.repo.FindById(id)
	if post == nil || post.PosterId != accountOwnerId {
		return false
	}
	return true
}

func (p *Post) UpdatePost(request *PostUpdateRequest) (*repo.Post, error) {
	err := p.ValidateInput(request, []*validation.FieldRules{
		validation.Field(&request.Record.Title, validation.Length(0, 250)),
		validation.Field(&request.AccountId, validation.Required),
		validation.Field(&request.PostId, validation.Required),
	})
	if err != nil {
		return nil, err
	}

	if !p.isPostOwnedByAccountId(request.PostId, request.AccountId) {
		return nil, p.Error().UnAuthorizedYouDoNotOwnThisResource()
	}

	record := request.Record.ConvertToPostModel()
	record.PosterId = request.AccountId
	record.ID = request.PostId

	updatedRecord, isUpdated := p.repo.Update(record)
	if !isUpdated {
		return nil, p.Error().InternalError()
	}

	return updatedRecord, nil

}

func (p *Post) DeletePost(req *PostDeleteRequest) error {
	err := p.ValidateInput(req, []*validation.FieldRules{
		validation.Field(&req.AccountId, validation.Required),
		validation.Field(&req.PostId, validation.Required),
	})

	if err != nil {
		return err
	}

	if !p.isPostOwnedByAccountId(req.PostId, req.AccountId) {
		return p.Error().UnAuthorizedYouDoNotOwnThisResource()
	}

	if !p.repo.Delete(req.PostId) {
		return p.Error().InternalError()
	}

	return nil
}

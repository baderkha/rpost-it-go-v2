package service

import (
	"rpost-it-go/pkg/api/repo"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/copier"
)

// CreateCommentRequest : used for creating comments, models the request
type CreateCommentRequest struct {
	AccountId string `json:"accountId"`
	PostId    string `json:"postId"`
	Comment   *CreateCommentJSON
}

// CreateCommentJSON : json body for the create comment request
type CreateCommentJSON struct {
	Text string `json:"text"`
}

// UpdateCommentJSON : json body for the update comment request
type UpdateCommentJSON struct {
	Text string `json:"text"`
}

// UpdateCommentRequest : used for updating comments , models the request
type UpdateCommentRequest struct {
	AccountId string `json:"accountId"`
	CommentId string `json:"commentId"`
	Comment   *UpdateCommentJSON
}

// DeletecommentRequest : used for deleting comments , models the request
type DeletecommentRequest struct {
	AccountId string `json:"accountId"`
	CommentId string `json:"commentId"`
}

func (ccj *CreateCommentJSON) ConvertToCommentModel() *repo.Comment {
	var commentModel repo.Comment
	copier.Copy(&commentModel, ccj)
	return &commentModel
}

func (ucj *UpdateCommentJSON) ConvertToCommentModel() *repo.Comment {
	var commentModel repo.Comment
	copier.Copy(&commentModel, ucj)
	return &commentModel
}

// Comment : Comment service used to interface with the persistence layer
type Comment struct {
	repo repo.ICommentRepo
	BaseService
}

func (c *Comment) GetCommentById(id string) (*repo.Comment, error) {
	if id == "" {
		return nil, c.Error().UserInputError("id", "expecting id")
	}
	recordInput := c.repo.FindById(id)
	if recordInput != nil && recordInput.ID == id {
		return recordInput, nil
	}
	return nil, c.Error().NotFoundResourceGeneric()
}

func (c *Comment) GetCommentsByPostId(postId string) (*[]repo.Comment, error) {
	if postId == "" {
		return nil, c.Error().UserInputError("postId", "expected post id")
	}
	return c.repo.FindCommentsByPostId(postId), nil
}

func (c *Comment) CreateComment(cr *CreateCommentRequest) (*repo.Comment, error) {
	recordInput := cr.Comment
	err := c.ValidateInput(recordInput, []*validation.FieldRules{
		validation.Field(&recordInput.Text, validation.Required),
	})
	if err != nil {
		return nil, err
	}
	err = c.ValidateInput(cr, []*validation.FieldRules{
		validation.Field(&cr.AccountId, validation.Required),
		validation.Field(&cr.PostId, validation.Required),
	})
	if err != nil {
		return nil, err
	}

	// safe after checks
	record := recordInput.ConvertToCommentModel()
	record.AccountOwnerId = cr.AccountId
	record.PostId = cr.PostId

	// ok now create record
	createdRecord, isCreated := c.repo.Create(record)
	if !isCreated {
		return nil, c.Error().InternalError()
	}
	return createdRecord, nil
}

func (c *Comment) UpdateComment(ucr *UpdateCommentRequest) (*repo.Comment, error) {
	recordInput := ucr.Comment

	err := c.ValidateInput(recordInput, []*validation.FieldRules{
		validation.Field(&recordInput.Text, validation.Required),
	})
	if err != nil {
		return nil, err
	}
	err = c.ValidateInput(ucr, []*validation.FieldRules{
		validation.Field(&ucr.AccountId, validation.Required),
		validation.Field(ucr.CommentId, validation.Required),
	})
	if err != nil {
		return nil, err
	}

	record := recordInput.ConvertToCommentModel()

	// safe after validation
	record.ID = ucr.CommentId
	record.AccountOwnerId = ucr.AccountId

	// ok now update
	updatedRecord, isUpdated := c.repo.Update(record)
	if !isUpdated {
		return nil, c.Error().InternalError()
	}

	return updatedRecord, nil

}

func (c *Comment) DeleteComment(dcr *DeletecommentRequest) error {
	err := c.ValidateInput(dcr, []*validation.FieldRules{
		validation.Field(&dcr.AccountId, validation.Required),
		validation.Field(&dcr.CommentId, validation.Required),
	})
	if err != nil {
		return err
	}
	// check that the comment belongs to user
	safeComment := c.repo.FindById(dcr.CommentId)
	if safeComment == nil || safeComment.AccountOwnerId != dcr.AccountId {
		return c.Error().UnAuthorizedYouDoNotOwnThisResource()
	}
	// ok safe delete
	isDeleted := c.repo.DeleteByIdAndAccountId(dcr.CommentId, dcr.AccountId)
	if !isDeleted {
		return c.Error().InternalError()
	}
	return nil
}

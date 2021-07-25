package service

import (
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/pkg/util/regex"

	"github.com/davecgh/go-spew/spew"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Community struct {
	repo repo.ICommunityRepo
	er   serviceErrorTemplate
}

func (c *Community) validateCreate(com *repo.Community) error {
	spew.Dump(com)
	err := validation.ValidateStruct(
		com,
		validation.Field(&com.ID, validation.Required, validation.Length(5, 40)),
		validation.Field(&com.AccountOwnerId, validation.Required),
		validation.Field(&com.AccountOwner, validation.NilOrNotEmpty),
		validation.Field(&com.Title, validation.Required, validation.Match(regex.ByName("alphanumeric_wout_spaces"))),
	)
	if err != nil {
		return c.er.CustomError(400, err.Error())
	}
	return nil
}

func (c *Community) validateUpdate(com *repo.Community) error {
	err := validation.ValidateStruct(
		com,
		validation.Field(com.ID, validation.NilOrNotEmpty),
		validation.Field(com.AccountOwnerId, validation.NilOrNotEmpty),
		validation.Field(com.AccountOwner, validation.NilOrNotEmpty),
		validation.Field(com.Title, validation.Required, validation.Match(regex.ByName("alphanumeric_wout_spaces"))),
	)
	if err != nil {
		return c.er.CustomError(400, err.Error())
	}
	return nil
}

func (c *Community) GetById(id string) (*repo.Community, error) {
	if id == "" {
		return nil, c.er.UserInputError("id", "missing")
	}
	com := c.repo.FindById(id)
	if com == nil || com.ID != id {
		return nil, c.er.NotFoundResourceGeneric()
	}
	return com, nil
}

func (c *Community) GetByApproximateId(id string) *[]repo.Community {
	return c.repo.FindByLikeIdInput(id)
}

func (c *Community) GetByAccountOwnerId(accountId string) *[]repo.Community {
	return c.repo.FindByAccountOwnerId(accountId)
}

func (c *Community) Create(com *repo.Community) (*repo.Community, error) {
	existing := c.repo.FindById(com.ID)
	if existing != nil && existing.ID == com.ID {
		return nil, c.er.UserInputError("id", "this community id / handle already exists")
	}
	err := c.validateCreate(com)
	if err != nil {
		return nil, err
	}
	com, err = c.repo.Create(com)
	if err != nil {
		return nil, c.er.InternalError()
	}
	return com, nil
}

func (c *Community) Update(id string, com *repo.Community) (*repo.Community, error) {
	err := c.validateUpdate(com)
	if err != nil {
		return nil, err
	}
	com.ID = id
	com, err = c.repo.Update(com)
	if err != nil {
		return nil, c.er.InternalError()
	}
	return com, nil
}

func (c *Community) Delete(id string) error {
	err := c.repo.Delete(id)
	if err != nil {
		return c.er.InternalError()
	}
	return nil
}

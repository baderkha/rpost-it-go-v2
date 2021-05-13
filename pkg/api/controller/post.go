package controller

import (
	"rpost-it-go/pkg/api/service"

	"github.com/gofiber/fiber"
)

type Post struct {
	Basecontroller
}

func (p *Post) GetById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	record, err := p.Service().GetPostById(id)
	if err != nil {
		p.StatusFromError(ctx, err)
		return
	}
	p.OK(ctx, record)
}

func (p *Post) GetAll(ctx *fiber.Ctx) {
	req := &service.PostRequest{
		AccountId:   ctx.Params("id"),
		CommunityId: ctx.Params("community-id"),
	}

	records, err := p.Service().GetPosts(req)
	if err != nil {
		p.StatusFromError(ctx, err)
		return
	}
	p.OK(ctx, records)
}

func (p *Post) Create(ctx *fiber.Ctx) {
	var record service.PostCreateJSON

	err := ctx.BodyParser(record)
	if err != nil {
		p.BodyCouldNotParse(ctx, err)
		return
	}

	request := service.PostCreateRequest{
		AccountId: ctx.Params("account-id"),
		Record:    &record,
	}

	createdPost, err := p.Service().CreatePost(&request)
	if err != nil {
		p.StatusFromError(ctx, err)
		return
	}

	p.Created(ctx, createdPost)
}

func (p *Post) Update(ctx *fiber.Ctx) {
	var record service.PostUpdateJSON
	err := ctx.BodyParser(record)
	if err != nil {
		p.BodyCouldNotParse(ctx, err)
		return
	}

	request := &service.PostUpdateRequest{
		AccountId: ctx.Params("account-id"),
		PostId:    ctx.Params("id"),
		Record:    &record,
	}
	updatedRecord, err := p.Service().UpdatePost(request)
	if err != nil {
		p.StatusFromError(ctx, err)
		return
	}
	p.Updated(ctx, updatedRecord)
}

func (p *Post) Delete(ctx *fiber.Ctx) {

	request := &service.PostDeleteRequest{
		AccountId: ctx.Params("account-id"),
		PostId:    ctx.Params("id"),
	}
	err := p.Service().DeletePost(request)
	if err != nil {
		p.StatusFromError(ctx, err)
		return
	}
	p.Deleted(ctx, "post")
}

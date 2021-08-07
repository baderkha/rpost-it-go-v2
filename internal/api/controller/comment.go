package controller

import (
	"rpost-it-go/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	Basecontroller
}

func (c *Comment) GetAllCommentsForPostById(ctx *fiber.Ctx) {
	postId := ctx.Params("id")
	comments, err := c.Service().GetCommentsByPostId(postId)
	if err != nil {
		c.StatusFromError(ctx, err)
		return
	}
	c.OK(ctx, comments)
}

func (c *Comment) GetCommentId(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	comment, err := c.Service().GetCommentById(id)
	if err != nil {
		c.StatusFromError(ctx, err)
		return
	}
	c.OK(ctx, comment)
}

func (c *Comment) CreateComment(ctx *fiber.Ctx) {
	var record service.CreateCommentJSON
	accountId := getAccountId(ctx)
	postId := ctx.Params("id")
	err := ctx.BodyParser(&record)
	if err != nil {
		c.BodyCouldNotParse(ctx, err)
		return
	}
	newComment, err := c.service.CreateComment(&service.CreateCommentRequest{
		AccountId: accountId,
		PostId:    postId,
		Comment:   &record,
	})

	if err != nil {
		c.StatusFromError(ctx, err)
		return
	}
	c.Created(ctx, newComment)
}

func (c *Comment) UpdateComment(ctx *fiber.Ctx) {
	var record service.UpdateCommentJSON
	accountId := getAccountId(ctx)
	commentId := ctx.Query("comment-id")
	err := ctx.BodyParser(&record)
	if err != nil {
		c.BodyCouldNotParse(ctx, err)
		return
	}
	newComment, err := c.service.UpdateComment(&service.UpdateCommentRequest{
		AccountId: accountId,
		CommentId: commentId,
		Comment:   &record,
	})

	if err != nil {
		c.StatusFromError(ctx, err)
		return
	}
	c.Created(ctx, newComment)
}

func (c *Comment) DeleteComment(ctx *fiber.Ctx) {

	err := c.Service().DeleteComment(&service.DeletecommentRequest{
		AccountId: getAccountId(ctx),
		CommentId: ctx.Params("id"),
	})

	if err != nil {
		c.StatusFromError(ctx, err)
		return
	}
	c.Deleted(ctx, "Comment")
}

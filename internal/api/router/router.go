package router

import (
	"rpost-it-go/internal/api/controller"

	"github.com/gofiber/fiber"
)

func GenerateRotues(app *fiber.App, ctrlr controller.App) {

	app.Get("", func(ctx *fiber.Ctx) {
		ctx.Status(200).SendString("api works")
	})

	app.Post("/login", ctrlr.Auth.Login)
	app.Post("/logout", ctrlr.Auth.Logout)
	api := app.Group("/api", ctrlr.Auth.Verify)
	accountRouter := api.Group("/accounts")
	{
		accountRouter.Get("", ctrlr.Account.Search)
		accountRouter.Get("/:id", ctrlr.Account.GetById)
		accountRouter.Get("/:id/communties", ctrlr.Community.GetByAccountId)
		accountRouter.Post("", ctrlr.Account.Create)
		accountRouter.Patch("/:id", ctrlr.Account.Update)
		accountRouter.Delete("/:id", ctrlr.Account.Delete)
	}

	communityRouter := api.Group("/communities")
	{
		communityRouter.Get("", ctrlr.Community.SearchCommunities)
		communityRouter.Get("/:id", ctrlr.Community.GetById)
		communityRouter.Post("", ctrlr.Community.Create)
		communityRouter.Patch("/:id", ctrlr.Community.Update)
		communityRouter.Delete("/:id", ctrlr.Community.Delete)
	}

	postRouter := api.Group("/posts")
	{
		postRouter.Get("", ctrlr.Post.GetAll)
		postRouter.Get("/:id", ctrlr.Post.GetById)
		postRouter.Get("/:id/comments", ctrlr.Comment.GetAllCommentsForPostById)
		postRouter.Post("/:id/comments", ctrlr.Comment.CreateComment)
		postRouter.Post("", ctrlr.Post.Create)
		postRouter.Patch("/:id", ctrlr.Post.Update)
		postRouter.Delete("/:id", ctrlr.Post.Delete)
	}

	commentRouter := api.Group("/comments")
	{
		commentRouter.Get("/:id", ctrlr.Comment.GetCommentId)
		commentRouter.Patch("/:id", ctrlr.Comment.UpdateComment)
		commentRouter.Delete("/:id", ctrlr.Comment.DeleteComment)
	}
}

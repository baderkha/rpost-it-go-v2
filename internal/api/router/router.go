package router

import (
	"rpost-it-go/internal/api/controller"

	"github.com/gofiber/fiber"
)

func GenerateRotues(app *fiber.App, ctrlr controller.App) {

	app.Get("", func(ctx *fiber.Ctx) {
		ctx.Status(200).SendString("api works")
	})

	// public routes
	{
		app.Post("/login", ctrlr.Auth.Login)

		app.Post("/accounts", ctrlr.Account.Create) // have verification ?
		app.Get("/accounts", ctrlr.Account.Search)
		app.Get("/accounts/:id/communties", ctrlr.Community.GetByAccountId)
		app.Get("/accounts/:id", ctrlr.Account.GetById)

		app.Get("/communities", ctrlr.Community.SearchCommunities)
		app.Get("/communities/:id", ctrlr.Community.GetById)

		app.Get("/posts", ctrlr.Post.GetAll)
		app.Get("/posts/:id", ctrlr.Post.GetById)
		app.Get("/posts/:id/comments", ctrlr.Comment.GetAllCommentsForPostById)

		app.Get("/comments/:id", ctrlr.Comment.GetCommentId)

		app.Post("/logout", ctrlr.Auth.Logout)
	}

	// private routes
	{
		api := app.Group("/private", ctrlr.Auth.Verify)

		accountRouter := api.Group("/accounts")
		{
			accountRouter.Patch("/:id", ctrlr.Account.Update)
			accountRouter.Delete("/:id", ctrlr.Account.Delete)
		}

		communityRouter := api.Group("/communities")
		{
			communityRouter.Post("", ctrlr.Community.Create)
			communityRouter.Patch("/:id", ctrlr.Community.Update)
			communityRouter.Delete("/:id", ctrlr.Community.Delete)
		}

		postRouter := api.Group("/posts")
		{

			postRouter.Post("/:id/comments", ctrlr.Comment.CreateComment)
			postRouter.Post("", ctrlr.Post.Create)
			postRouter.Patch("/:id", ctrlr.Post.Update)
			postRouter.Delete("/:id", ctrlr.Post.Delete)
		}

		commentRouter := api.Group("/comments")
		{

			commentRouter.Patch("/:id", ctrlr.Comment.UpdateComment)
			commentRouter.Delete("/:id", ctrlr.Comment.DeleteComment)
		}
	}
}

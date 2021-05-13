package router

import (
	"rpost-it-go/pkg/api/controller"

	"github.com/gofiber/fiber"
)

func GenerateRotues(app *fiber.App, ctrlr controller.App) {

	accountRouter := app.Group("accounts")
	{
		accountRouter.Get("", ctrlr.Account.Search)
		accountRouter.Get("/:id", ctrlr.Account.GetById)
		accountRouter.Get("/:id/communties", ctrlr.Community.GetByAccountId)
		accountRouter.Post("", ctrlr.Account.Create)
		accountRouter.Patch("/:id", ctrlr.Account.Update)
		accountRouter.Delete("/:id", ctrlr.Account.Delete)
	}
	communityRouter := app.Group("communities")
	{
		communityRouter.Get("", ctrlr.Community.SearchCommunities)
		communityRouter.Get("/:id", ctrlr.Community.GetById)
		communityRouter.Post("", ctrlr.Community.Create)
		communityRouter.Patch("/:id", ctrlr.Community.Update)
		communityRouter.Delete("/:id", ctrlr.Community.Delete)
	}

	postRouter := app.Group("posts")
	{
		postRouter.Get("", ctrlr.Post.GetAll)
		postRouter.Get("/:id", ctrlr.Post.GetById)
		postRouter.Post("", ctrlr.Post.Create)
		postRouter.Patch("/:id", ctrlr.Post.Update)
		postRouter.Delete("/:id", ctrlr.Post.Delete)
	}
}

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
		accountRouter.Post("", ctrlr.Account.Create)
		accountRouter.Patch("/:id", ctrlr.Account.Update)
		accountRouter.Delete("/:id", ctrlr.Account.Delete)
	}
}

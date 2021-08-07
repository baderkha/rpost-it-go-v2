package router

import (
	"rpost-it-go/internal/api/controller"
	"rpost-it-go/pkg/thirdparty/swagger"
	"rpost-it-go/pkg/util/http"

	"github.com/gofiber/fiber/v2"
)

func GenerateRotues(app *fiber.App, ctrlr controller.App) {
	swaggerMiddleware := swagger.NewMiddleware("./docs.json", "/")

	app = swaggerMiddleware.Register(app)
	// public routes
	{
		app.Post("/login", http.FiberRouteV2(ctrlr.Auth.Login))

		app.Post("/accounts", http.FiberRouteV2(ctrlr.Account.Create)) // have verification ?
		app.Get("/accounts", http.FiberRouteV2(ctrlr.Account.Search))
		app.Get("/accounts/:id/communties", http.FiberRouteV2(ctrlr.Community.GetByAccountId))
		app.Get("/accounts/:id", http.FiberRouteV2(ctrlr.Account.GetById))

		app.Get("/communities", http.FiberRouteV2(ctrlr.Community.SearchCommunities))
		app.Get("/communities/:id", http.FiberRouteV2(ctrlr.Community.GetById))

		app.Get("/posts", http.FiberRouteV2(ctrlr.Post.GetAll))
		app.Get("/posts/:id", http.FiberRouteV2(ctrlr.Post.GetById))
		app.Get("/posts/:id/comments", http.FiberRouteV2(ctrlr.Comment.GetAllCommentsForPostById))
		app.Get("/verify/:id", http.FiberRouteV2(ctrlr.Account.VerifyCreation))

		app.Get("/comments/:id", http.FiberRouteV2(ctrlr.Comment.GetCommentId))

		app.Post("/logout", http.FiberRouteV2(ctrlr.Auth.Logout))
	}

	// private routes
	{
		api := app.Group("/private", http.FiberRouteV2(ctrlr.Auth.Verify))

		accountRouter := api.Group("/accounts")
		{
			accountRouter.Patch("/:id", http.FiberRouteV2(ctrlr.Account.Update))
			accountRouter.Delete("/:id", http.FiberRouteV2(ctrlr.Account.Delete))
		}

		communityRouter := api.Group("/communities")
		{
			communityRouter.Post("", http.FiberRouteV2(ctrlr.Community.Create))
			communityRouter.Patch("/:id", http.FiberRouteV2(ctrlr.Community.Update))
			communityRouter.Delete("/:id", http.FiberRouteV2(ctrlr.Community.Delete))
		}

		postRouter := api.Group("/posts")
		{
			postRouter.Post("/:id/comments", http.FiberRouteV2(ctrlr.Comment.CreateComment))
			postRouter.Post("", http.FiberRouteV2(ctrlr.Post.Create))
			postRouter.Patch("/:id", http.FiberRouteV2(ctrlr.Post.Update))
			postRouter.Delete("/:id", http.FiberRouteV2(ctrlr.Post.Delete))
		}

		commentRouter := api.Group("/comments")
		{

			commentRouter.Patch("/:id", http.FiberRouteV2(ctrlr.Comment.UpdateComment))
			commentRouter.Delete("/:id", http.FiberRouteV2(ctrlr.Comment.DeleteComment))
		}
	}
}

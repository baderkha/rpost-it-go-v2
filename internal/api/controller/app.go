package controller

import (
	"fmt"
	"rpost-it-go/internal/api/service"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type App struct {
	Account   Account
	Community Community
	Post      Post
	Comment   Comment
	Auth      IAuth
}

func getAccountId(ctx *fiber.Ctx) string {
	accountID := (ctx.Locals("request-account-id")).(string)
	if accountID == "" {
		fmt.Println("account id is empty even though route needed auth , put middle ware on or do something")
		panic("should not happen !")
	}
	return accountID
}

func New(db *gorm.DB) App {
	ser := service.New(db)
	return App{
		Account: Account{
			service: &ser,
		},
		Community: Community{
			service: &ser,
		},
		Post: Post{
			Basecontroller: Basecontroller{
				service: &ser,
			},
		},
		Comment: Comment{
			Basecontroller: Basecontroller{
				service: &ser,
			},
		},
		Auth: &SessionAuth{
			Basecontroller: Basecontroller{
				service: &ser,
			},
			DefaultAuthDurationDays: 7,
			DomainPrefix:            "local.api.postrealm.com",
			IsLocalHost:             true,
		},
	}
}

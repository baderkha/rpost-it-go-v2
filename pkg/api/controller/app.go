package controller

import (
	"rpost-it-go/pkg/api/service"

	"gorm.io/gorm"
)

type App struct {
	Account   Account
	Community Community
	Post      Post
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
	}
}

package controller

import (
	"rpost-it-go/pkg/api/service"

	"gorm.io/gorm"
)

type App struct {
	Account   Account
	Community Community
	Post      Post
	Comment   Comment
	Auth      IAuth
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
			DomainPrefix:            "postrealm.com",
		},
	}
}

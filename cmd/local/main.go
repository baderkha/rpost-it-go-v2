package main

import (
	"rpost-it-go/internal/api"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	app, db := api.SetupDbAndRouter()

	err := app.Listen(":4040")
	spew.Dump(err)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

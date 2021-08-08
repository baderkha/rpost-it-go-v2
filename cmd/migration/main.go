package main

import (
	"rpost-it-go/internal/api"
	"rpost-it-go/internal/api/repo"
)

func main() {
	db := api.NewDb()
	// migrate
	_ = db.AutoMigrate(repo.Account{})
	_ = db.AutoMigrate(repo.Community{})
	_ = db.AutoMigrate(repo.Post{})
	_ = db.AutoMigrate(repo.Session{})

	// close connection once done
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

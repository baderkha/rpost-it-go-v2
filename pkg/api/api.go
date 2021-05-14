package api

import (
	"log"
	"os"
	"rpost-it-go/pkg/api/controller"
	"rpost-it-go/pkg/api/repo"
	"rpost-it-go/pkg/api/router"
	"time"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newDb() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	dsn := "root:root@tcp(127.0.0.1:3320)/rpost-it?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func Start() {
	db := newDb()

	app := fiber.New()

	router.GenerateRotues(app, controller.New(db))
	_ = app.Listen("8000")
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func Migrate() {
	db := newDb()
	// migrate
	db.AutoMigrate(repo.Account{})
	db.AutoMigrate(repo.Community{})
	db.AutoMigrate(repo.Post{})

	// close connection once done
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

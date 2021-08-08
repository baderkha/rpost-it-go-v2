package api

import (
	"log"
	"os"
	"rpost-it-go/internal/api/config"
	"rpost-it-go/internal/api/controller"
	"rpost-it-go/internal/api/router"
	"time"

	"github.com/gofiber/fiber/v2"
	loggerFiber "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func NewDb() *gorm.DB {
	if db == nil {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)

		dsn := config.Get().DBConfig.GetDSN()
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic(err)
		}
		return db
	}

	return db

}

func SetupDbAndRouter() (*fiber.App, *gorm.DB) {
	db := NewDb()

	app := fiber.New()

	router.GenerateRotues(app, controller.New(db))
	app.Use(loggerFiber.New())

	return app, db
}

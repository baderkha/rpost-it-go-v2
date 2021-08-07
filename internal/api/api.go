package api

import (
	"log"
	"os"
	"rpost-it-go/internal/api/config"
	"rpost-it-go/internal/api/controller"
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/internal/api/router"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/gofiber/fiber/v2"
	loggerFiber "github.com/gofiber/fiber/v2/middleware/logger"
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

	dsn := config.Get().DBConfig.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func Start() {
	spew.Dump("here")
	db := newDb()

	app := fiber.New()

	router.GenerateRotues(app, controller.New(db))
	app.Use(loggerFiber.New())
	err := app.Listen(":4040")

	spew.Dump(err)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func Migrate() {
	db := newDb()
	// migrate
	_ = db.AutoMigrate(repo.Account{})
	_ = db.AutoMigrate(repo.Community{})
	_ = db.AutoMigrate(repo.Post{})
	_ = db.AutoMigrate(repo.Session{})

	// close connection once done
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

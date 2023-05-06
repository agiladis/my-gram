package main

import (
	"fmt"
	"log"
	"my-gram/entity"
	"my-gram/router"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_HOST     = os.Getenv("DB_HOST")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_DATABASE = os.Getenv("DB_DATABASE")
	)

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		// Log to see changes
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	db.Debug().Migrator().DropTable(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.Socialmedia{})
	db.Debug().AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.Socialmedia{})

	router.StartServer(db.Debug()).Run(":3000")
}

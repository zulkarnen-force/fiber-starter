package config

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zulkarnen-force/fiber-starter/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Config struct {
	Port string
}

var AppConfig Config


func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
	   log.Fatalf("Error loading .env file: %v", err)
    }

    AppConfig = Config{
	   Port: getEnv("PORT", "3000"),
    }
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=db user=user password=password dbname=starter port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return nil, err
	}

	DB = db

	return db, nil

}

func InitFiber() *fiber.App {
    return fiber.New()
}

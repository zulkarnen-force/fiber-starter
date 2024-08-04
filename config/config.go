package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zulkarnen-force/fiber-starter/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github/instaShop_assessment/features/domain"
)

func NewPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "user=technology dbname=instashop password=gbenga sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Order{}, &domain.OrderItem{})

	fmt.Println("Database connection successful")
	return db, nil
}

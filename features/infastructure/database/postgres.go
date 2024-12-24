package database

import (
	"github.com/jinzhu/gorm"

	"github/instaShop_assessment/features/domain"
)

func NewPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=ecommerce password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Order{}, &domain.OrderItem{})

	return db, nil
}

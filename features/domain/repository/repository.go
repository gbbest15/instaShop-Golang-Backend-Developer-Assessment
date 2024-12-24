package repository

import "github/instaShop_assessment/features/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
}

type ProductRepository interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint) error
}

type OrderRepository interface {
	Create(order *domain.Order) error
	CreateOrderItems(items []domain.OrderItem) error
	GetByUserID(userID uint) ([]domain.Order, error)
	GetByID(id uint) (*domain.Order, error)
	Update(order *domain.Order) error
}

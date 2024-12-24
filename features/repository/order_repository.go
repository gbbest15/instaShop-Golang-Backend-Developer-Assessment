package repository

import (
	"github.com/jinzhu/gorm"

	domain "github/instaShop_assessment/features/domain"
	repo "github/instaShop_assessment/features/domain/repository"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repo.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) CreateOrderItems(items []domain.OrderItem) error {
	return r.db.Create(&items).Error
}

func (r *orderRepository) GetByUserID(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Where("user_id = ?", userID).Preload("Products").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Products").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

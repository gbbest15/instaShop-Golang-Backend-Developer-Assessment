package repository

import (
	"fmt"

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
	if order == nil {
		return fmt.Errorf("order cannot be nil")
	}
	return r.db.Create(order).Error
}

func (r *orderRepository) CreateOrderItems(items []domain.OrderItem) error {
	if len(items) == 0 {
		return fmt.Errorf("no order items to insert")
	}

	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			fmt.Println("Panic occurred during order item creation:", err)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for i := range items {
		if items[i].ProductID == 0 || items[i].Quantity == 0 {
			tx.Rollback()
			return fmt.Errorf("invalid order item: %+v", items[i])
		}
	}

	fmt.Println("Order items before creation:", items)

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

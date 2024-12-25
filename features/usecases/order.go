package usecase

import (
	"errors"
	"fmt"

	domain "github/instaShop_assessment/features/domain"
	repo "github/instaShop_assessment/features/domain/repository"
)

type OrderUseCase struct {
	orderRepo   repo.OrderRepository
	productRepo repo.ProductRepository
}

func NewOrderUseCase(orderRepo repo.OrderRepository, productRepo repo.ProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (u *OrderUseCase) Create(order *domain.Order) error {
	if order == nil {
		return fmt.Errorf("order cannot be nil")
	}

	order.Status = "Pending"

	var total float64

	for i := range order.Products {
		product, err := u.productRepo.GetByID(order.Products[i].ProductID)
		if err != nil {
			return fmt.Errorf("failed to get product: %w", err)
		}

		order.Products[i].Price = product.Price
		total += product.Price * float64(order.Products[i].Quantity)
	}

	order.Total = total

	if err := u.orderRepo.Create(order); err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	for i := range order.Products {
		order.Products[i].OrderID = order.ID
	}

	// Create the order items
	if err := u.orderRepo.CreateOrderItems(order.Products); err != nil {
		return fmt.Errorf("failed to create order items: %w", err)
	}

	return nil
}

func (u *OrderUseCase) GetUserOrders(userID uint) ([]domain.Order, error) {
	return u.orderRepo.GetByUserID(userID)
}

func (u *OrderUseCase) CancelOrder(orderID, userID uint) error {
	order, err := u.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("unauthorized")
	}

	if order.Status != "Pending" {
		return errors.New("only pending orders can be cancelled")
	}

	order.Status = "Cancelled"
	return u.orderRepo.Update(order)
}

func (u *OrderUseCase) UpdateStatus(orderID uint, status string) error {
	order, err := u.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	order.Status = status
	return u.orderRepo.Update(order)
}

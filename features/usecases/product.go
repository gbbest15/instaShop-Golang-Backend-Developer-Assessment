package usecase

import (
	domain "github/instaShop_assessment/features/domain"
	repo "github/instaShop_assessment/features/domain/repository"
)

type ProductUseCase struct {
	repo repo.ProductRepository
}

func NewProductUseCase(repo repo.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (p *ProductUseCase) Create(product *domain.Product) error {
	return p.repo.Create(product)
}

func (p *ProductUseCase) GetAll() ([]domain.Product, error) {
	return p.repo.GetAll()
}

func (p *ProductUseCase) GetByID(id uint) (*domain.Product, error) {
	return p.repo.GetByID(id)
}

func (p *ProductUseCase) Update(product *domain.Product) error {
	return p.repo.Update(product)
}

func (p *ProductUseCase) Delete(id uint) error {
	return p.repo.Delete(id)
}

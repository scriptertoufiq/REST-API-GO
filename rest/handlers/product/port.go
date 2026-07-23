package product

import "ecommerce/domain"

type Service interface {
	Create(p domain.Product) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Count() (int64, error)
	GetByID(productID int) *domain.Product
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}

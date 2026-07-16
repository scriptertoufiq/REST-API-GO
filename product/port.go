package product

import "ecommerce/domain"

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	List() ([]*domain.Product, error)
	GetByID(productID int) *domain.Product
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}

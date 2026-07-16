package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) *service {
	return &service{
		productRepo: productRepo,
	}
}

func (s *service) Create(product domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(product)

}

func (s *service) List() ([]*domain.Product, error) {
	return s.productRepo.List()
}

func (s *service) GetByID(productID int) *domain.Product {
	return s.productRepo.GetByID(productID)
}

func (s *service) Delete(productID int) error {
	return s.productRepo.Delete(productID)
}

func (s *service) Update(product domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(product)
}

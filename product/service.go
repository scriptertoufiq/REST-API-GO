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

func (s *service) List(page, limit int) ([]*domain.Product, error) {
	return s.productRepo.List(page, limit)
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

func (s *service) Count() (int64, error) {
	return s.productRepo.Count()
}

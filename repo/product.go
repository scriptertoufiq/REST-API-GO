package repo

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	GetByID(productID int) *Product
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productsList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) GetByID(productID int) *Product {
	for _, product := range r.productsList {
		if product.ID == productID {
			return product
		}
	}
	return nil
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productsList) + 1
	r.productsList = append(r.productsList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productsList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productsList, nil
}

func (r *productRepo) Delete(productID int) error {
	for i, product := range r.productsList {
		if product.ID == productID {
			r.productsList = append(r.productsList[:i], r.productsList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product not found")
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for i, p := range r.productsList {
		if p.ID == product.ID {
			r.productsList[i] = &product
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is the first product",
		Price:       19.99,
		ImgURL:      "https://example.com/product1.jpg",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is the second product",
		Price:       29.99,
		ImgURL:      "https://example.com/product2.jpg",
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is the third product",
		Price:       39.99,
		ImgURL:      "https://example.com/product3.jpg",
	}

	prd4 := &Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is the fourth product",
		Price:       49.99,
		ImgURL:      "https://example.com/product4.jpg",
	}

	prd5 := &Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is the fifth product",
		Price:       59.99,
		ImgURL:      "https://example.com/product5.jpg",
	}

	r.productsList = []*Product{
		prd1,
		prd2,
		prd3,
		prd4,
		prd5,
	}
}

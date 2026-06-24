package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

var ProductsList []Product

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is the first product",
		Price:       19.99,
		ImgURL:      "https://example.com/product1.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is the second product",
		Price:       29.99,
		ImgURL:      "https://example.com/product2.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is the third product",
		Price:       39.99,
		ImgURL:      "https://example.com/product3.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Product 4",
		Description: "This is the fourth product",
		Price:       49.99,
		ImgURL:      "https://example.com/product4.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Product 5",
		Description: "This is the fifth product",
		Price:       59.99,
		ImgURL:      "https://example.com/product5.jpg",
	}

	ProductsList = []Product{prd1, prd2, prd3, prd4, prd5}
}

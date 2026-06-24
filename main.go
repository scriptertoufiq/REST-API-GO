package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(getProduct))
	mux.Handle("POST /create-product", http.HandlerFunc(addProduct))

	fmt.Println("Server is running on port 8080")
	globalRouter := globalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}

}

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

	productsList = []Product{prd1, prd2, prd3, prd4, prd5}
}

package product

import (
	repo "ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURL:      req.ImgURL,
	})
	if err != nil {
		fmt.Println("Error creating product:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// newProduct.ID = len(database.ProductsList) + 1
	// database.ProductsList = append(database.ProductsList, newProduct)

	util.SendData(w, createdProduct, 201)

}

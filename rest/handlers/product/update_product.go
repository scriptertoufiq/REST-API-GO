package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct database.Product

	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Ensure the ID comes from the URL, not the request body.
	updatedProduct.ID = id

	product := database.Update(updatedProduct)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, product, http.StatusOK)
}

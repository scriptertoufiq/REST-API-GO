package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	updatedProduct.ID = id
	product := database.Update(updatedProduct)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, product, http.StatusOK)
}

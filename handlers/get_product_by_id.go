package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") //
	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for idx, product := range database.ProductsList {
		if product.ID == id {
			util.SendData(w, database.ProductsList[idx], 200)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)

}

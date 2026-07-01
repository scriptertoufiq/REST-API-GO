package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Get(id)
	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w, product, 200)
}

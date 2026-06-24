package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductsList) + 1
	database.ProductsList = append(database.ProductsList, newProduct)

	util.SendData(w, newProduct, 201)

}

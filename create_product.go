package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productsList) + 1
	productsList = append(productsList, newProduct)

	sendData(w, newProduct, 201)

}

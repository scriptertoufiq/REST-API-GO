package main

import (
	"net/http"
)



func getProduct(w http.ResponseWriter, r *http.Request) {

	sendData(w, productsList, 200)

}
package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.List(), 200)

}

package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var loginUser RequestLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginUser)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := database.Find(loginUser.Email, loginUser.Password)

	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	util.SendData(w, user, 200)

}

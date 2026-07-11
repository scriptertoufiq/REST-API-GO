package user

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newUser = newUser.Store()

	util.SendData(w, newUser, 201)

}

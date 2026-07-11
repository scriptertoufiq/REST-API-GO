package user

import (
	"ecommerce/config"
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

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

	cnf := config.GetConfig()

	accessToken, err := util.CreateJWT(cnf.JwtSecretKey, util.Payload{
		Sub:         fmt.Sprintf("%d", user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Error creating JWT", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, 200)

}

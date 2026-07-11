package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := parts[1]
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	jwtSignature := tokenParts[2]

	if jwtHeader == "" || jwtPayload == "" || jwtSignature == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	message := jwtHeader + "." + jwtPayload
	cnf := config.GetConfig()

	byteArrSecret := []byte(cnf.JwtSecretKey)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	hash := h.Sum(nil)
	newSignature := base64UrlEncode(hash)

	if newSignature != jwtSignature {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	newProduct = database.Store(newProduct)

	// newProduct.ID = len(database.ProductsList) + 1
	// database.ProductsList = append(database.ProductsList, newProduct)

	util.SendData(w, newProduct, 201)

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

 func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, World!");
}
func aboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is the about page");
}

type Product struct {
	ID    int `json:"id"`
	Title  string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"img_url"`
}




func getProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productsList)

	
}

func addProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	w.WriteHeader(201)

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
	
}

var productsList [] Product





func main() {	
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloWorld)
	mux.HandleFunc("/about",aboutPage)
	mux.HandleFunc("/product",getProduct)
	mux.HandleFunc("/create-product",addProduct)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)


	if err != nil {
		fmt.Println("Error starting server:", err)
	}else {		
		fmt.Println("Server started successfully")
	}




}

func init() {
	prd1 := Product{
		ID: 1,
		Title: "Product 1",
		Description: "This is the first product",
		Price: 19.99,
		ImgURL: "https://example.com/product1.jpg",
	}
	prd2 := Product{
		ID: 2,
		Title: "Product 2",
		Description: "This is the second product",
		Price: 29.99,
		ImgURL: "https://example.com/product2.jpg",
	}
	prd3 := Product{
		ID: 3,
		Title: "Product 3",
		Description: "This is the third product",
		Price: 39.99,
		ImgURL: "https://example.com/product3.jpg",
	}
	prd4 := Product{
		ID: 4,
		Title: "Product 4",
		Description: "This is the fourth product",
		Price: 49.99,
		ImgURL: "https://example.com/product4.jpg",
	}
	prd5 := Product{
		ID: 5,
		Title: "Product 5",
		Description: "This is the fifth product",
		Price: 59.99,
		ImgURL: "https://example.com/product5.jpg",
	}

	productsList = []Product{prd1, prd2, prd3, prd4, prd5}		
}
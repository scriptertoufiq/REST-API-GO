package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProduct))
	mux.Handle("POST /create-product", http.HandlerFunc(handlers.AddProduct))

	fmt.Println("Server is running on port 8080")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

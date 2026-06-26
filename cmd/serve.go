package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProduct),
		middleware.Arekta,
	))
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(handlers.AddProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))

	fmt.Println("Server is running on port 8080")
	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

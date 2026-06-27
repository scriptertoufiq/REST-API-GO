package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()

	// manager.Use(middleware.Logger, middleware.Hudai, middleware.CorsWithPreflight)

	mux := http.NewServeMux()
	initRoutes(mux, manager)
	wrappedMux := manager.WrapMux(
		mux,
		middleware.Logger,
		middleware.Hudai,
		middleware.CorsWithPreflight,
	)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

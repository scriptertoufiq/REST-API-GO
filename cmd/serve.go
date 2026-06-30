package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	config := config.GetConfig()

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
	fmt.Println("Starting service:", config.ServiceName, "Version:", config.Version, "running on port:", config.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.HttpPort), wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

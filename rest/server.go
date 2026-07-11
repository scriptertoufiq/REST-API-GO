package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)

	fmt.Println("Starting service:", cnf.ServiceName, "Version:", cnf.Version, "running on port:", cnf.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cnf.HttpPort), wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

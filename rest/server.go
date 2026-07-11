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
	config         *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	config *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start() {

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

	fmt.Println("Starting service:", server.config.ServiceName, "Version:", server.config.Version, "running on port:", server.config.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", server.config.HttpPort), wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully")
	}
}

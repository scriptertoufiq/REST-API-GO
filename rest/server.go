package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
)

type Server struct {
	config         *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	config *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		config:         config,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	fmt.Printf(
		"Starting %s v%s on :%d\n",
		s.config.ServiceName,
		s.config.Version,
		s.config.HttpPort,
	)

	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", s.config.HttpPort),
		manager.WrapMux(mux),
	); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

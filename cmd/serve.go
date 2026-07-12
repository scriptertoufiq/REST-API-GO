package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cfg := config.GetConfig()

	// Repositories
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	// Middlewares
	middlewares := middleware.NewMiddlewares(cfg)

	// Handlers
	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewUserHandler(middlewares, userRepo)

	// Server
	server := rest.NewServer(
		cfg,
		productHandler,
		userHandler,
	)

	server.Start()
}

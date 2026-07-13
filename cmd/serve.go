package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cfg := config.GetConfig()

	db, err := db.NewDBConnection()
	if err != nil {
		println("Error connecting to the database:", err)
		return
	}

	// Repositories
	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)

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

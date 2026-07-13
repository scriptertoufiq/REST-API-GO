package cmd

import (
	"fmt"

	"ecommerce/config"
	dbpkg "ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cfg := config.GetConfig()

	dbConn, err := dbpkg.NewDBConnection(cfg)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// Repositories
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

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

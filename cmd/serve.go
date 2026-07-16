package cmd

import (
	"fmt"

	"ecommerce/config"
	"ecommerce/infra/db"
	dbpkg "ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
)

func Serve() {
	cfg := config.GetConfig()

	dbConn, err := dbpkg.NewDBConnection(cfg)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	err = db.MigrateDB(dbConn, "./migrations")
	if err != nil {
		fmt.Println("Error migrating the database:", err)
		return
	}

	// Repositories
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	// domains
	userService := user.NewUserService(userRepo)
	productSvc := product.NewService(productRepo)

	// Middlewares
	middlewares := middleware.NewMiddlewares(cfg)

	// Handlers
	productHandler := productHandler.NewHandler(middlewares, productSvc)
	userHandler := userHandler.NewUserHandler(cfg, userService)

	// Server
	server := rest.NewServer(
		cfg,
		productHandler,
		userHandler,
	)

	server.Start()
}

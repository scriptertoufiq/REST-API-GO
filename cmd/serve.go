package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {

	config := config.GetConfig()
	middlewares := middleware.NewMiddlewares(config)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()
	server := rest.NewServer(config, productHandler, userHandler, reviewHandler)
	server.Start()
}

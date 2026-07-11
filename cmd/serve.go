package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
)

func Serve() {

	config := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()
	server := rest.NewServer(productHandler, userHandler, reviewHandler)
	server.Start(config)
}

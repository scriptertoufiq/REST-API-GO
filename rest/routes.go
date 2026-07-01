package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(handlers.GetProduct),
		middleware.Arekta,
	))
	mux.Handle("POST /create-product", mngr.With(http.HandlerFunc(handlers.AddProduct)))
	mux.Handle("GET /products/{id}", mngr.With(http.HandlerFunc(handlers.GetProductByID)))
}

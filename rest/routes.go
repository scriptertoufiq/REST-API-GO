package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(handlers.GetProduct),
	))
	mux.Handle("POST /create-product", mngr.With(
		http.HandlerFunc(handlers.AddProduct),
		middleware.AuthenticationJWT,
	))
	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.AuthenticationJWT,
	))
	mux.Handle("PUT /products/{id}", mngr.With(
		http.HandlerFunc(handlers.UpdateProduct),
		middleware.AuthenticationJWT,
	))
	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(handlers.DeleteProduct),
		middleware.AuthenticationJWT,
	))

}

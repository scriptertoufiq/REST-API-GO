package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(h.GetProduct),
	))
	mux.Handle("POST /create-product", manager.With(
		http.HandlerFunc(h.AddProduct),
		middleware.AuthenticationJWT,
	))
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(h.GetProductByID),
		middleware.AuthenticationJWT,
	))
	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(h.UpdateProduct),
		middleware.AuthenticationJWT,
	))
	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(h.DeleteProduct),
		middleware.AuthenticationJWT,
	))
}

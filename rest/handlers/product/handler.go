package product

import (
	middleware "ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	service     Service
}

func NewHandler(middlewares *middleware.Middlewares, service Service) *Handler {
	return &Handler{
		middlewares: middlewares,
		service:     service,
	}
}

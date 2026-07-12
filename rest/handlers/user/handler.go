package user

import (
	"ecommerce/repo"
	middleware "ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	userRepo    repo.UserRepo
}

func NewUserHandler(
	middlewares *middleware.Middlewares,
	userRepo repo.UserRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		userRepo:    userRepo,
	}
}

package user

import (
	"ecommerce/config"
)

type Handler struct {
	cnf *config.Config
	svc Service
}

func NewUserHandler(cnf *config.Config, svc Service) *Handler {
	return &Handler{
		cnf: cnf,
		svc: svc,
	}
}

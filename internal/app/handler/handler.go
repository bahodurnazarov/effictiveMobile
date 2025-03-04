package handler

import "github.com/bahodurnazarov/effictiveMobile/internal/app/service"

type Handler struct {
	Service *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{Service: svc}
}

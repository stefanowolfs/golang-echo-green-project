package handler

import (
	"green-echo/internal/app/green/service"
)

type Handler struct {
	userService service.UserService
}

func NewHandler(s service.UserService) *Handler {
	return &Handler{
		userService: s,
	}
}

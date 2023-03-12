package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	user := v1.Group("/users")
	user.GET("", h.GetById)
}

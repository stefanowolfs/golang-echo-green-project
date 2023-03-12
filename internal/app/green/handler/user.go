package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// GetById godoc
// @Summary Get a user
// @Description Get a user of the system
// @ID get-user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponse
// @Failure 400 {object} errs.AppError
// @Failure 401 {object} errs.AppError
// @Failure 422 {object} errs.AppError
// @Failure 404 {object} errs.AppError
// @Failure 500 {object} errs.AppError
// @Router /users/:id [get]
func (h *Handler) GetById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 0, 64)
	if err != nil {
		id = 0
	}

	u, err := h.userService.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response := UserResponse{}
	response.FromUser(u)
	return c.JSON(http.StatusOK, response)
}

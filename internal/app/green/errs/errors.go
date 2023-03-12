package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
	Error   error  `json:",omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func NewConflictError(message string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: message}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

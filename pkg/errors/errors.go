package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Newf(code int, format string, args ...interface{}) *AppError {
	return New(code, fmt.Sprintf(format, args...))
}

func NotFound(message string) *AppError {
	return New(http.StatusNotFound, message)
}

func NotFoundf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusNotFound, format, args...)
}

func BadRequest(message string) *AppError {
	return New(http.StatusBadRequest, message)
}

func BadRequestf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusBadRequest, format, args...)
}

func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, message)
}

func Unauthorizedf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusUnauthorized, format, args...)
}

func Forbidden(message string) *AppError {
	return New(http.StatusForbidden, message)
}

func Forbiddenf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusForbidden, format, args...)
}

func UnprocessableEntity(message string) *AppError {
	return New(http.StatusUnprocessableEntity, message)
}

func UnprocessableEntityf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusUnprocessableEntity, format, args...)
}

func InternalServerError(message string) *AppError {
	return New(http.StatusInternalServerError, message)
}

func InternalServerErrorf(format string, args ...interface{}) *AppError {
	return Newf(http.StatusInternalServerError, format, args...)
}

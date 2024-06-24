package handler

import (
	"net/http"

	"github.com/dxckboi/hugeman-exam/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
	Success bool   `json:"success"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"internal server error"`
	Success bool   `json:"success" example:"false"`
}

func ResponseOK(c *gin.Context, result any) {
	response := Response{
		Success: true,
		Message: "ok",
		Result:  result,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseCreated(c *gin.Context, result any) {
	response := Response{
		Success: true,
		Message: "created",
		Result:  result,
	}

	c.JSON(http.StatusCreated, response)
}

func ResponseError(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	response := ErrorResponse{
		Success: false,
		Message: "internal server error",
	}

	if e, ok := err.(*errors.AppError); ok {
		status = e.Code
		response.Message = e.Message
	}

	c.JSON(status, response)
}

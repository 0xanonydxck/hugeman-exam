package service

import (
	"time"

	"github.com/dxckboi/hugeman-exam/internal/model"
)

type TodoResponse struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      model.Status `json:"status"`
	Image       string       `json:"image"`
	CreatedAt   *time.Time   `json:"created_at"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"base64|datauri,required"`
	Status      string `json:"status" validate:"required,oneof=IN_PROGRESS COMPLETED"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      string `json:"status"`
}

type AllTodoQuery struct {
	Search  string `form:"search"`
	Sort    string `form:"sort" validate:"oneof=created_at title status"`
	Descend bool   `form:"descend"`
}

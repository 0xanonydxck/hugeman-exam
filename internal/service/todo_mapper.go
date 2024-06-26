package service

import (
	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/dxckboi/hugeman-exam/pkg/util"
	"github.com/dxckboi/hugeman-exam/pkg/validator"
)

func TodoModelToResponse(todo *model.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Image:       string(todo.Image),
		CreatedAt:   todo.CreatedAt,
	}
}

func CreateTodoRequestToModel(req *CreateTodoRequest) *model.Todo {
	return &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      model.Status(req.Status),
		Image:       []byte(req.Image),
	}
}

func UpdateTodoRequestToModel(req *UpdateTodoRequest) (*model.Todo, error) {
	t := new(model.Todo)
	if !util.IsEmptyString(req.Title) {
		if err := validator.Var(req.Title, "min=3,max=100"); err != nil {
			return nil, err
		}

		t.Title = req.Title
	}

	if !util.IsEmptyString(req.Description) {
		t.Description = req.Description
	}

	if !util.IsEmptyString(req.Image) {
		if err := validator.Var(req.Image, "base64|datauri"); err != nil {
			return nil, err
		}

		t.Image = []byte(req.Image)
	}

	if !util.IsEmptyString(req.Status) {
		if err := validator.Var(req.Status, "oneof=IN_PROGRESS COMPLETED"); err != nil {
			return nil, err
		}

		t.Status = model.Status(req.Status)
	}

	return t, nil
}

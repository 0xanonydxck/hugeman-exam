package repo

import (
	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/gofrs/uuid"
)

type TodoRepo interface {
	All(...*AllTodoOption) ([]*model.Todo, error)
	Get(id uuid.UUID) (*model.Todo, error)
	Create(todo *model.Todo) error
	Update(id uuid.UUID, todo *model.Todo) (*model.Todo, error)
	Delete(id uuid.UUID) error
}

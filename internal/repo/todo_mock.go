package repo

import (
	"time"

	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"
	"go.openly.dev/pointy"
)

type todoMock struct {
	mock.Mock
}

func NewMockTodoRepo() *todoMock {
	return &todoMock{}
}

func (m *todoMock) All() ([]*model.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*model.Todo), args.Error(1)
}

func (m *todoMock) Get(id uuid.UUID) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *todoMock) Create(todo *model.Todo) error {
	todo.ID = uuid.Must(uuid.NewV4())
	todo.CreatedAt = pointy.Pointer(time.Now())
	args := m.Called()
	return args.Error(0)
}

func (m *todoMock) Update(id uuid.UUID, todo *model.Todo) (*model.Todo, error) {
	todo.ID = id
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *todoMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

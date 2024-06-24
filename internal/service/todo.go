package service

type TodoService interface {
	All(query *AllTodoQuery) ([]*TodoResponse, error)
	Get(id string) (*TodoResponse, error)
	Create(todo *CreateTodoRequest) (*TodoResponse, error)
	Update(id string, todo *UpdateTodoRequest) (*TodoResponse, error)
	SetInProgress(id string) error
	SetCompleted(id string) error
	Delete(id string) error
}

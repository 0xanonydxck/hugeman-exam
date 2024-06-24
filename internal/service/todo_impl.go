package service

import (
	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/dxckboi/hugeman-exam/internal/repo"
	"github.com/dxckboi/hugeman-exam/pkg/errors"
	"github.com/dxckboi/hugeman-exam/pkg/validator"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type todoService struct {
	repo repo.TodoRepo
}

func NewTodoService(repo repo.TodoRepo) *todoService {
	return &todoService{
		repo: repo,
	}
}

func (s *todoService) All(query *AllTodoQuery) ([]*TodoResponse, error) {
	todos, err := s.repo.All()
	if err != nil {
		log.Error().Err(err).Msg("service::All() - failed to get all todo tasks")
		return nil, err
	}

	result := make([]*TodoResponse, 0)
	for _, todo := range todos {
		result = append(result, &TodoResponse{
			ID:          todo.ID.String(),
			Title:       todo.Title,
			Description: todo.Description,
			Status:      todo.Status,
			Image:       string(todo.Image),
			CreatedAt:   todo.CreatedAt,
		})
	}

	return result, nil
}

func (s *todoService) Get(id string) (*TodoResponse, error) {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::Get() - failed to parse id to uuid")
		return nil, errors.ErrParsedUUID
	}

	todo, err := s.repo.Get(parsedID)
	if err != nil {
		log.Error().Err(err).Msg("service::Get() - failed to get todo task")
		return nil, err
	}

	return &TodoResponse{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Image:       string(todo.Image),
		CreatedAt:   todo.CreatedAt,
	}, nil
}

func (s *todoService) Create(todo *CreateTodoRequest) (*TodoResponse, error) {
	if err := validator.Struct(todo); err != nil {
		log.Error().Err(err).Msg("service::Create() - failed to validate todo request")
		return nil, err
	}

	t := &model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Image:       []byte(todo.Image),
		Status:      model.Status(todo.Status),
	}

	err := s.repo.Create(t)
	if err != nil {
		log.Error().Err(err).Msg("service::Create() - failed to create todo task")
		return nil, err
	}

	return &TodoResponse{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Image:       string(t.Image),
		CreatedAt:   t.CreatedAt,
	}, nil
}

func (s *todoService) Update(id string, todo *UpdateTodoRequest) (*TodoResponse, error) {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::Update() - failed to parse id to uuid")
		return nil, errors.ErrParsedUUID
	}

	t := new(model.Todo)
	t.Title = todo.Title
	t.Description = todo.Description
	t.Status = model.Status(todo.Status)
	image := []byte(todo.Image)
	if len(image) > 0 {
		t.Image = image
	}

	t, err = s.repo.Update(parsedID, t)
	if err != nil {
		log.Error().Err(err).Msg("service::Update() - failed to update todo task")
		return nil, err
	}

	return &TodoResponse{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Image:       string(t.Image),
		CreatedAt:   t.CreatedAt,
	}, nil

}

func (s *todoService) SetInProgress(id string) error {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::SetInProgress() - failed to parse id to uuid")
		return errors.ErrParsedUUID
	}

	t := &model.Todo{
		Status: model.IN_PROGRESS,
	}

	_, err = s.repo.Update(parsedID, t)
	if err != nil {
		log.Error().Err(err).Msg("service::SetInProgress() - failed to update todo task")
		return err
	}

	return nil
}

func (s *todoService) SetCompleted(id string) error {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::SetCompleted() - failed to parse id to uuid")
		return errors.ErrParsedUUID
	}

	t := &model.Todo{
		Status: model.COMPLETED,
	}

	_, err = s.repo.Update(parsedID, t)
	if err != nil {
		log.Error().Err(err).Msg("service::SetCompleted() - failed to update todo task")
		return err
	}

	return nil
}

func (s *todoService) Delete(id string) error {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::Delete() - failed to parse id to uuid")
		return errors.ErrParsedUUID
	}

	err = s.repo.Delete(parsedID)
	if err != nil {
		log.Error().Err(err).Msg("service::Delete() - failed to delete todo task")
		return err
	}

	return nil
}

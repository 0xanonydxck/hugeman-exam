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
	opt, err := AllTodoQueryToOption(query)
	if err != nil {
		log.Error().Err(err).Msg("service::All() - failed to map query to option")
		return nil, err
	}

	todos, err := s.repo.All(opt)
	if err != nil {
		log.Error().Err(err).Msg("service::All() - failed to get all todo tasks")
		return nil, err
	}

	result := make([]*TodoResponse, 0)
	for _, todo := range todos {
		result = append(result, TodoModelToResponse(todo))
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

	return TodoModelToResponse(todo), nil
}

func (s *todoService) Create(todo *CreateTodoRequest) (*TodoResponse, error) {
	if err := validator.Struct(todo); err != nil {
		log.Error().Err(err).Msg("service::Create() - failed to validate todo request")
		return nil, err
	}

	t := CreateTodoRequestToModel(todo)
	if err := s.repo.Create(t); err != nil {
		log.Error().Err(err).Msg("service::Create() - failed to create todo task")
		return nil, err
	}

	return TodoModelToResponse(t), nil
}

func (s *todoService) Update(id string, req *UpdateTodoRequest) (*TodoResponse, error) {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::Update() - failed to parse id to uuid")
		return nil, errors.ErrParsedUUID
	}

	todo, err := UpdateTodoRequestToModel(req)
	if err != nil {
		log.Error().Err(err).Msg("service::Update() - failed to map todo request to model")
		return nil, err
	}

	todo, err = s.repo.Update(parsedID, todo)
	if err != nil {
		log.Error().Err(err).Msg("service::Update() - failed to update todo task")
		return nil, err
	}

	return TodoModelToResponse(todo), nil

}

func (s *todoService) SetInProgress(id string) error {
	parsedID, err := uuid.FromString(id)
	if err != nil {
		log.Error().Err(err).Msg("service::SetInProgress() - failed to parse id to uuid")
		return errors.ErrParsedUUID
	}

	todo := &model.Todo{
		Status: model.IN_PROGRESS,
	}

	_, err = s.repo.Update(parsedID, todo)
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

	todo := &model.Todo{
		Status: model.COMPLETED,
	}

	_, err = s.repo.Update(parsedID, todo)
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

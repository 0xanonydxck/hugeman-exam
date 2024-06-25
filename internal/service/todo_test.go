package service_test

import (
	"testing"
	"time"

	"github.com/dxckboi/hugeman-exam/internal/model"
	"github.com/dxckboi/hugeman-exam/internal/repo"
	"github.com/dxckboi/hugeman-exam/internal/service"
	"github.com/dxckboi/hugeman-exam/pkg/constant"
	"github.com/dxckboi/hugeman-exam/pkg/validator"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTodoAll(t *testing.T) {
	type TestCase struct {
		Name         string
		RepoResult   []*model.Todo
		RepoError    error
		ExpectResult []*service.TodoResponse
		ExpectError  error
	}

	now := time.Now()
	testCases := []TestCase{
		{
			Name: "Success",
			RepoResult: []*model.Todo{
				{
					ID:          uuid.FromStringOrNil("f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1"),
					Title:       "Task 1",
					Description: "Description 1",
					Status:      model.IN_PROGRESS,
					Image:       []byte(constant.HM_LOGO),
					CreatedAt:   &now,
				},
				{
					ID:          uuid.FromStringOrNil("e4b3b8f1-6fbb-4b9d-8a4d-0c2e77e1f7a2"),
					Title:       "Task 2",
					Description: "Description 2",
					Status:      model.COMPLETED,
					Image:       []byte(constant.HM_LOGO),
					CreatedAt:   &now,
				},
			},
			RepoError: nil,
			ExpectResult: []*service.TodoResponse{
				{
					ID:          "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
					Title:       "Task 1",
					Description: "Description 1",
					Status:      model.IN_PROGRESS,
					Image:       constant.HM_LOGO,
					CreatedAt:   &now,
				},
				{
					ID:          "e4b3b8f1-6fbb-4b9d-8a4d-0c2e77e1f7a2",
					Title:       "Task 2",
					Description: "Description 2",
					Status:      model.COMPLETED,
					Image:       constant.HM_LOGO,
					CreatedAt:   &now,
				},
			},
			ExpectError: nil,
		},
		{
			Name:         "Success Empty",
			RepoResult:   make([]*model.Todo, 0),
			RepoError:    nil,
			ExpectResult: make([]*service.TodoResponse, 0),
			ExpectError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("All").Return(tc.RepoResult, tc.RepoError)

			svc := service.NewTodoService(repo)
			result, err := svc.All(nil)

			assert.Equal(t, tc.ExpectResult, result)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}

func TestTodoGet(t *testing.T) {
	type TestCase struct {
		Name         string
		TodoID       string
		RepoResult   *model.Todo
		RepoError    error
		ExpectResult *service.TodoResponse
		ExpectError  error
	}

	now := time.Now()
	testCases := []TestCase{
		{
			Name:   "Success",
			TodoID: "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
			RepoResult: &model.Todo{
				ID:          uuid.FromStringOrNil("f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1"),
				Title:       "Task 1",
				Description: "Description 1",
				Status:      model.IN_PROGRESS,
				Image:       []byte(constant.HM_LOGO),
				CreatedAt:   &now,
			},
			RepoError: nil,
			ExpectResult: &service.TodoResponse{
				ID:          "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
				Title:       "Task 1",
				Description: "Description 1",
				Status:      model.IN_PROGRESS,
				Image:       constant.HM_LOGO,
				CreatedAt:   &now,
			},
			ExpectError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Get", uuid.FromStringOrNil(tc.TodoID)).Return(tc.RepoResult, tc.RepoError)

			svc := service.NewTodoService(repo)
			result, err := svc.Get(tc.TodoID)

			assert.Equal(t, tc.ExpectResult, result)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}

func TestTodoCreate(t *testing.T) {
	type TestCase struct {
		Name        string
		TodoRequest *service.CreateTodoRequest
		RepoError   error
		ExpectError error
	}

	testCases := []TestCase{
		{
			Name: "Success",
			TodoRequest: &service.CreateTodoRequest{
				Title:       "Task 1",
				Description: "Description 1",
				Image:       constant.HM_LOGO,
				Status:      string(model.IN_PROGRESS),
			},
			RepoError:   nil,
			ExpectError: nil,
		},
	}

	validator.Init()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Create").Return(tc.RepoError)

			svc := service.NewTodoService(repo)
			result, err := svc.Create(tc.TodoRequest)
			assert.Equal(t, tc.ExpectError, err)
			assert.NotEmpty(t, result.ID)
			assert.Equal(t, tc.TodoRequest.Title, result.Title)
			assert.Equal(t, tc.TodoRequest.Description, result.Description)
			assert.Equal(t, model.Status(tc.TodoRequest.Status), result.Status)
			assert.Equal(t, tc.TodoRequest.Image, result.Image)
			assert.NotNil(t, result.CreatedAt)
		})
	}
}

func TestTodoUpdate(t *testing.T) {
	type TestCase struct {
		Name        string
		TodoID      string
		TodoRequest *service.UpdateTodoRequest
		RepoError   error
		RepoResult  *model.Todo
		ExpectError error
	}

	testCases := []TestCase{
		{
			Name:   "Success",
			TodoID: "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
			TodoRequest: &service.UpdateTodoRequest{
				Title:       "Task 1",
				Description: "Description 1",
				Image:       constant.HM_LOGO,
				Status:      string(model.IN_PROGRESS),
			},
			RepoResult: &model.Todo{
				ID:          uuid.FromStringOrNil("f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1"),
				Title:       "Task 1",
				Description: "Description 1",
				Status:      model.IN_PROGRESS,
				Image:       []byte(constant.HM_LOGO),
				CreatedAt:   nil,
			},
			RepoError:   nil,
			ExpectError: nil,
		},
	}

	validator.Init()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Update", mock.Anything).Return(tc.RepoResult, tc.RepoError)

			svc := service.NewTodoService(repo)
			result, err := svc.Update(tc.TodoID, tc.TodoRequest)
			assert.Equal(t, tc.ExpectError, err)
			assert.NotEmpty(t, result.ID)
			assert.Equal(t, tc.TodoRequest.Title, result.Title)
			assert.Equal(t, tc.TodoRequest.Description, result.Description)
			assert.Equal(t, model.Status(tc.TodoRequest.Status), result.Status)
			assert.Equal(t, tc.TodoRequest.Image, result.Image)
		})
	}
}

func TestSetInProgress(t *testing.T) {
	type TestCase struct {
		Name        string
		TodoID      string
		RepoResult  *model.Todo
		RepoError   error
		ExpectError error
	}

	testCases := []TestCase{
		{
			Name:        "Success",
			TodoID:      "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
			RepoResult:  &model.Todo{},
			RepoError:   nil,
			ExpectError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Update", mock.Anything).Return(tc.RepoResult, tc.RepoError)

			svc := service.NewTodoService(repo)
			err := svc.SetInProgress(tc.TodoID)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}

func TestSetCompleted(t *testing.T) {
	type TestCase struct {
		Name        string
		TodoID      string
		RepoResult  *model.Todo
		RepoError   error
		ExpectError error
	}

	testCases := []TestCase{
		{
			Name:        "Success",
			TodoID:      "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
			RepoResult:  &model.Todo{},
			RepoError:   nil,
			ExpectError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Update", mock.Anything).Return(tc.RepoResult, tc.RepoError)

			svc := service.NewTodoService(repo)
			err := svc.SetCompleted(tc.TodoID)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}

func TestTodoDelete(t *testing.T) {
	type TestCase struct {
		Name        string
		TodoID      string
		RepoError   error
		ExpectError error
	}

	testCases := []TestCase{
		{
			Name:        "Success",
			TodoID:      "f4b3b8f1-7fbb-4b9d-8a4d-0c2e77e1f7a1",
			RepoError:   nil,
			ExpectError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repo := repo.NewMockTodoRepo()
			repo.On("Delete", mock.Anything).Return(tc.RepoError)

			svc := service.NewTodoService(repo)
			err := svc.Delete(tc.TodoID)
			assert.Equal(t, tc.ExpectError, err)
		})
	}
}

package usecase_test

import (
	"TaskManagement/db"
	mocks "TaskManagement/db/mock"
	taskmodel "TaskManagement/internal/model/task"
	"TaskManagement/internal/usecase"
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_GetTasks(t *testing.T) {
	mockQuerier := new(mocks.MockQuerier)
	uc := usecase.New(mockQuerier)

	var niltasks []db.Task // это nil-срез нужного типа
	//mockQuerier.On("ListTasks", mock.Anything).Return(tasks, nil)

	mockQuerier.On("ListTasks", mock.Anything).Return(niltasks, errors.New("db error"))
	tasks := uc.GetTasks()
	assert.Nil(t, tasks)
	//assert.Error(t, err)
}

func TestUseCase_GetTask(t *testing.T) {
	mockQuerier := new(mocks.MockQuerier)
	uc := usecase.New(mockQuerier)

	dbTask := db.Task{
		Taskid:         123,
		TaskNumber:     "T-123",
		Description:    sql.NullString{String: "Test task", Valid: true},
		Priority:       sql.NullInt32{Int32: 1, Valid: true},
		CreateDatetime: time.Now(),
	}

	mockQuerier.On("GetTaskByID", mock.Anything, int32(123)).Return(dbTask, nil)

	task, err := uc.GetTask(123)

	assert.NoError(t, err)
	assert.Equal(t, 123, task.Taskid)
	assert.Equal(t, "T-123", task.TaskNumber)
}

func TestUseCase_CreateTask(t *testing.T) {
	mockQuerier := new(mocks.MockQuerier)
	uc := usecase.New(mockQuerier)

	req := taskmodel.CreateTaskRequest{
		TaskNumber:  "T-999",
		Description: "New task",
		Priority:    3,
	}

	mockQuerier.On("CreateTask", mock.Anything, mock.AnythingOfType("db.CreateTaskParams")).Return(nil)

	err := uc.CreateTask(context.Background(), req)

	assert.NoError(t, err)
	mockQuerier.AssertExpectations(t)
}

func TestUseCase_DeleteTask(t *testing.T) {
	mockQuerier := new(mocks.MockQuerier)
	uc := usecase.New(mockQuerier)

	mockQuerier.On("DeleteTask", mock.Anything, int32(42)).Return(nil)

	err := uc.DeleteTask(context.Background(), 42)

	assert.NoError(t, err)
}

func TestUseCase_UpdateTask(t *testing.T) {
	mockQuerier := new(mocks.MockQuerier)
	uc := usecase.New(mockQuerier)

	desc := "Updated desc"
	prio := 5
	taskNum := "T-999"

	req := taskmodel.UpdateTaskRequest{
		Description: &desc,
		Priority:    &prio,
		TaskNumber:  &taskNum,
	}

	dbTask := db.Task{
		Taskid:      100,
		TaskNumber:  "T-100",
		Description: sql.NullString{String: "Updated desc", Valid: true},
		Priority:    sql.NullInt32{Int32: 5, Valid: true},
	}

	mockQuerier.On("UpdateTask", mock.Anything, mock.AnythingOfType("db.UpdateTaskParams")).Return(dbTask, nil)

	task, err := uc.UpdateTask(context.Background(), req, 100)

	assert.NoError(t, err)
	assert.Equal(t, "Updated desc", task.Description)
	assert.Equal(t, 5, task.Priority)
}

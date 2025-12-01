// internal/mocks/querier_mock.go
package mocks

import (
	"TaskManagement/db"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockQuerier struct {
	mock.Mock
}

func (m *MockQuerier) ListTasks(ctx context.Context) ([]db.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]db.Task), args.Error(1)
}

func (m *MockQuerier) GetTaskByID(ctx context.Context, id int32) (db.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.Task), args.Error(1)
}

func (m *MockQuerier) CreateTask(ctx context.Context, arg db.CreateTaskParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockQuerier) UpdateTask(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.Task), args.Error(1)
}

func (m *MockQuerier) DeleteTask(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockQuerier) ListUsers(ctx context.Context) ([]db.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]db.User), args.Error(1)
}

func (m *MockQuerier) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockQuerier) CreateUser(ctx context.Context, arg db.CreateUserParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockQuerier) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.User), args.Error(1)
}

func (m *MockQuerier) DeleteUser(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

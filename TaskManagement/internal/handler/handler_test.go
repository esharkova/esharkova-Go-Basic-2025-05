package handler_test

import (
	"TaskManagement/internal/handler"
	task "TaskManagement/internal/model/task"
	user "TaskManagement/internal/model/user"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "TaskManagement/internal/handler/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateUser(t *testing.T) {
	validUser := user.CreateUserRequest{
		FirstName: "John Doe",
		LastName:  "Doe",
	}

	tests := []struct {
		name           string // description of this test case
		body           string
		mockSetup      func(*mock.MockuseCase)
		expectedStatus int
	}{
		{
			name: "Success",
			body: func() string {
				b, _ := json.Marshal(validUser)
				return string(b)
			}(),
			mockSetup: func(m *mock.MockuseCase) {
				m.EXPECT().CreateUser(gomock.Any(), validUser).Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUC := mock.NewMockuseCase(ctrl)
			tt.mockSetup(mockUC)

			h := handler.New(mockUC)
			router := gin.New()
			router.POST("/users", h.CreateUser())

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

		})
	}
}

func TestHandler_CreateTask(t *testing.T) {

	validTask := task.CreateTaskRequest{
		Description: "Test description",
		TaskNumber:  "TaskNumber",
		Priority:    1,
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		body           string
		mockSetup      func(*mock.MockuseCase)
		expectedStatus int
	}{
		{
			name: "Success",
			body: func() string {
				b, _ := json.Marshal(validTask)
				return string(b)
			}(),
			mockSetup: func(m *mock.MockuseCase) {
				m.EXPECT().CreateTask(gomock.Any(), validTask).Return(nil)
			},
			expectedStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUC := mock.NewMockuseCase(ctrl)
			mockUC.EXPECT().CreateTask(gomock.Any(), validTask).Return(nil)

			handler := handler.New(mockUC)
			router := gin.New()
			router.POST("/tasks", handler.CreateTask())

			body, _ := json.Marshal(validTask)
			req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusCreated, w.Code)

			var resp task.Task
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

		})
	}
}

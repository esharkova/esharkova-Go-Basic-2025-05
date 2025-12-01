package repository_test

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"TaskManagement/internal/repository"
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// resetState очищает глобальное состояние и удаляет файлы перед каждым тестом
func resetState() {
	repository.MuUser.Lock()
	repository.Users = nil
	repository.MuUser.Unlock()

	repository.MuTask.Lock()
	repository.Tasks = nil
	repository.MuTask.Unlock()

	repository.LastUsersCount = 0
	repository.LastTasksCount = 0

	// Удаляем тестовые файлы
	os.Remove("users.json")
	os.Remove("tasks.json")
}

// чтение строк в файле
func readJSONLines(filename string) ([]json.RawMessage, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []json.RawMessage
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, json.RawMessage(line))
	}
	return lines, scanner.Err()
}

func GenerateUser(i int) taskUser.User {

	newUser := taskUser.User{
		Userid:    i,
		FirstName: "FirstName" + strconv.Itoa(i),
		LastName:  "LastName" + strconv.Itoa(i),
	}

	newUser.AddPassport("Passport" + strconv.Itoa(i))

	return newUser

}

func GenerateTask(i int) task.Task {

	newTask := task.Task{
		Taskid:             i,
		TaskNumber:         "TaskNumber" + strconv.Itoa(i),
		Description:        "TaskDescription" + strconv.Itoa(i),
		CreateDateTime:     time.Now(),
		CompletionDateTime: time.Now().AddDate(0, 0, 7),
	}

	return newTask

}

// Вспомогательная функция для указателей на строку
func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func TestAddUser(t *testing.T) {
	resetState()

	user := taskUser.User{
		Userid:    1,
		FirstName: "User1FirstName",
		LastName:  "User1LastName",
	}
	repository.AddUser(user)

	assert.NotEqual(t, len(repository.Users), 0)
	assert.Equal(t, len(repository.Users), 1)

	//сравниваем пользователя в слайсе
	if len(repository.Users) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(repository.Users))
	}
	if repository.Users[0].Userid != 1 || repository.Users[0].FirstName != "User1FirstName" || repository.Users[0].LastName != "User1LastName" {
		t.Errorf("User data mismatch: %+v", repository.Users[0])
	}

	lines, err := readJSONLines("users.json")
	assert.NoError(t, err)
	assert.Equal(t, len(lines), 1)

	//сравниваем пользователя в файле
	var savedUser taskUser.User
	if err := json.Unmarshal(lines[0], &savedUser); err != nil {
		t.Fatalf("Failed to unmarshal user from file: %v", err)
	}
	if savedUser.Userid != 1 || savedUser.FirstName != "User1FirstName" || savedUser.LastName != "User1LastName" {
		t.Errorf("Saved user mismatch: %+v", savedUser)
	}

}

func TestAddTask(t *testing.T) {
	resetState()

	currentDateTime := time.Now()

	tk := task.Task{
		Taskid:         1,
		TaskNumber:     "TaskNumber1",
		Description:    "TaskNumber1Description",
		CreateDateTime: currentDateTime,
		Priority:       1,
	}

	repository.AddTask(tk)

	assert.NotEqual(t, len(repository.Tasks), 0)
	require.Len(t, repository.Tasks, 1, "Expected 1 task in memory")

	assert.Equal(t, repository.Tasks[0].Taskid, tk.Taskid)
	assert.Equal(t, repository.Tasks[0].TaskNumber, tk.TaskNumber)
	assert.Equal(t, repository.Tasks[0].Description, tk.Description)
	assert.Equal(t, repository.Tasks[0].CreateDateTime, tk.CreateDateTime)
	assert.Equal(t, repository.Tasks[0].Priority, tk.Priority)

	lines, err := readJSONLines("tasks.json")
	require.NoError(t, err)
	require.Len(t, lines, 1)

	var savedTask task.Task
	require.NoError(t, json.Unmarshal(lines[0], &savedTask), "Failed to unmarshal task from JSON")

	assert.Equal(t, savedTask.Taskid, tk.Taskid)
	assert.Equal(t, savedTask.TaskNumber, tk.TaskNumber)
	assert.Equal(t, savedTask.Description, tk.Description)
	assert.Equal(t, savedTask.Priority, tk.Priority)

}

func TestGetUserById(t *testing.T) {
	resetState()

	newUser := GenerateUser(1)

	repository.AddUser(newUser)

	user, err := repository.GetUserById(1)
	require.NoError(t, err)

	assert.Equal(t, user.Userid, newUser.Userid)
	assert.Equal(t, user.FirstName, newUser.FirstName)
	assert.Equal(t, user.LastName, newUser.LastName)
	assert.Equal(t, user.GetPassport(), newUser.GetPassport())

	_, err = repository.GetUserById(999)
	require.Error(t, err, "Expected error for non-existent user")
	assert.Equal(t, "Пользователь с идентификатором 999 не найден", err.Error())

}

func TestGetTaskById(t *testing.T) {
	resetState()

	newTask := GenerateTask(1)

	repository.AddTask(newTask)

	task, err := repository.GetTaskById(1)
	require.NoError(t, err)
	assert.Equal(t, task.Taskid, newTask.Taskid)
	assert.Equal(t, task.TaskNumber, newTask.TaskNumber)
	assert.Equal(t, task.Description, newTask.Description)
	assert.Equal(t, task.CreateDateTime, newTask.CreateDateTime)
	assert.Equal(t, task.CompletionDateTime, newTask.CompletionDateTime)
	assert.Equal(t, task.Priority, newTask.Priority)

	_, err = repository.GetTaskById(999)
	require.Error(t, err, "Expected error for non-existent user")
	assert.Equal(t, "Задача с идентификатором 999 не найдена", err.Error())
}

func TestUpdateUser(t *testing.T) {
	resetState()

	user := GenerateUser(1)

	repository.AddUser(user)

	newUser := taskUser.UpdateUserRequest{
		FirstName: strPtr("NewFirstName"),
		LastName:  strPtr("NewLastName"),
	}

	updated, err := repository.UpdateUser(1, newUser)

	require.NoError(t, err)
	assert.Equal(t, updated.Userid, user.Userid)
	assert.Equal(t, updated.FirstName, "NewFirstName")
	assert.Equal(t, updated.LastName, "NewLastName")

	// Проверяем файл
	lines, err := readJSONLines("users.json")
	require.NoError(t, err)
	require.Len(t, lines, 1)

	var saved taskUser.User
	require.NoError(t, json.Unmarshal(lines[0], &saved), "Failed to unmarshal task from JSON")

	assert.Equal(t, saved.Userid, user.Userid)
	assert.Equal(t, saved.FirstName, "NewFirstName")
	assert.Equal(t, saved.LastName, "NewLastName")

}

func TestUpdateTask(t *testing.T) {
	resetState()

	currTask := GenerateTask(1)

	repository.AddTask(currTask)

	newTask := task.UpdateTaskRequest{
		TaskNumber:  strPtr("NewTaskNumber"),
		Description: strPtr("NewDescription"),
		Priority:    intPtr(2),
	}

	updated, err := repository.UpdateTask(1, newTask)
	require.NoError(t, err)
	assert.Equal(t, updated.Taskid, currTask.Taskid)
	assert.Equal(t, updated.TaskNumber, "NewTaskNumber")
	assert.Equal(t, updated.Description, "NewDescription")
	assert.Equal(t, updated.Priority, 2)

	lines, err := readJSONLines("tasks.json")
	require.NoError(t, err)
	require.Len(t, lines, 1)

	var saved task.Task
	require.NoError(t, json.Unmarshal(lines[0], &saved), "Failed to unmarshal task from JSON")
	assert.Equal(t, saved.Taskid, currTask.Taskid)
	assert.Equal(t, saved.TaskNumber, "NewTaskNumber")
	assert.Equal(t, saved.Description, "NewDescription")
	assert.Equal(t, saved.Priority, 2)
}

func TestDeleteUser(t *testing.T) {
	resetState()

	repository.AddUser(GenerateUser(1))
	repository.AddUser(GenerateUser(2))

	err := repository.DeleteUser(1)
	require.NoError(t, err)

	assert.NotEqual(t, len(repository.Users), 0)
	require.Len(t, repository.Users, 1, "Expected 1 user in memory")

	assert.Equal(t, repository.Users[0].Userid, 2)
	assert.Equal(t, repository.Users[0].FirstName, "FirstName2")
	assert.Equal(t, repository.Users[0].LastName, "LastName2")

	lines, err := readJSONLines("users.json")
	require.NoError(t, err)
	require.Len(t, lines, 1)

	var saved taskUser.User
	require.NoError(t, json.Unmarshal(lines[0], &saved), "Failed to unmarshal task from JSON")

	assert.Equal(t, saved.Userid, 2)
	assert.Equal(t, saved.FirstName, "FirstName2")
	assert.Equal(t, saved.LastName, "LastName2")

	// Удаление несуществующего
	err = repository.DeleteUser(999)
	require.Error(t, err, "Expected error for non-existent user")
	assert.Equal(t, "Пользователь с идентификатором 999 не найден", err.Error())
}

func TestDeleteTask(t *testing.T) {
	resetState()

	task1 := GenerateTask(1)
	task2 := GenerateTask(2)

	repository.AddTask(task1)
	repository.AddTask(task2)

	err := repository.DeleteTask(1)
	require.NoError(t, err)

	assert.NotEqual(t, len(repository.Tasks), 0)
	require.Len(t, repository.Tasks, 1, "Expected 1 task in memory")

	assert.Equal(t, repository.Tasks[0].Taskid, task2.Taskid)
	assert.Equal(t, repository.Tasks[0].TaskNumber, task2.TaskNumber)
	assert.Equal(t, repository.Tasks[0].Description, task2.Description)
	assert.Equal(t, repository.Tasks[0].CreateDateTime, task2.CreateDateTime)
	assert.Equal(t, repository.Tasks[0].Priority, task2.Priority)

	lines, err := readJSONLines("tasks.json")
	require.NoError(t, err)
	require.Len(t, lines, 1)

	var savedTask task.Task
	require.NoError(t, json.Unmarshal(lines[0], &savedTask), "Failed to unmarshal task from JSON")

	assert.Equal(t, savedTask.Taskid, task2.Taskid)
	assert.Equal(t, savedTask.TaskNumber, task2.TaskNumber)
	assert.Equal(t, savedTask.Description, task2.Description)
	assert.Equal(t, savedTask.Priority, task2.Priority)

	// Удаление несуществующего
	err = repository.DeleteTask(999)
	require.Error(t, err, "Expected error for non-existent task")
	assert.Equal(t, "Задача с идентификатором 999 не найдена", err.Error())

}

func TestProcessValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tm      repository.TaskManage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := repository.ProcessValue(tt.tm)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ProcessValue() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ProcessValue() succeeded unexpectedly")
			}
		})
	}
}

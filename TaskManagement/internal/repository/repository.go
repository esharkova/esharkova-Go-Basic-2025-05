package repository

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

var LastUsersCount, LastTasksCount int

var Users []*taskUser.User
var Tasks []*task.Task

var MuUser sync.Mutex
var MuTask sync.Mutex
var MuUserCount sync.Mutex
var MuTaskCount sync.Mutex

type TaskManage interface {
	Insert() int
}

func ProcessValue(tm TaskManage) error {

	switch v := tm.(type) {

	case taskUser.User:
		MuUser.Lock()
		fmt.Println("Processing User", v)
		Users = append(Users, &v)
		MuUser.Unlock()
	case task.Task:
		MuTask.Lock()
		fmt.Println("Processing Task", v)
		Tasks = append(Tasks, &v)
		MuTask.Unlock()
	default:
		fmt.Println("Unknown type", v)
		errors.New("Unknown type")

	}

	return nil

}

func GetCopyUsers(currentUsers []*taskUser.User) []*taskUser.User {
	MuUser.Lock()
	defer MuUser.Unlock()
	copiedUsers := make([]*taskUser.User, len(currentUsers))
	copy(copiedUsers, currentUsers)
	return copiedUsers

}

func GetCopyTasks(currentTasks []*task.Task) []*task.Task {
	MuTask.Lock()
	defer MuTask.Unlock()
	copiedTasks := make([]*task.Task, len(currentTasks))
	copy(copiedTasks, currentTasks)
	return copiedTasks

}

func AddUser(user taskUser.User) {
	MuUser.Lock()
	defer MuUser.Unlock()
	Users = append(Users, &user)
	fmt.Println("Добавлен пользователь в слайс")
	WriteUserToFile(user)
}

func AddTask(task task.Task) {
	MuTask.Lock()
	defer MuTask.Unlock()
	Tasks = append(Tasks, &task)
	fmt.Println("Добавлена задача в слайс")
	WriteTaskToFile(task)

}

func UpdateTask(id int, newTask task.UpdateTaskRequest) (task.Task, error) {
	MuTask.Lock()
	defer MuTask.Unlock()

	for _, t := range Tasks {
		if t.Taskid == id {
			if newTask.TaskNumber != nil {

				t.TaskNumber = *newTask.TaskNumber
			}

			if newTask.Description != nil {
				t.Description = *newTask.Description

			}

			if newTask.Priority != nil {
				t.Priority = *newTask.Priority

			}

			UpdateTaskInFile(id, newTask)

			return *t, nil
		}
	}

	return task.Task{}, fmt.Errorf("Задача с идентификатором %d не найдена", id)
}

func GetTaskById(id int) (task.Task, error) {
	for _, t := range Tasks {
		if t.Taskid == id {
			return *t, nil
		}
	}

	return task.Task{}, fmt.Errorf("Задача с идентификатором %d не найдена", id)
}

func DeleteTask(id int) error {

	MuTask.Lock()
	defer MuTask.Unlock()

	for i, t := range Tasks {
		if t.Taskid == id {

			Tasks = append(Tasks[:i], Tasks[i+1:]...)

			DeleteTaskInFile(id)

			return nil
		}
	}

	return fmt.Errorf("Задача с идентификатором %d не найдена", id)
}

func GetUserById(id int) (taskUser.User, error) {
	for _, u := range Users {
		if u.Userid == id {
			return *u, nil
		}
	}

	return taskUser.User{}, fmt.Errorf("Пользователь с идентификатором %d не найден", id)
}

func UpdateUser(id int, newUser taskUser.UpdateUserRequest) (taskUser.User, error) {

	MuUser.Lock()
	defer MuUser.Unlock()

	for _, u := range Users {
		if u.Userid == id {
			if newUser.FirstName != nil {

				u.FirstName = *newUser.FirstName
			}

			if newUser.LastName != nil {
				u.LastName = *newUser.LastName

			}

			UpdateUserInFile(id, newUser)

			return *u, nil
		}
	}

	return taskUser.User{}, fmt.Errorf("Пользователь с идентификатором %d не найден", id)
}

func DeleteUser(id int) error {

	MuUser.Lock()
	defer MuUser.Unlock()

	for i, u := range Users {
		if u.Userid == id {

			Users = append(Users[:i], Users[i+1:]...)

			DeleteUserInFile(id)

			return nil
		}
	}

	return fmt.Errorf("Пользователь с идентификатором %d не найден", id)
}

func GetLastUserCount() int {
	MuUserCount.Lock()
	defer MuUserCount.Unlock()
	return LastUsersCount
}

func GetLastTasksCount() int {
	MuTaskCount.Lock()
	defer MuTaskCount.Unlock()
	return LastTasksCount
}

func SetLastUserCount(count int) {
	MuUserCount.Lock()
	defer MuUserCount.Unlock()
	LastUsersCount = count
}

func SetLastTasksCount(count int) {
	MuTaskCount.Lock()
	defer MuTaskCount.Unlock()
	LastTasksCount = count
}

func WriteUserToFile(user taskUser.User) {

	file, err := os.OpenFile("users.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer file.Close()

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
	} else {
		_, err = file.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
		fmt.Println("Добавлен пользователь в файл")
	}

}

func WriteTaskToFile(task task.Task) {

	file, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer file.Close()

	jsonData, err := json.Marshal(task)

	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
	} else {
		_, err = file.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}
		fmt.Println("Добавлена задача в файл")
	}

}

func ReadUsersFromFileAndAddToSlice() {

	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл users.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		var user taskUser.User
		err := json.Unmarshal(line, &user)
		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", err)
			continue
		}
		Users = append(Users, &user)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла users:", err)
		return
	}

	fmt.Printf("Прочитано %d записей в файле users.json:\n", len(Users))
	for _, user := range Users {
		fmt.Printf("  %+v\n", *user)
	}

	SetLastUserCount(len(Users))

}

func ReadTasksFromFileAndAddToSlice() {

	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл tasks.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		var task task.Task
		err := json.Unmarshal(line, &task)
		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", err)
			continue
		}
		Tasks = append(Tasks, &task)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла tasks:", err)
		return
	}

	fmt.Printf("Прочитано %d записей в файле tasks.json:\n", len(Tasks))
	for _, task := range Tasks {
		fmt.Printf("  %+v\n", *task)
	}

	SetLastTasksCount(len(Tasks))

}

func UpdateTaskInFile(id int, newTask task.UpdateTaskRequest) {

	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл tasks.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаём временный файл для записи обновлённых данных
	tempFile, err := os.CreateTemp("", "jsonl_*.tmp")
	if err != nil {
		fmt.Println("не удалось создать временный файл: %w", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Bytes()
		var task task.Task

		err := json.Unmarshal(line, &task)

		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", line, err)
			tempFile.WriteString(string(line) + "\n")
			continue
		}

		if task.Taskid == id {
			if newTask.TaskNumber != nil {
				task.TaskNumber = *newTask.TaskNumber
			}
			if newTask.Description != nil {
				task.Description = *newTask.Description
			}
			if newTask.Priority != nil {
				task.Priority = *newTask.Priority
			}
			found = true
		}

		updatedLine, err := json.Marshal(task)
		if err != nil {
			fmt.Println("ошибка сериализации: %w", err)
			return
		}

		tempFile.Write(append(updatedLine, '\n'))

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла tasks:", err)
		return
	}

	if !found {
		fmt.Printf("В файле не найдена задача с идентификатором %d ", id)
		return
	}

	tempFile.Close()
	file.Close()

	err = os.Rename(tempFile.Name(), "tasks.json")

	if err != nil {
		fmt.Printf("не удалось заменить файл tasks.json", err)
		return
	}
}

func DeleteTaskInFile(id int) {

	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл tasks.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаём временный файл для записи обновлённых данных
	tempFile, err := os.CreateTemp("", "jsonl_*.tmp")
	if err != nil {
		fmt.Println("не удалось создать временный файл: %w", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Bytes()
		var task task.Task

		err := json.Unmarshal(line, &task)

		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", line, err)
			tempFile.WriteString(string(line) + "\n")
			continue
		}

		if task.Taskid == id {
			found = true
			continue
		}

		updatedLine, err := json.Marshal(task)
		if err != nil {
			fmt.Println("ошибка сериализации: %w", err)
			return
		}

		tempFile.Write(append(updatedLine, '\n'))

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла tasks:", err)
		return
	}

	if !found {
		fmt.Printf("В файле не найдена задача с идентификатором %d ", id)
		return
	}

	tempFile.Close()
	file.Close()

	err = os.Rename(tempFile.Name(), "tasks.json")

	if err != nil {
		fmt.Printf("не удалось заменить файл tasks.json", err)
		return
	}
}

func UpdateUserInFile(id int, newUser taskUser.UpdateUserRequest) {

	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл users.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаём временный файл для записи обновлённых данных
	tempFile, err := os.CreateTemp("", "jsonl_*.tmp")
	if err != nil {
		fmt.Println("не удалось создать временный файл: %w", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Bytes()
		var user taskUser.User

		err := json.Unmarshal(line, &user)

		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", line, err)
			tempFile.WriteString(string(line) + "\n")
			continue
		}

		if user.Userid == id {
			if newUser.FirstName != nil {
				user.FirstName = *newUser.FirstName
			}
			if newUser.LastName != nil {
				user.LastName = *newUser.LastName
			}
			found = true
		}

		updatedLine, err := json.Marshal(user)
		if err != nil {
			fmt.Println("ошибка сериализации: %w", err)
			return
		}

		tempFile.Write(append(updatedLine, '\n'))

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла users:", err)
		return
	}

	if !found {
		fmt.Printf("В файле не найден пользователь с идентификатором %d ", id)
		return
	}

	tempFile.Close()
	file.Close()

	err = os.Rename(tempFile.Name(), "users.json")

	if err != nil {
		fmt.Printf("не удалось заменить файл users.json", err)
		return
	}
}

func DeleteUserInFile(id int) {
	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл users.json не найден!")
			return
		}
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаём временный файл для записи обновлённых данных
	tempFile, err := os.CreateTemp("", "jsonl_*.tmp")
	if err != nil {
		fmt.Println("не удалось создать временный файл: %w", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Bytes()
		var user taskUser.User

		err := json.Unmarshal(line, &user)

		if err != nil {
			fmt.Println("Ошибка при разборе JSON:", line, err)
			tempFile.WriteString(string(line) + "\n")
			continue
		}

		if user.Userid == id {
			found = true
			continue
		}

		updatedLine, err := json.Marshal(user)
		if err != nil {
			fmt.Println("ошибка сериализации: %w", err)
			return
		}

		tempFile.Write(append(updatedLine, '\n'))

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла users:", err)
		return
	}

	if !found {
		fmt.Printf("В файле не найден пользователь с идентификатором %d ", id)
		return
	}

	tempFile.Close()
	file.Close()

	err = os.Rename(tempFile.Name(), "users.json")

	if err != nil {
		fmt.Printf("не удалось заменить файл users.json", err)
		return
	}
}

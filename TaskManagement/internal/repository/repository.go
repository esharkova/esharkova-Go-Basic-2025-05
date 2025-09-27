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

var Users []*taskUser.User
var Tasks []*task.Task

var MuUser sync.Mutex
var MuTask sync.Mutex

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
}

func AddTask(task task.Task) {
	MuTask.Lock()
	defer MuTask.Unlock()
	Tasks = append(Tasks, &task)
	fmt.Println("Добавлена задача в слайс")
}

func WriteUserToFile(user taskUser.User) {

	MuUser.Lock()
	defer MuUser.Unlock()

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

	MuTask.Lock()
	defer MuTask.Unlock()

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

}

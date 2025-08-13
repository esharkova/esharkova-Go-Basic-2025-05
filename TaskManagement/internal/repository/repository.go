package repository

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var Users []*taskUser.User
var Tasks []*task.Task

var mu sync.Mutex

type TaskManage interface {
	Insert() int
}

func ProcessValue(tm TaskManage) {

	switch v := tm.(type) {

	case taskUser.User:
		fmt.Println("Processing User", v)
		Users = append(Users, &v)
	case task.Task:
		fmt.Println("Processing Task", v)
		Tasks = append(Tasks, &v)
	default:
		fmt.Println("Unknown type", v)

	}

}

func PrintSlice() {
	fmt.Println("Total users: ", strconv.Itoa(len(Users)), Users)
	fmt.Println("Total tasks: ", strconv.Itoa(len(Tasks)), Tasks)
}

func AddUser(user taskUser.User) {
	mu.Lock()
	defer mu.Unlock()
	Users = append(Users, &user)
}

func AddTask(task task.Task) {
	mu.Lock()
	defer mu.Unlock()
	Tasks = append(Tasks, &task)
}

func LogSlices() {

	var lastUsersCount, lastTasksCount int
	var lastUserSlice []*taskUser.User
	var lastTasksSlice []*task.Task

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		mu.Lock()
		currentUsersCount := len(Users)
		currentTasksCount := len(Tasks)

		if currentUsersCount != lastUsersCount {
			newUsers := Users[lastUsersCount:currentUsersCount]
			log.Printf("Добавились пользователи: %+v\n", newUsers)
			lastUsersCount = currentUsersCount
			lastUserSlice = append(lastUserSlice, newUsers...)
		}

		if currentTasksCount != lastTasksCount {
			newTasks := Tasks[lastTasksCount:currentTasksCount]
			log.Printf("Добавились задачи: %+v\n", newTasks)
			lastTasksCount = currentTasksCount
			lastTasksSlice = append(lastTasksSlice, newTasks...)
		}

		mu.Unlock()
	}

}

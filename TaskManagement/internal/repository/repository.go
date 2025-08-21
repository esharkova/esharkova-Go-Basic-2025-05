package repository

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"errors"
	"fmt"
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
	copiedUsers := make([]*taskUser.User, len(currentUsers))
	copy(copiedUsers, currentUsers)
	return copiedUsers

}

func GetCopyTasks(currentTasks []*task.Task) []*task.Task {
	copiedTasks := make([]*task.Task, len(currentTasks))
	copy(copiedTasks, currentTasks)
	return copiedTasks

}

func AddUser(user taskUser.User) {
	MuUser.Lock()
	defer MuUser.Unlock()
	Users = append(Users, &user)
}

func AddTask(task task.Task) {
	MuTask.Lock()
	defer MuTask.Unlock()
	Tasks = append(Tasks, &task)
}

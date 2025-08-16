package repository

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
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

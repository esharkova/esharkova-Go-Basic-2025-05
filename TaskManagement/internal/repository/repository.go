package repository

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"fmt"
	"strconv"
)

var Users []*taskUser.User
var Tasks []*task.Task

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

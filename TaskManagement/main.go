package main

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	service "TaskManagement/internal/service"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	service.CreateItems()

}

func TerminalCreating() {

	var FirstName string
	var LastName string
	var Passport string
	var TaskDescription string

	WorkStatus := task.Status{1, "В работе"}

	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&FirstName)

	fmt.Print("Введите фамилию пользователя: ")
	fmt.Scan(&LastName)

	fmt.Print("Введите документ пользователя: ")
	fmt.Scan(&Passport)

	newUser := taskUser.User{}
	newUser.Userid = rand.Intn(100)
	newUser.FirstName = FirstName
	newUser.LastName = LastName
	newUser.AddPassport(Passport)

	fmt.Println("Добавлен пользователь: ", newUser.FirstName, newUser.LastName, newUser.GetPassport())

	fmt.Print("Введите задачу для пользователя: ")
	fmt.Scan(&TaskDescription)

	newTask := task.Task{
		Taskid:             rand.Intn(100),
		TaskNumber:         strconv.Itoa(rand.Intn(100000)),
		Description:        TaskDescription,
		CreateDateTime:     time.Now(),
		CompletionDateTime: time.Now().AddDate(0, 0, 7),
	}

	newTaskStatus := task.TaskStatus{
		Taskid:        newTask.Taskid,
		Userid:        newUser.Userid,
		Statusid:      WorkStatus.Statusid,
		StartDateTime: time.Now(),
	}

	fmt.Println("Добавлена задача: ", newTask.Description, " Срок выполнения ", newTask.CompletionDateTime.Format(time.DateOnly), " Статус: ", WorkStatus.StatusName, " с ", newTaskStatus.StartDateTime.Format(time.DateOnly))

}

package main

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"TaskManagement/internal/repository"
	"TaskManagement/internal/service"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	taskChannel := make(chan task.Task, 10)
	userChannel := make(chan taskUser.User, 10)

	go service.LogSlices()

	go func() {

		ticker1 := time.NewTicker(5 * time.Second)
		defer ticker1.Stop()

		for range ticker1.C {
			taskChannel <- service.CreateTask()
			fmt.Println("Add Task in Task Channel")

		}

	}()

	go func() {
		ticker2 := time.NewTicker(1 * time.Second)
		defer ticker2.Stop()

		for range ticker2.C {
			userChannel <- service.CreateUser()
			fmt.Println("Add User in User Channel")

		}
	}()

	fmt.Println("Горутина добавления объектов в каналы завершила свое выполнение")

	go func() {

		for val := range userChannel {

			repository.AddUser(val)
		}

	}()

	go func() {

		for val := range taskChannel {

			repository.AddTask(val)
		}

	}()

	time.Sleep(30 * time.Second)

	service.PrintSlice()

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

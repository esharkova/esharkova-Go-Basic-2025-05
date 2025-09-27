package main

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	"TaskManagement/internal/repository"
	"TaskManagement/internal/service"
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {

	repository.ReadUsersFromFileAndAddToSlice()
	repository.ReadTasksFromFileAndAddToSlice()

	taskChannel := make(chan task.Task, 10)
	userChannel := make(chan taskUser.User, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отменяем контекст при завершении функции main

	var wg sync.WaitGroup

	wg.Add(6)

	go service.LogSlices(ctx, &wg)

	go func(ctx context.Context) {

		ticker1 := time.NewTicker(3 * time.Second)
		defer ticker1.Stop()

		for {
			select {
			case <-ticker1.C:
				taskChannel <- service.CreateTask()
				fmt.Println("Добавлена задача в канал")
			case <-ctx.Done():
				fmt.Println(ctx.Err().Error())
				fmt.Println("Горутина добавления задач получила отмену контекста")
				wg.Done()
				return
			}
		}

	}(ctx)

	go func(ctx context.Context) {
		ticker2 := time.NewTicker(1 * time.Second)
		defer ticker2.Stop()

		for {
			select {
			case <-ticker2.C:
				userChannel <- service.CreateUser()
				fmt.Println("Добавлен пользователь в канал")
			case <-ctx.Done():
				fmt.Println(ctx.Err().Error())
				fmt.Println("Горутина добавления пользователей получила отмену контекста")
				wg.Done()
				return
			}
		}

	}(ctx)

	go func(ctx context.Context) {

		for {
			select {
			case user := <-userChannel:
				repository.AddUser(user)
				repository.WriteUserToFile(user)
			case <-ctx.Done():
				fmt.Println(ctx.Err().Error())
				fmt.Println("Горутина добавления пользователей в слайс получила отмену контекста")
				wg.Done()
				return
			}
		}

	}(ctx)

	go func(ctx context.Context) {

		for {
			select {
			case task := <-taskChannel:
				repository.AddTask(task)
				repository.WriteTaskToFile(task)
			case <-ctx.Done():
				fmt.Println(ctx.Err().Error())
				fmt.Println("Горутина добавления задач в слайс получила отмену контекста")
				wg.Done()
				return
			}
		}

	}(ctx)

	go gracefulShutdown(cancel, &wg)

	wg.Wait()

	fmt.Println("Все горутины завершили выполнение")
	service.PrintSlice()

}

func gracefulShutdown(cancel context.CancelFunc, wg *sync.WaitGroup) {
	// Создаем канал для получения сигналов
	sigs := make(chan os.Signal, 1)

	// Уведомляем канал о поступлении указанных сигналов
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Ожидаем сигнал (нажмите Ctrl+C)...")
	sig := <-sigs

	cancel()

	fmt.Println("Получен сигнал:", sig)
	fmt.Println("Выход из программы.")

	wg.Done()
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

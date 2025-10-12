package service

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	repository "TaskManagement/internal/repository"
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func CreateItems() {

	var item repository.TaskManage
	var array = [10]int{}

	for i := 0; i < len(array); i++ {

		array[i] = rand.Intn(30)

	}

	fmt.Println("Массив случайных чисел ", array)

	//добавляем в случайном порядке разные структуры
	for _, i := range array {

		//если четное значение элемента массива, то добавляем задачу
		if i%2 == 0 {

			newTask := task.Task{
				Taskid:             i,
				TaskNumber:         "TaskNumber" + strconv.Itoa(i),
				Description:        "TaskDescription" + strconv.Itoa(i),
				CreateDateTime:     time.Now(),
				CompletionDateTime: time.Now().AddDate(0, 0, 7),
			}

			item = newTask
			// иначе пользователя
		} else {

			newUser := taskUser.User{
				Userid:    i,
				FirstName: "FirstName" + strconv.Itoa(i),
				LastName:  "LastName" + strconv.Itoa(i),
			}

			newUser.AddPassport("Passport" + strconv.Itoa(i))

			item = newUser

		}

		item.Insert()
		repository.ProcessValue(item)

	}

	PrintSlice()
}

func CreateUser() taskUser.User {

	i := rand.Intn(100)

	newUser := taskUser.User{
		Userid:    i,
		FirstName: "FirstName" + strconv.Itoa(i),
		LastName:  "LastName" + strconv.Itoa(i),
	}

	newUser.AddPassport("Passport" + strconv.Itoa(i))

	return newUser

}

func CreateTask() task.Task {

	i := rand.Intn(100)

	newTask := task.Task{
		Taskid:             i,
		TaskNumber:         "TaskNumber" + strconv.Itoa(i),
		Description:        "TaskDescription" + strconv.Itoa(i),
		CreateDateTime:     time.Now(),
		CompletionDateTime: time.Now().AddDate(0, 0, 7),
	}

	return newTask

}

func LogSlices(ctx context.Context, wg *sync.WaitGroup) {

	var lastUsersCount, lastTasksCount int

	var lastUserSlice []*taskUser.User
	var lastTasksSlice []*task.Task

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			copiedUsers := repository.GetCopyUsers(repository.Users)
			copiedTasks := repository.GetCopyTasks(repository.Tasks)

			lastUsersCount = repository.GetLastUserCount()
			lastTasksCount = repository.GetLastTasksCount()

			currentUsersCount := len(copiedUsers)
			currentTasksCount := len(copiedTasks)

			if currentUsersCount != lastUsersCount {
				newUsers := copiedUsers[lastUsersCount:currentUsersCount]
				log.Printf("Добавились пользователи: %+v\n", newUsers)
				lastUserSlice = append(lastUserSlice, newUsers...)
				lastUsersCount = currentUsersCount
				repository.SetLastUserCount(lastUsersCount)

			}

			if currentTasksCount != lastTasksCount {
				newTasks := copiedTasks[lastTasksCount:currentTasksCount]
				log.Printf("Добавились задачи: %+v\n", newTasks)
				lastTasksSlice = append(lastTasksSlice, newTasks...)
				lastTasksCount = currentTasksCount
				repository.SetLastTasksCount(lastTasksCount)

			}
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			fmt.Println("Горутина логирования получила отмену контекста")
			wg.Done()
			return
		}
	}

}

func PrintSlice() {

	copiedUsers := repository.GetCopyUsers(repository.Users)
	copiedTasks := repository.GetCopyTasks(repository.Tasks)

	fmt.Println("Total users: ", strconv.Itoa(len(copiedUsers)))
	for _, user := range copiedUsers {
		fmt.Printf("  %+v\n", *user)
	}
	fmt.Println("Total tasks: ", strconv.Itoa(len(copiedTasks)))
	for _, task := range copiedTasks {
		fmt.Printf("  %+v\n", *task)
	}

}

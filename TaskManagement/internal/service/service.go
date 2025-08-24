package service

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	repository "TaskManagement/internal/repository"
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

func LogSlices() {

	var lastUsersCount, lastTasksCount int
	var lastUserSlice []*taskUser.User
	var lastTasksSlice []*task.Task

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {

		copiedUsers := repository.GetCopyUsers(repository.Users)
		copiedTasks := repository.GetCopyTasks(repository.Tasks)

		currentUsersCount := len(copiedUsers)
		currentTasksCount := len(copiedTasks)

		if currentUsersCount != lastUsersCount {
			newUsers := copiedUsers[lastUsersCount:currentUsersCount]
			log.Printf("Добавились пользователи: %+v\n", newUsers)
			lastUsersCount = currentUsersCount
			lastUserSlice = append(lastUserSlice, newUsers...)
		}

		if currentTasksCount != lastTasksCount {
			newTasks := copiedTasks[lastTasksCount:currentTasksCount]
			log.Printf("Добавились задачи: %+v\n", newTasks)
			lastTasksCount = currentTasksCount
			lastTasksSlice = append(lastTasksSlice, newTasks...)
		}

	}

}

func PrintSlice() {

	copiedUsers := repository.GetCopyUsers(repository.Users)
	copiedTasks := repository.GetCopyTasks(repository.Tasks)

	fmt.Println("Total users: ", strconv.Itoa(len(copiedUsers)), copiedUsers)
	fmt.Println("Total tasks: ", strconv.Itoa(len(copiedTasks)), copiedTasks)

}

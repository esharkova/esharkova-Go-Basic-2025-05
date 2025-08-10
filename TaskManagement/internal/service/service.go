package service

import (
	task "TaskManagement/internal/model/task"
	taskUser "TaskManagement/internal/model/user"
	repository "TaskManagement/internal/repository"
	"fmt"
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

	repository.PrintSlice()
}

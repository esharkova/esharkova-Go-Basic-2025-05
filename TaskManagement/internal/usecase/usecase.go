package usecase

import (
	"TaskManagement/internal/model/task"
	user "TaskManagement/internal/model/user"
	"TaskManagement/internal/repository"
	"context"
	"math/rand"
	"time"
)

type UseCase struct {
}

func (u *UseCase) GetTasks() []*task.Task {

	return repository.GetCopyTasks(repository.Tasks)

}

func (u *UseCase) GetUsers() []*user.User {

	return repository.GetCopyUsers(repository.Users)

}

func (u *UseCase) CreateTask(context context.Context, args task.CreateTaskRequest) task.Task {

	i := rand.Intn(100)

	newTask := task.Task{
		Taskid:             i,
		TaskNumber:         args.TaskNumber,
		Description:        args.Description,
		CreateDateTime:     time.Now(),
		CompletionDateTime: time.Now().AddDate(0, 0, 7),
		Priority:           args.Priority,
	}
	repository.AddTask(newTask)

	return newTask
}

func (u *UseCase) GetTask(id int) (task.Task, error) {

	return repository.GetTaskById(id)

}

func (u *UseCase) UpdateTask(context context.Context, args task.UpdateTaskRequest, id int) (task.Task, error) {

	updTask, err := repository.UpdateTask(id, args)

	return updTask, err

}

func (u *UseCase) DeleteTask(context context.Context, id int) error {

	err := repository.DeleteTask(id)

	return err

}

func (u *UseCase) GetUser(id int) (user.User, error) {

	return repository.GetUserById(id)

}

func (u *UseCase) CreateUser(context context.Context, args user.CreateUserRequest) user.User {

	i := rand.Intn(100)

	newUser := user.User{
		Userid:    i,
		FirstName: args.FirstName,
		LastName:  args.LastName,
	}

	repository.AddUser(newUser)

	return newUser

}

func (u *UseCase) UpdateUser(context context.Context, args user.UpdateUserRequest, id int) (user.User, error) {

	updUser, err := repository.UpdateUser(id, args)

	return updUser, err
}

func (u *UseCase) DeleteUser(context context.Context, id int) error {

	err := repository.DeleteUser(id)

	return err
}

func New() *UseCase {
	return &UseCase{}
}

package usecase

import (
	"TaskManagement/db"
	dbtask "TaskManagement/db"
	"TaskManagement/internal/model/task"
	user "TaskManagement/internal/model/user"
	"context"
	"database/sql"
	"log"
	"time"
)

//go:generate mockgen -source=usecase.go -destination=mock/usecase.go -package=mocks

type UseCase struct {
	queries db.Querier
}

func (u *UseCase) GetTasks() []*task.Task {

	var items []dbtask.Task

	items, err := u.queries.ListTasks(context.Background())

	if err != nil {
		return nil
	}
	log.Printf("items: %+v\n", items)

	return mapDBTasksToTasks(items)

}

func (u *UseCase) GetUsers() []*user.User {

	var items []dbtask.User

	items, _ = u.queries.ListUsers(context.Background())

	return MapDBUsersToUsers(items)

}

func (u *UseCase) CreateTask(context context.Context, args task.CreateTaskRequest) error {

	error := u.queries.CreateTask(context, MapCreateTaskRequestToParams(args))

	return error
}

func (u *UseCase) GetTask(id int) (task.Task, error) {

	task, err := u.queries.GetTaskByID(context.Background(), int32(id))

	return *mapDBTaskToTask(task), err

}

func (u *UseCase) UpdateTask(context context.Context, args task.UpdateTaskRequest, id int) (task.Task, error) {

	updTask, err := u.queries.UpdateTask(context, MapUpdateTaskRequestToParams(id, args))

	return *mapDBTaskToTask(updTask), err

}

func (u *UseCase) DeleteTask(context context.Context, id int) error {

	err := u.queries.DeleteTask(context, int32(id))

	return err

}

func (u *UseCase) GetUser(id int) (user.User, error) {

	user, err := u.queries.GetUserByID(context.Background(), int32(id))

	return *MapDBUserToUser(user), err

}

func (u *UseCase) CreateUser(context context.Context, args user.CreateUserRequest) error {

	err := u.queries.CreateUser(context, MapCreateUserRequestToParams(args))

	return err

}

func (u *UseCase) UpdateUser(context context.Context, args user.UpdateUserRequest, id int) (user.User, error) {

	updUser, err := u.queries.UpdateUser(context, MapUpdateUserRequestToParams(id, args))

	return *MapDBUserToUser(updUser), err
}

func (u *UseCase) DeleteUser(context context.Context, id int) error {

	err := u.queries.DeleteUser(context, int32(id))

	return err
}

func mapDBTaskToTask(dbTask db.Task) *task.Task {
	return &task.Task{
		Taskid:             int(dbTask.Taskid),
		TaskNumber:         dbTask.TaskNumber,
		Description:        dbTask.Description.String,
		CreateDateTime:     dbTask.CreateDatetime,
		CompletionDateTime: dbTask.CompletionDatetime.Time,
		Priority:           int(dbTask.Priority.Int32),
	}
}

func mapDBTasksToTasks(dbTasks []db.Task) []*task.Task {
	tasks := make([]*task.Task, len(dbTasks))
	for i, t := range dbTasks {
		tasks[i] = mapDBTaskToTask(t)
	}
	return tasks
}

func MapCreateTaskRequestToParams(req task.CreateTaskRequest) dbtask.CreateTaskParams {
	return dbtask.CreateTaskParams{
		TaskNumber: req.TaskNumber,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
		CreateDatetime: time.Now(),
		CompletionDatetime: sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		},
		Priority: sql.NullInt32{
			Int32: int32(req.Priority),
			Valid: true,
		},
	}
}

func MapUpdateTaskRequestToParams(id int, req task.UpdateTaskRequest) dbtask.UpdateTaskParams {

	taskNumber := *req.TaskNumber

	description := sql.NullString{}
	if req.Description != nil {
		description = sql.NullString{
			String: *req.Description,
			Valid:  true,
		}
	}

	priority := sql.NullInt32{}
	if req.Priority != nil {
		priority = sql.NullInt32{
			Int32: int32(*req.Priority),
			Valid: true,
		}
	}

	return dbtask.UpdateTaskParams{
		Taskid:      int32(id),
		TaskNumber:  taskNumber,
		Description: description,
		Priority:    priority,
	}
}

func MapDBUserToUser(dbUser db.User) *user.User {

	u := &user.User{
		Userid:    int(dbUser.Userid),
		FirstName: dbUser.Firstname,
		LastName:  dbUser.Lastname,
	}

	u.AddPassport(dbUser.PassportNumber)

	return u
}

func MapDBUsersToUsers(dbUsers []db.User) []*user.User {
	users := make([]*user.User, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = MapDBUserToUser(u)
	}
	return users
}

func MapCreateUserRequestToParams(req user.CreateUserRequest) dbtask.CreateUserParams {
	return dbtask.CreateUserParams{
		Firstname:      req.FirstName,
		Lastname:       req.LastName,
		PassportNumber: "",
	}
}

func MapUpdateUserRequestToParams(id int, req user.UpdateUserRequest) dbtask.UpdateUserParams {
	var firstname, lastname string

	if req.FirstName != nil {
		firstname = *req.FirstName
	}
	if req.LastName != nil {
		lastname = *req.LastName
	}

	return dbtask.UpdateUserParams{
		Userid:    int32(id),
		Firstname: firstname,
		Lastname:  lastname,
	}
}

func New(queries db.Querier) *UseCase {
	return &UseCase{queries: queries}
}

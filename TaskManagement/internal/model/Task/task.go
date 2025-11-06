package task

import (
	"fmt"
	"time"
)

/*Задача */
type Task struct {
	Taskid             int       /*идентификатор*/
	TaskNumber         string    /*номер задачи*/
	Description        string    /*описание*/
	CreateDateTime     time.Time /*дата создания*/
	CompletionDateTime time.Time /*дата выполнения*/
	Priority           int       /*приоритет*/
}

type CreateTaskRequest struct {
	TaskNumber  string `json:"taskNumber" binding:"required,min=2,max=10"`
	Description string `json:"description" binding:"required"`
	Priority    int    `json:"priority" binding:"required"`
}
type UpdateTaskRequest struct {
	TaskNumber  *string `json:"taskNumber" binding:"min=2,max=10"`
	Description *string `json:"description"`
	Priority    *int    `json:"priority"`
}

/*Status - справочник возможных статусов задачи (н-р, "планируется", "в процессе", "завершено")*/
type Status struct {
	Statusid   int
	StatusName string
}

/*Статус задачи*/
type TaskStatus struct {
	Taskid        int
	Userid        int
	Statusid      int
	StartDateTime time.Time /*дата начала установки статуса*/
	EndDateTime   time.Time /*дата окончания установки статуса*/

}

func (t Task) Insert() int {

	fmt.Println("Добавлена задача ", t.Taskid)

	return t.Taskid

}

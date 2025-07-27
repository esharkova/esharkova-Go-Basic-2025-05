package task

import (
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

package db

import "time"

type Status string

/*
	export const D_TASK_STATUS = {
	    DONE: "done",
	    IN_PROGRESS: "in_progress",
	    TODO: "todo",
	    BACKLOG: "backlog",
	};
*/
const (
	DONE        Status = "done"
	IN_PROGRESS Status = "in_progress"
	TODO        Status = "todo"
	BACKLOG     Status = "backlog"
)

/*
	{
		id: "",
	    title: `Task`,
	    description: null,
	    status: STATUS.BACKLOG,
	    flag: null,
	    links: [], //Links to other tasks
	};
*/
type Task struct {
	Id          string  `gorm:"primarykey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      Status  `json:"status"`
	Flag        *string `json:"flag"`

	//LINKS

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ProjectId string
	Project   Project
}

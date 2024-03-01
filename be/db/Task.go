package db

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Status string

type StatusWrapper struct {
	Status Status `json:"status" binding:"required,oneof=done in_progress todo backlog"`
}

const (
	DONE        Status = "done"
	IN_PROGRESS Status = "in_progress"
	TODO        Status = "todo"
	BACKLOG     Status = "backlog"
)

type Task struct {
	Id          string  `gorm:"primarykey" json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status      Status  `json:"status" binding:"required,oneof=done in_progress todo backlog"`
	Flag        *string `json:"flag"`
	Archived    bool    `json:"archived"`

	Tags []*Tag `gorm:"many2many:task_tags;constraint:OnDelete:CASCADE" json:"tags,omitempty"`

	//LINKS

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ProjectId string  `json:"project_id"`
	Project   Project `json:"-"`
}

type TaskCreate struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Status      Status  `json:"status" binding:"required,oneof=done in_progress todo backlog"`
	Flag        *string `json:"flag"`
	Tags        []*Tag  `json:"tags" binding:"required"`
}

func (t *TaskCreate) Task(projectId string) Task {
	for _, t := range t.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}
	return Task{
		Id:          ulid.Make().String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Flag:        t.Flag,
		Tags:        t.Tags,
		ProjectId:   projectId,
	}
}

type TaskUpdate struct {
	TaskCreate
	Id       string `json:"id" binding:"required"`
	Archived bool   `json:"archived"`
}

func (t *TaskUpdate) Task(projectId string) Task {
	for _, t := range t.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}
	return Task{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Flag:        t.Flag,
		Tags:        t.Tags,
		ProjectId:   projectId,
		Archived:    t.Archived,
	}
}

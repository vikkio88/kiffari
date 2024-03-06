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

type Category string

const (
	FEATURE Category = "feature"
	BUG     Category = "bug"
	DOC     Category = "doc"
	SPIKE   Category = "spike"
	CLEANUP Category = "cleanup"
)

type Task struct {
	Id          string   `gorm:"primarykey" json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      Status   `json:"status" binding:"required,oneof=done in_progress todo backlog"`
	Category    Category `json:"category" binding:"required,oneof=feature bug doc spike cleanup"`
	Priority    int      `json:"priority"`
	Flag        *string  `json:"flag"`
	Archived    bool     `json:"archived"`

	Tags []*Tag `gorm:"many2many:task_tags;constraint:OnDelete:CASCADE" json:"tags,omitempty"`

	//LINKS
	// https://gorm.io/docs/has_many.html#Self-Referential-Has-Many

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ProjectId string  `json:"project_id"`
	Project   Project `json:"project,omitempty"`
}

type TaskCreate struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Status      Status   `json:"status" binding:"required,oneof=done in_progress todo backlog"`
	Category    Category `json:"category" binding:"required,oneof=feature bug doc spike cleanup"`
	Priority    int      `json:"priority"`
	Flag        *string  `json:"flag"`
	Tags        []*Tag   `json:"tags" binding:"required"`
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
		Category:    t.Category,
		Priority:    t.Priority,
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
		Category:    t.Category,
		Priority:    t.Priority,
		Flag:        t.Flag,
		Tags:        t.Tags,
		ProjectId:   projectId,
		Archived:    t.Archived,
	}
}

package db

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Note struct {
	Id        string     `gorm:"primarykey;size:16" json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Tags      []*Tag     `gorm:"many2many:note_tags;constraint:OnDelete:CASCADE" json:"tags,omitempty"`
	DueDate   *time.Time `json:"due_date"`
	Archived  bool       `json:"archived"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type NoteItem struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	DueDate   *time.Time `json:"due_date"`
	Archived  bool       `json:"archived"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type NoteUpdate struct {
	NoteCreate
	Id       string `json:"id" binding:"required"`
	Archived bool   `json:"archived"`
}

func (n NoteUpdate) Note() Note {
	for _, t := range n.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}

	return Note{
		Id:       n.Id,
		Title:    n.Title,
		Body:     n.Body,
		DueDate:  n.DueDate,
		Archived: n.Archived,
		Tags:     n.Tags,
	}

}

type NoteCreate struct {
	Title   string     `json:"title" binding:"required"`
	Body    string     `json:"body" binding:"required"`
	DueDate *time.Time `json:"due_date" binding:"omitempty,dateInTheFuture" time_format:"2006-01-02T15:04:05Z07:00"`
	Tags    []*Tag     `json:"tags" binding:"required"`
}

func (n NoteCreate) Note() Note {
	for _, t := range n.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}

	return Note{
		Id:      ulid.Make().String(),
		Title:   n.Title,
		Body:    n.Body,
		DueDate: n.DueDate,
		Tags:    n.Tags,
	}

}

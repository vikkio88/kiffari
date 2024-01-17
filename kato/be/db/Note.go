package db

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Note struct {
	Id        string    `gorm:"primarykey;size:16" json:"id"`
	Body      string    `json:"body"`
	Tags      []*Tag    `gorm:"many2many:note_tags;constraint:OnDelete:CASCADE" json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteItem struct {
	Id        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteUpdate struct {
	NoteCreate
	Id string `json:"id" binding:"required"`
}

func (n NoteUpdate) Note() Note {
	for _, t := range n.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}

	return Note{
		Id:   n.Id,
		Body: n.Body,
		Tags: n.Tags,
	}

}

type NoteCreate struct {
	Body string `json:"body" binding:"required"`
	Tags []*Tag `json:"tags" binding:"required"`
}

func (n NoteCreate) Note() Note {
	for _, t := range n.Tags {
		if t.Id == "" {
			t.Id = ulid.Make().String()
			t.Value = normaliseTag(t.Label)
		}
	}

	return Note{
		Id:   ulid.Make().String(),
		Body: n.Body,
		Tags: n.Tags,
	}

}

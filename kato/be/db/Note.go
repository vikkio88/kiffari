package db

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Note struct {
	Id        string    `gorm:"primarykey;size:16" json:"id"`
	Content   string    `json:"content" `
	Tags      []Tag     `gorm:"many2many:note_tags;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteCreate struct {
	Content string `json:"content" binding:"required"`
	Tags    []Tag  `json:"tags" binding:"required"`
}

func (n NoteCreate) Note() Note {
	return Note{
		Id:      ulid.Make().String(),
		Content: n.Content,
		Tags:    n.Tags,
	}

}

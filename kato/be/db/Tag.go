package db

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

type Tag struct {
	Id    string  `gorm:"primarykey" json:"id" binding:"required"`
	Value string  `gorm:"uniqueIndex" json:"-"`
	Label string  `json:"label"`
	Notes []*Note `gorm:"many2many:note_tags;constraint:OnDelete:CASCADE" json:"notes,omitempty"`
}

type TagCreate struct {
	Label string `json:"label" binding:"required"`
}

func (t *TagCreate) Tag() Tag {
	value := t.Label
	return Tag{
		Id:    ulid.Make().String(),
		Value: normaliseTag(value),
		Label: value,
	}
}

func normaliseTag(value string) string {
	return strings.ToLower(value)
}

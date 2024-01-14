package db

import (
	"fmt"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct {
	g *gorm.DB
}

func NewDb(fileName string) *Db {
	g, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?_foreign_keys=on", fileName)), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(g)
	return &Db{g}
}

func (d *Db) InsertTag(value TagCreate) bool {
	t := value.Tag()
	tx := d.g.Save(t)

	if tx.Error != nil && strings.Contains(tx.Error.Error(), "UNIQUE") {
		ot := d.GetTagByValue(t.Value)
		ot.Label = t.Label
		tx := d.g.Save(ot)
		return tx.Error == nil
	}

	return true
}

func (d *Db) GetTagByValue(value string) Tag {
	var t Tag
	d.g.Model(&Tag{}).Where("value", value).Find(&t)

	return t
}

func (d *Db) GetAllTags() []Tag {
	var tags []Tag
	d.g.Model(&Tag{}).Find(&tags)

	return tags
}

func (d *Db) FilterTags(value string) []Tag {
	var tags []Tag
	value = normaliseTag(value)
	d.g.Model(&Tag{}).Where("value LIKE ?", fmt.Sprintf("%%%s%%", value)).Find(&tags)

	return tags
}

func (d *Db) InsertNote(value NoteCreate) bool {
	d.g.Save(value.Note())
	return true
}

func (d *Db) GetAllNotes() []Note {
	var notes []Note

	d.g.Model(&Note{}).Preload("Tags").Find(&notes)

	return notes
}

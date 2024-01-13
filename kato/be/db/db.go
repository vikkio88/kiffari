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
	t := value.NewTag()
	tx := d.g.Save(t)

	if tx.Error != nil && strings.Contains(tx.Error.Error(), "UNIQUE") {
		ot := d.GetByValue(t.Value)
		ot.Label = t.Label
		tx := d.g.Save(ot)
		return tx.Error == nil
	}

	return true

}

func (d *Db) GetByValue(value string) Tag {
	var t Tag
	d.g.Model(&Tag{}).Where("value", value).Find(&t)

	return t
}

func (d *Db) GetAll() []Tag {
	var tags []Tag
	d.g.Model(&Tag{}).Find(&tags)

	return tags
}

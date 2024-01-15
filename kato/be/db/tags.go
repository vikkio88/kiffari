package db

import (
	"fmt"
	"strings"
)

func (d *Db) InsertTag(value TagCreate) (Tag, bool) {
	t := value.Tag()
	tx := d.g.Save(t)

	if tx.Error != nil && strings.Contains(tx.Error.Error(), "UNIQUE") {
		ot := d.GetTagByValue(t.Value)
		ot.Label = t.Label
		tx := d.g.Save(ot)

		return ot, tx.Error == nil

	}

	return t, true
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

package db

func (d *Db) InsertNote(value NoteCreate) (Note, bool) {
	n := value.Note()
	trx := d.g.Create(&n)
	return n, trx.RowsAffected == 1
}

func (d *Db) GetLatest() []NoteItem {
	var notes []NoteItem

	d.g.Model(&Note{}).Order("updated_at DESC, created_at DESC").Where("archived", false).Limit(5).Find(&notes)

	return notes
}

func (d *Db) GetAllNotes() []NoteItem {
	var notes []NoteItem
	d.g.Model(&Note{}).Order("updated_at DESC, created_at DESC").Find(&notes)

	return notes
}

func (d *Db) GetArchivedNotes() []NoteItem {
	var notes []NoteItem

	d.g.Model(&Note{}).Order("updated_at DESC, created_at DESC").Where("archived", true).Find(&notes)

	return notes
}

func (d *Db) GetNoteById(id string) (Note, bool) {
	var note Note

	trx := d.g.Model(&Note{}).Preload("Tags").Find(&note, "Id = ?", id)

	return note, trx.RowsAffected == 1
}

func (d *Db) UpdateNote(n NoteUpdate) (string, bool) {
	un := n.Note()

	d.g.Model(&un).Association("Tags").Replace(un.Tags)
	trx := d.g.Omit("created_at").Save(&un)
	if trx.RowsAffected != 1 {
		return "", false
	}
	return n.Id, true
}

func (d *Db) SetArchivedNote(id string, archived bool) bool {
	trx := d.g.Model(&Note{}).Where("id", id).Update("archived", archived)

	return trx.RowsAffected == 1
}

func (d *Db) DeleteNote(id string) bool {
	trx := d.g.Where("id", id).Delete(&Note{})

	return trx.RowsAffected > 0
}

package db

func (d *Db) InsertNote(value NoteCreate) (Note, bool) {
	n := value.Note()
	trx := d.g.Create(&n)
	return n, trx.RowsAffected == 1
}

func (d *Db) GetAllNotes() []NoteItem {
	var notes []NoteItem

	d.g.Model(&Note{}).Find(&notes)

	return notes
}

func (d *Db) GetNoteById(id string) (Note, bool) {
	var note Note

	trx := d.g.Model(&Note{}).Preload("Tags").Find(&note, "Id = ?", id)

	return note, trx.RowsAffected == 1
}

func (d *Db) UpdateNote(n NoteUpdate) (string, bool) {
	un := n.Note()
	trx := d.g.Save(un)
	if trx.RowsAffected != 1 {
		return "", false
	}
	return n.Id, true
}

func (d *Db) DeleteNote(id string) bool {
	trx := d.g.Where("id", id).Delete(&Note{})

	return trx.RowsAffected > 0
}

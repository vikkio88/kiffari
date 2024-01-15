package db

func (d *Db) InsertNote(value NoteCreate) string {
	n := value.Note()
	d.g.Create(&n)
	return n.Id
}

func (d *Db) GetAllNotes() []Note {
	var notes []Note

	d.g.Model(&Note{}).Preload("Tags").Find(&notes)

	return notes
}

func (d *Db) GetNoteById(id string) (Note, bool) {
	var note Note

	trx := d.g.Model(&Note{}).Preload("Tags").Find(&note, "Id = ?", id)

	return note, trx.RowsAffected == 1
}

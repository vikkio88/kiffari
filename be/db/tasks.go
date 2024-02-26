package db

func (d *Db) GetTaskById(id string) (Task, bool) {
	var t Task

	trx := d.g.Model(&Task{}).Preload("Tags").Find(&t, "Id = ?", id)

	return t, trx.RowsAffected == 1
}

func (d *Db) CreateTask(tc TaskCreate, projectId string) (Task, bool) {
	t := tc.Task(projectId)
	trx := d.g.Create(&t)

	return t, trx.RowsAffected == 1
}

func (d *Db) UpdateTask(tu TaskUpdate, projectId string) (string, bool) {
	t := tu.Task(projectId)

	d.g.Model(&t).Association("Tags").Replace(t.Tags)
	trx := d.g.Omit("created_at").Save(&t)
	if trx.RowsAffected != 1 {
		return "", false
	}
	return t.Id, true
}

// func (d *Db) MoveTask(id string, status Status) bool {

// }

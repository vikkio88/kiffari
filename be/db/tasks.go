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

// func (d *Db) UpdateTask(p TaskUpdate) (string, bool) {
// 	pu := p.Project()
// 	trx := d.g.Omit("created_at").Save(&pu)
// 	if trx.RowsAffected != 1 {
// 		return "", false
// 	}
// 	return p.Id, true
// }

// func (d *Db) MoveTask(id string, status Status) bool {

// }

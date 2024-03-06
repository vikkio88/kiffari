package db

import "gorm.io/gorm"

func (d *Db) GetProjectById(id string) (ProjectDto, bool) {
	var p Project

	trx := d.g.Model(&Project{}).Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Where("tasks.archived != true").Order("tasks.status DESC, tasks.priority DESC, tasks.updated_at DESC")
	}).
		Preload("Tasks.Tags").Find(&p, "Id = ?", id)

	return p.Dto(), trx.RowsAffected == 1
}

func (d *Db) GetProjectWithArchivedTasksById(id string) (ProjectDto, bool) {
	var p Project

	trx := d.g.Model(&Project{}).Preload("Tasks", "archived == false").
		Preload("Tasks", "archived = true").
		Preload("Tasks.Tags").Find(&p, "Id = ?", id)

	return p.Dto(), trx.RowsAffected == 1
}

func (d *Db) CreateProject(p ProjectCreate) (ProjectDto, bool) {
	pc := p.Project()
	trx := d.g.Create(&pc)

	return pc.Dto(), trx.RowsAffected == 1
}

func (d *Db) UpdateProject(p ProjectUpdate) (ProjectDto, bool) {
	pu := p.Project()
	trx := d.g.Omit("created_at").Save(&pu)
	if trx.RowsAffected != 1 {
		return ProjectDto{}, false
	}
	return pu.Dto(), true
}

func (d *Db) GetAllProjects() []ProjectDto {
	var ps []Project

	d.g.Model(&Project{}).Order("updated_at DESC").Find(&ps)

	dtos := make([]ProjectDto, len(ps))
	for i, p := range ps {
		dtos[i] = p.Dto()
	}

	return dtos
}

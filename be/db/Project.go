package db

import (
	"kato-be/libs"
	"time"

	"github.com/oklog/ulid/v2"
)

/*
const project = {
    id: "somePrjectId",
    name: "Kiffari",
    description: "A Kanban-ish approach to managing side projects.",
    links: [
      { label: "Source", href: "https://github.com" },
      { label: "Website", href: "https://github.com" },
    ],
  };
*/

type Project struct {
	Id          string    `gorm:"primarykey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"-"`
	Links       string    `json:"-"`
	Config      string    `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`

	Tasks []Task `json:"-"`
}

type ProjectDto struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Links       []Link        `json:"links"`
	Config      ProjectConfig `json:"config"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`

	Tasks []Task `json:"tasks,omitempty"`
}

func (p *Project) Dto() ProjectDto {
	return ProjectDto{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Links:       libs.JsonDecode(p.Links, []Link{}),
		Config:      libs.JsonDecode(p.Config, ProjectConfig{}),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Tasks:       p.Tasks,
	}

}

type ProjectCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Links       []Link `json:"links"`

	Config ProjectConfig `json:"config"`
}

func (p *ProjectCreate) Project() Project {
	links := libs.JsonStringify(p.Links, "[]")
	config := libs.JsonStringify(p.Config, "{}")
	return Project{
		Id:          ulid.Make().String(),
		Name:        p.Name,
		Description: p.Description,
		Links:       links,
		Config:      config,
	}
}

type ProjectUpdate struct {
	ProjectCreate
	Id string `json:"id" binding:"required"`
}

func (p *ProjectUpdate) Project() Project {
	links := libs.JsonStringify(p.Links, "[]")
	config := libs.JsonStringify(p.Config, "{}")
	return Project{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Links:       links,
		Config:      config,
	}
}

type Link struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type ProjectConfig struct {
	WipLimit int `json:"wip_limit"`
}

package db

import (
	"encoding/json"
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
	Description string    `json:"description"`
	Links       string    `json:"links"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Tasks []Task
}

type ProjectDto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Links       []Link    `json:"links"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Tasks []Task `json:"tasks,omitempty"`
}

func (p *Project) Dto() ProjectDto {
	var links []Link
	json.Unmarshal([]byte(p.Links), &links)

	if links == nil {
		links = []Link{}
	}

	return ProjectDto{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Links:       links,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Tasks:       p.Tasks,
	}

}

type ProjectCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Links       []Link `json:"links"`
}

func (p *ProjectCreate) Project() Project {
	links, err := json.Marshal(p.Links)
	if err != nil {
		links = []byte("[]")
	}

	return Project{
		Id:          ulid.Make().String(),
		Name:        p.Name,
		Description: p.Description,
		Links:       string(links),
	}
}

type Link struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

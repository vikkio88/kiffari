package db

import (
	"kato-be/conf"

	"gorm.io/gorm"
)

func migrate(g *gorm.DB) {
	if conf.KiffariEnabled {
		g.AutoMigrate(&Tag{}, &Note{}, &Passkey{}, &Task{}, &Project{})
		return
	}
	g.AutoMigrate(&Tag{}, &Note{}, &Passkey{})
}

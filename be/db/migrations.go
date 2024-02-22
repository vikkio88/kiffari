package db

import "gorm.io/gorm"

func migrate(g *gorm.DB) {
	g.AutoMigrate(&Tag{}, &Note{}, &Passkey{})
}

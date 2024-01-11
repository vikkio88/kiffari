package main

import (
	"fmt"
	"kato-be/db"
)

func main() {
	d := db.NewDb("testing.db")
	// t := db.TagDto{}
	// t.Id = ulid.Make().String()
	// t.Value = "blabla"
	// d.G.Save(&t)
	var t []db.TagDto
	d.G.Model(&db.TagDto{}).Find(&t)

	fmt.Println(t)
}

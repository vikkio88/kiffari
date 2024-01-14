package main

import (
	"fmt"
	"kato-be/db"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	d := db.NewDb("testing.db")
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/tags", func(c *gin.Context) {
		var newTag db.TagCreate
		if c.Bind(&newTag) == nil {
			c.JSON(http.StatusOK, gin.H{"result": d.InsertTag(newTag)})
		}
	})

	r.GET("/tags", func(c *gin.Context) {
		filterValue := c.Query("q")
		if filterValue == "" {
			c.JSON(http.StatusOK, d.GetAllTags())
			return
		}

		c.JSON(http.StatusOK, d.FilterTags(filterValue))
	})

	r.GET("/notes", func(c *gin.Context) {
		c.JSON(http.StatusOK, d.GetAllNotes())
	})

	r.POST("/notes", func(c *gin.Context) {
		var newNote db.NoteCreate
		err := c.Bind(&newNote)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"result": d.InsertNote(newNote)})
			return
		}

		fmt.Println(err.Error())

	})
	r.Run(":5111") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

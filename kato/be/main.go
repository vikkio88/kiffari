package main

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig

}

func main() {
	d := db.NewDb("testing.db")
	r := gin.Default()

	r.Use(cors.New(corsConfig()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/tags", func(c *gin.Context) {
		var newTag db.TagCreate
		if c.Bind(&newTag) == nil {
			t, ok := d.InsertTag(newTag)
			if ok {
				c.JSON(http.StatusOK, t)
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": "could not create"})
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

	r.GET("/notes/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if note, ok := d.GetNoteById(id); ok {
			c.JSON(http.StatusOK, note)
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	r.POST("/notes", func(c *gin.Context) {
		var newNote db.NoteCreate
		err := c.Bind(&newNote)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"id": d.InsertNote(newNote)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})

	})

	r.Run(":5111")
}

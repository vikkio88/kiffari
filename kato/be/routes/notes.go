package routes

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(r gin.IRouter, d *db.Db) {
	r.GET("/notes", func(c *gin.Context) {
		latest := c.Query("latest")
		if latest == "true" {
			c.JSON(http.StatusOK, d.GetLatest())
			return
		}

		c.JSON(http.StatusOK, d.GetAllNotes())
	})

	r.GET("/notes/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		note, ok := d.GetNoteById(id)
		SuccessOr404(c, note, ok)
	})

	r.POST("/notes", func(c *gin.Context) {
		var newNote db.NoteCreate
		err := c.Bind(&newNote)
		if err == nil {
			n, ok := d.InsertNote(newNote)
			SuccessOr400(c, n, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	})

	r.PUT("/notes/:id", func(c *gin.Context) {
		var updateNote db.NoteUpdate
		err := c.Bind(&updateNote)
		if err == nil {
			updateNote.Id = c.Params.ByName("id")
			id, ok := d.UpdateNote(updateNote)
			SuccessOr400(c, gin.H{"result": id}, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	})

	r.DELETE("/notes/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.DeleteNote(id)
		if ok {
			c.JSON(http.StatusOK, gin.H{"result": true})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete"})
	})
}

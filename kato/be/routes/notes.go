package routes

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(r *gin.Engine, d *db.Db) {
	r.GET("/notes", func(c *gin.Context) {
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

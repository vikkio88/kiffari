package routes

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(r gin.IRouter, d *db.Db) {
	r.GET("/dash/notes", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"pinned":    d.GetPinnedNotes(),
				"reminders": d.GetReminderNotes(),
				"latest":    d.GetLatest(),
			},
		)
	})

	r.GET("/notes", func(c *gin.Context) {
		if isParamTrue(c, "archived") {
			c.JSON(http.StatusOK, d.GetArchivedNotes())
			return
		}

		filterValue := c.Query("q")
		if filterValue == "" {
			c.JSON(http.StatusOK, d.GetAllNotes())
			return
		}

		//TODO: maybe add an opts to parse in here, so you can pass values like ASC/DESC Archived and Dates?
		titleOnly := c.Query("t")
		c.JSON(http.StatusOK, d.FilterNotes(filterValue, titleOnly == "true"))
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
			n, ok := d.CreateNote(newNote)
			SuccessOr400(c, n, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})

	r.POST("/notes/:id/archive", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetArchivedNote(id, true)
		SuccessOr400(c, gin.H{"result": id}, ok)
	})

	r.DELETE("/notes/:id/archive", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetArchivedNote(id, false)
		SuccessOr400(c, gin.H{"result": id}, ok)
	})

	r.POST("/notes/:id/pin", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetPinnedNote(id, true)
		SuccessOr400(c, gin.H{"result": id}, ok)
	})

	r.DELETE("/notes/:id/pin", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetPinnedNote(id, false)
		SuccessOr400(c, gin.H{"result": id}, ok)
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

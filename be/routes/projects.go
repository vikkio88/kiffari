package routes

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(r gin.IRouter, d *db.Db) {
	r.GET("/projects", func(c *gin.Context) {
		c.JSON(http.StatusOK, d.GetAllProjects())
	})

	r.POST("/projects", func(c *gin.Context) {
		var newP db.ProjectCreate
		err := c.Bind(&newP)
		if err == nil {
			n, ok := d.CreateProject(newP)
			SuccessOr400(c, n, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})

	r.PUT("/projects/:id", func(c *gin.Context) {
		var updateProject db.ProjectUpdate
		err := c.Bind(&updateProject)
		if err == nil {
			updateProject.Id = c.Params.ByName("id")
			id, ok := d.UpdateProject(updateProject)
			SuccessOr400(c, gin.H{"result": id}, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})
}

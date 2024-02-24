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
}

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

	r.GET("/projects/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		p, ok := d.GetProjectById(id)
		SuccessOr404(c, p, ok)
	})

	r.POST("/projects/:id/task", func(c *gin.Context) {
		projectId := c.Params.ByName("id")
		var nTask db.TaskCreate
		err := c.Bind(&nTask)
		if err == nil {
			n, ok := d.CreateTask(nTask, projectId)
			SuccessOr400(c, n, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})

	r.PUT("/projects/:id/task", func(c *gin.Context) {
		projectId := c.Params.ByName("id")
		var uTask db.TaskUpdate
		err := c.Bind(&uTask)
		if err == nil {
			id, ok := d.UpdateTask(uTask, projectId)
			SuccessOr400(c, gin.H{"result": id}, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

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

	r.PUT("/projects/:id/task/:taskId/status", func(c *gin.Context) {
		projectId := c.Params.ByName("id")
		taskId := c.Params.ByName("taskId")
		var statusW db.StatusWrapper
		err := c.Bind(&statusW)
		if err == nil {
			ok, err := d.MoveTask(projectId, taskId, statusW.Status)
			SuccessOr400WithError(c, gin.H{"result": ok}, err)
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
			p, ok := d.UpdateProject(updateProject)
			SuccessOr400(c, p, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		task, ok := d.GetTaskById(id)
		SuccessOr404(c, task, ok)
	})

	r.POST("/tasks/:id/archive", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetArchivedTask(id, true)
		SuccessOr400(c, gin.H{"result": id}, ok)
	})

	r.DELETE("/tasks/:id/archive", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.SetArchivedTask(id, false)
		SuccessOr400(c, gin.H{"result": id}, ok)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.DeleteTask(id)
		if ok {
			c.JSON(http.StatusOK, gin.H{"result": true})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete"})
	})
}

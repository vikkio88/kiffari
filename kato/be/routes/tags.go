package routes

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TagRoutes(r *gin.Engine, d *db.Db) {
	r.POST("/tags", func(c *gin.Context) {
		var newTag db.TagCreate
		if c.Bind(&newTag) == nil {
			t, ok := d.InsertTag(newTag)
			SuccessOr400(c, t, ok)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create"})
	})

	r.GET("/tags", func(c *gin.Context) {
		filterValue := c.Query("q")
		if filterValue == "" {
			c.JSON(http.StatusOK, d.GetAllTags())
			return
		}

		c.JSON(http.StatusOK, d.FilterTags(filterValue))
	})

	r.DELETE("/tags/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		ok := d.DeleteTag(id)
		if ok {
			c.JSON(http.StatusOK, gin.H{"result": true})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete"})
	})
}
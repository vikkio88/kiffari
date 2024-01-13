package main

import (
	"kato-be/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	d := db.NewDb("testing.db")
	r := gin.Default()

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
			c.JSON(http.StatusOK, d.GetAll())
			return
		}

		c.JSON(http.StatusOK, d.Filter(filterValue))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

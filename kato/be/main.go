package main

import (
	"kato-be/db"
	"kato-be/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
	}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders(
		"Access-Control-Allow-Headers",
		"access-control-allow-origin, access-control-allow-headers",
		"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With",
		"Authorization",
	)
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

	r.POST("/login", func(c *gin.Context) {
		var pk db.PasskeyClear
		if c.Bind(&pk) == nil {
			t, err := d.CheckPk(pk)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": t})

			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not login"})

	})

	private := r.Group("")
	private.Use(AuthRequired)

	routes.TagRoutes(private, d)
	routes.NoteRoutes(private, d)

	r.Run(":5111")
}

func AuthRequired(c *gin.Context) {
	if c.GetHeader("Authorization") != "user" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauth"})
		return
	}

	c.Next()
}

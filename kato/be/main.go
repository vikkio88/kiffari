package main

import (
	"kato-be/conf"
	"kato-be/db"
	"kato-be/routes"
	"kato-be/validators"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	d := db.NewDb("testing.db")
	r := gin.Default()
	gin.SetMode(conf.GinMode)
	//TODO: check the Proxy warning
	/*
		https://github.com/gin-gonic/gin/issues/2809
		https://github.com/gin-gonic/gin/pull/3449/files#diff-b335630551682c19a781afebcf4d07bf978fb1f8ac04c6bf87428ed5106870f5L2251
	*/

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("dateInTheFuture", validators.DateInTheFuture)
	}

	r.Use(cors.New(corsConfig()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": conf.Version,
		})
	})

	routes.AuthRoutes(r, d)

	private := r.Group("")
	private.Use(AuthRequired(d))

	routes.TagRoutes(private, d)
	routes.NoteRoutes(private, d)

	r.Run(":5111")
}

// TODO: move to middleware module
func AuthRequired(db *db.Db) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := ""
		if token = c.GetHeader("Authorization"); token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauth"})
			return
		}

		token = strings.TrimSpace(token)
		if !db.IsTokenValid(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}

		c.Next()
	}
}

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

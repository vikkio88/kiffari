package main

import (
	"fmt"
	"kato-be/conf"
	"kato-be/db"
	"kato-be/routes"
	"kato-be/validators"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	d := db.NewDb()
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

	r.Use(static.ServeRoot("/", "./static"))

	r.GET("api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": conf.Version,
		})
	})

	routes.AuthRoutes(r, d)

	private := r.Group("api")
	private.Use(AuthRequired(d))

	routes.TagRoutes(private, d)
	routes.NoteRoutes(private, d)

	r.Run(fmt.Sprintf(":%s", conf.Port))
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
	corsConfig.AllowOrigins = conf.Cors
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

package main

import (
	"fmt"
	"kato-be/conf"
	"kato-be/db"
	"kato-be/middlewares"
	"kato-be/routes"
	"kato-be/validators"
	"net/http"

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

	r.GET("api/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": conf.Version,
			"kiffari": conf.KiffariEnabled,
		})
	})
	routes.AuthRoutes(r, d)

	private := r.Group("api")
	private.Use(middlewares.AuthRequired(d))

	routes.TagRoutes(private, d)
	routes.NoteRoutes(private, d)

	if conf.KiffariEnabled {
		routes.ProjectRoutes(private, d)
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("./static")
	})

	r.Run(fmt.Sprintf(":%s", conf.Port))
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

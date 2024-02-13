package main

import (
	"kato-be/conf"
	"kato-be/db"
	"kato-be/routes"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

// TODO: move to validator module
var dateInTheFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	if date.IsZero() {
		return true
	}

	today := time.Now()
	return !today.After(date)
}

func main() {
	d := db.NewDb("testing.db")
	r := gin.Default()
	gin.SetMode(conf.GinMode)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("dateInTheFuture", dateInTheFuture)
	}

	r.Use(cors.New(corsConfig()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": conf.Version,
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

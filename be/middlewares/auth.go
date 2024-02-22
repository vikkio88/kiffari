package middlewares

import (
	"kato-be/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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

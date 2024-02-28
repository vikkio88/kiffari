package middlewares

import (
	"kato-be/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthToken(c *gin.Context) string {
	return strings.TrimSpace(c.GetHeader("Authorization"))
}

func AuthRequired(db *db.Db) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := ""
		if token = GetAuthToken(c); token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauth"})
			return
		}

		if !db.IsTokenValid(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}

		c.Next()
	}
}

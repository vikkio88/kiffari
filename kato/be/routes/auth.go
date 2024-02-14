package routes

import (
	"kato-be/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r gin.IRouter, d *db.Db) {
	r.POST("/login", func(c *gin.Context) {
		var pk db.PasskeyClear
		if c.Bind(&pk) == nil {

			t, err := d.CheckPk(pk)

			// if the password is not initialised the fist login sets it
			if err != nil && strings.Contains(err.Error(), "Initialised") {
				if d.StorePassword(pk.Hash()) {
					t, _ := d.CheckPk(pk)
					c.JSON(http.StatusOK, gin.H{"token": t})
					return
				}
			}

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": t})

			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not login"})

	})
}

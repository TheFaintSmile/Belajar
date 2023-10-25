package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}

func PendidikAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := ExtractTokenRole(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		if role == "Siswa" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Forbidden Endpoint",
			})
			return
		}
		c.Next()
	}
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRole chỉ cho phép người có role phù hợp
func RequireRole(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != required {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

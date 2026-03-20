package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authentication(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
		return
	}

	token := strings.Split(header, " ")[1]
	userId, role, err := m.jwt.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token/failed to validate token",
		})
	}

	c.Set("userId", userId)
	c.Set("role", role)
	c.Next()
}

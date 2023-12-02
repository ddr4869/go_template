package middleware

import (
	"github.com/gin-gonic/gin"
)

func MwAbortJson(c *gin.Context, statusCode int, err error) {

	c.JSON(statusCode, gin.H{"error": err.Error()})
	c.Abort()
}

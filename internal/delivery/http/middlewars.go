package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func recoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Connection", "close")
				c.IndentedJSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Errorf("%s", err).Error(),
				})
			}
		}()
		c.Next()
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Correlation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var xReqId = c.GetHeader("X-Request-Id")
		if xReqId == "" {
			xReqId = uuid.New().String()
		}

		c.Request.Header.Set("X-Request-Id", xReqId)
		c.Next()
	}
}
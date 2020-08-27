package midware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"use-gin/logger"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")

		if requestID == "" {
			u4, err := uuid.NewV4()
			if err != nil {
				logger.Log.Errorf("create uuid failed: %v", err)
			}
			requestID = u4.String()
			logger.Log.Warnf(requestID)
		}

		// NOTE: handler function can be get X-Request-Id by c.Get("X-Request-Id")
		c.Set("X-Request-Id", requestID)

		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

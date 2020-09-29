package midware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"ginner/logger"
)

// RequestID trace request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")

		if requestID == "" {
			u4, err := uuid.NewV4()
			if err != nil {
				logger.Log.Errorf("create uuid failed: %v", err)
			}
			requestID = u4.String()
		}

		// NOTE: handler function can be get X-Request-Id by c.Get("requestID")
		c.Set("requestID", requestID)

		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}

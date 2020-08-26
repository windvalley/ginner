package midware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/util"
)

// A client(ip) will be dennied when it's access rate is over busrtSize/second.
func TrafficLimit(burstSize int) gin.HandlerFunc {
	var limiter = util.NewIPRateLimiter(1, burstSize)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		limiter := limiter.GetLimiter(clientIP)
		if limiter.Allow() {
			c.Next()
			return
		}

		err := errors.New(fmt.Sprintf(
			"%s request over %d, dennied", clientIP, burstSize))
		err1 := errcode.New(errcode.TooManyRequestError, err)
		handler.SendResponse(c, err1, nil)

		c.Abort()
		return
	}
}

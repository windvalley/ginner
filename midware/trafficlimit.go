package midware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/util"
)

// UserTrafficLimiter A client(ip) will be dennied when it's access rate is over busrtSize/second.
func UserTrafficLimiter(burstSize int) gin.HandlerFunc {
	var limiter = util.NewIPRateLimiter(1, burstSize)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		limiter := limiter.GetLimiter(clientIP)
		if limiter.Allow() {
			c.Next()
			return
		}

		err := fmt.Errorf("%s requests over %d/second, dennied", clientIP, burstSize)
		err1 := errcode.New(errcode.TooManyRequestError, err)
		handler.SendResponse(c, err1, nil)

		c.Abort()
		return
	}
}

// GlobalTrafficLimiter request will be dennied when the total requests in one seconds over burstSize.
func GlobalTrafficLimiter(burstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(1, burstSize)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}

		err := fmt.Errorf("global requests rate over %d/second, dennied", burstSize)
		err1 := errcode.New(errcode.TooManyRequestError, err)
		handler.SendResponse(c, err1, nil)

		c.Abort()
		return
	}
}

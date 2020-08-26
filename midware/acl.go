package midware

import (
	"github.com/gin-gonic/gin"

	"use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
)

// ACL access limited
func ACL() gin.HandlerFunc {
	allowURLMap := make(map[string]bool)
	allowIPMap := make(map[string]bool)

	for _, v := range config.Conf().ACL.AllowURL {
		allowURLMap[v] = true
	}

	for _, v := range config.Conf().ACL.AllowIP {
		allowIPMap[v] = true
	}

	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		ip := c.ClientIP()

		if !allowURLMap[requestPath] && len(allowURLMap) != 0 {
			err := errcode.New(errcode.AccessForbiddenError, nil)
			err.Add(requestPath + " is not allowed")
			handler.SendResponse(c, err, nil)

			c.Abort()
			return
		}

		if !allowIPMap[ip] && len(allowIPMap) != 0 {
			err := errcode.New(errcode.AccessForbiddenError, nil)
			err.Add(ip + " is not allowed")
			handler.SendResponse(c, err, nil)

			c.Abort()
			return
		}

		c.Next()
		return
	}
}

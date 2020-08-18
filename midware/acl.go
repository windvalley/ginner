package midware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"use-gin/config"
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

	return func(ctx *gin.Context) {
		requestPath := ctx.Request.URL.Path
		ip := strings.Split(ctx.Request.RemoteAddr, ":")[0]
		if ip == "[" {
			ipTmp := strings.Split(ctx.Request.RemoteAddr, "[")[1]
			ip = strings.Split(ipTmp, "]")[0]
		}

		if !allowURLMap[requestPath] && len(allowURLMap) != 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, &handler.Response{
				Code:    "AccessForbidden",
				Message: requestPath + " is not allowed",
				Data:    nil,
			})
		}

		if !allowIPMap[ip] && len(allowIPMap) != 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, &handler.Response{
				Code:    "AccessForbidden",
				Message: ip + " is not allowed",
				Data:    nil,
			})
		}

		ctx.Next()
		return
	}
}

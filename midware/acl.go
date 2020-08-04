package midware

import (
	"github.com/gin-gonic/gin"

	"use-gin/config"
)

// ACL access limited
func ACL() gin.HandlerFunc {
	permitURLMap := make(map[string]bool)
	permitIPMap := make(map[string]bool)

	for _, v := range config.Config().ACL.AllowURL {
		permitURLMap[v] = true
	}

	for _, v := range config.Config().ACL.AllowIP {
		permitIPMap[v] = true
	}

	return func(ctx *gin.Context) {
		if permitURLMap[ctx.Request.URL.Path] || len(permitURLMap) == 0 {
			ip := ctx.Request.Header.Get("X-Real-Ip")
			if permitIPMap[ip] || len(permitIPMap) == 0 {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatus(403)
	}
}

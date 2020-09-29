package midware

import (
	"github.com/gin-gonic/gin"

	"ginner/config"
)

// UserAudit enable user audit log
func UserAudit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(config.UserAuditEnableKey, true)

		c.Next()
	}
}

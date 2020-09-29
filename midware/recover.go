package midware

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
)

// Recover recover from panic and send fitting response to client
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errStr := fmt.Sprintf("%s", err)
				debugStack := errStr
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					debugStack += v + " "
				}

				err1 := errcode.New(errcode.ServerPanicError, errors.New(debugStack))
				api.SendResponse(c, err1, nil)
			}
		}()

		c.Next()
	}
}

package midware

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/template"
	"use-gin/util"
)

// Exception catch panic Globally
func Exception() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					debugStack += v + "<br>"
				}

				subject := fmt.Sprintf("[Alert] %s error", config.Conf().AppName)

				body := strings.ReplaceAll(
					template.AlertMail,
					"{ErrorMsg}",
					fmt.Sprintf("%s", err),
				)
				body = strings.ReplaceAll(
					body,
					"{RequestTime}",
					time.Now().Format("2006-01-02 15:04:05"),
				)
				body = strings.ReplaceAll(
					body,
					"{RequestURL}",
					c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI,
				)
				body = strings.ReplaceAll(
					body,
					"{RequestUA}",
					c.Request.UserAgent(),
				)
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{Debug Stack}", debugStack)

				if err := util.SendMail(
					config.Conf().Mail.MailTo,
					subject,
					body,
				); err != nil {
					err1 := errcode.New(errcode.ServerPanicError, err)
					handler.SendResponse(c, err1, nil)
				}
			}
		}()

		c.Next()
	}
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"use-gin/errcode"
	"use-gin/logger"
	"use-gin/util"
)

// Response response with json format
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// SendResponse response JSON
func SendResponse(c *gin.Context, err error, data interface{}) {
	status, code, message, sysErrMsg := errcode.DecodeErr(err)

	c.JSON(status, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})

	// write to error.log
	if status != http.StatusOK && status != http.StatusCreated {
		requestID := util.GetRequestID(c)
		username := util.GetUsername(c)

		logger.Log.WithFields(logrus.Fields{
			"username":        username,
			"client_ip":       c.ClientIP(),
			"request_method":  c.Request.Method,
			"request_uri":     c.Request.URL.Path,
			"http_status":     c.Writer.Status(),
			"latency_time":    nil,
			"request_proto":   c.Request.Proto,
			"request_referer": c.Request.Referer(),
			"request_body":    c.Request.PostForm.Encode(),
			"request_id":      requestID,
			"request_ua":      c.Request.UserAgent(),
			"response_code":   code,
			"response_msg":    message,
		}).Error(sysErrMsg)
	}
}

// SendFile response with file
func SendFile(c *gin.Context, filepath string, filename string) {
	// filepath is the fullpath in server,
	// filename is the file name of the user save to.
	c.FileAttachment(filepath, filename)
}

// SendString response with txt
func SendString(c *gin.Context, text string) {
	c.String(http.StatusOK, text)
}

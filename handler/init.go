package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"use-gin/errcode"
	"use-gin/logger"
)

// response json format
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// response JSON
func SendResponse(c *gin.Context, err error, data interface{}) {
	status, code, message := errcode.DecodeErr(err)

	c.JSON(status, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})

	if status != http.StatusOK {
		logger.Log.Error(err)
	}
}

// response file
func SendFile(c *gin.Context, filepath string, filename string) {
	// filepath is the fullpath in server,
	// filename is the file name of the user save to.
	c.FileAttachment(filepath, filename)
}

// response txt
func SendString(c *gin.Context, text string) {
	c.String(http.StatusOK, text)
}

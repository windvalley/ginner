package api

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v1"
)

// HandleInfluxdbDemo a demo of handle influxdb
func HandleInfluxdbDemo(c *gin.Context) {
	err := service.HandleInfluxdbDemo()
	if err != nil {
		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

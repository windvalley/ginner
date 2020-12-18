package api

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v1"
)

// HandleKafkaDemo a demo of handle kafka
func HandleKafkaDemo(c *gin.Context) {
	err := service.HandleKafkaDemo(c)
	if err != nil {
		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

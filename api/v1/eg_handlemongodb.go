package api

import (
	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v1"
)

// HandleMongodbDemo a demo of handle mongodb
func HandleMongodbDemo(c *gin.Context) {
	err := service.HandleMongodbDemo()
	if err != nil {
		if err == errcode.ErrRecordNotFound {
			api.SendResponse(c, errcode.RecordNotFoundError, nil)
			return
		}

		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}

package apiv1

import (
	"github.com/gin-gonic/gin"

	"use-gin/api"
)

// SignatureDemo for testing API Signature Authentication
func SignatureDemo(c *gin.Context) {
	api.SendResponse(c, nil, "hello world!")
}

package handler

import "github.com/gin-gonic/gin"

func Status(c *gin.Context) {
	SendString(c, "ok")
}

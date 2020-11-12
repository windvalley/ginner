package util

import (
	"github.com/gin-gonic/gin"

	"ginner/logger"
)

type pagination struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page-size" binding:"required"`
}

// Paginate If page and page-size validate error, it will return defaultPageSize records.
func Paginate(c *gin.Context, defaultPageSize int) (offset, limit, page, pageSize int) {
	var r pagination
	if err := c.ShouldBind(&r); err != nil {
		logger.Log.Debugf(
			"use the default values of page(1) and page-size(%d), bind error: %v",
			defaultPageSize, err)
		return 0, defaultPageSize, 1, defaultPageSize
	}

	if r.Page > 0 && r.PageSize > 0 {
		offset = (r.Page - 1) * r.PageSize
		limit = r.PageSize
	} else {
		return 0, defaultPageSize, 1, defaultPageSize
	}

	// offset and limit used for quering mysql,
	// and r.page and r.pageSize used for responding to users.
	return offset, limit, r.Page, r.PageSize
}

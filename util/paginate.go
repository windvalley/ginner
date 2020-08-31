package util

import (
	"fmt"
	"strconv"

	"use-gin/logger"

	"github.com/gin-gonic/gin"
)

// Paginate The request url must has page and page_size parameter,
// and if not have this two parameters, the server will return all records.
func Paginate(c *gin.Context) (int, int, int, int) {
	page, pageSize := 0, 0
	if pageStr := c.Query("page"); pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			logger.Log.Warnf("page(string) convert to int error: %v", err)
		}
	}
	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		var err error
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			logger.Log.Warnf("page_size(string) convert to int error: %v", err)
		}
	}

	// return all records
	offset, limit := -1, -1

	// paginate
	if page > 0 && pageSize > 0 {
		offset = (page - 1) * pageSize
		limit = pageSize
	} else {
		fmt.Println(page, pageSize)
		page, pageSize = -1, -1
	}

	// offset and limit used for mysql query;
	// page and pageSize used for return to users.
	return offset, limit, page, pageSize
}

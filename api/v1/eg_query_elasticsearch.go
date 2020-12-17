package api

import (
	"strings"

	"github.com/gin-gonic/gin"

	"ginner/api"
	"ginner/errcode"
	"ginner/service/v1"
	"ginner/util"
)

// FilterRecordsReq request parameters of the client
type FilterRecordsReq struct {
	service.FilterParams
	LeafIDs string `form:"leaf_ids" binding:"required"`
}

type filterRecordsResp struct {
	Page     int                      `json:"page"`
	Pagesize int                      `json:"page_size"`
	Count    int                      `json:"count"`
	Domains  []map[string]interface{} `json:"domains"`
}

// FilterRecordsFromES filter or search records from es
func FilterRecordsFromES(c *gin.Context) {
	var r FilterRecordsReq
	if err := c.ShouldBind(&r); err != nil {
		err1 := errcode.New(errcode.ValidationError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	departIDs := strings.Split(r.LeafIDs, ",")

	offset, limit, page, pageSize := util.Paginate(c, 20)
	resp, count, err := service.FilterRecordsFromES(
		departIDs, r.FilterParams, offset, limit)
	if err != nil {
		err1 := errcode.New(errcode.CustomInternalServerError, err)
		err1.Add(err)
		api.SendResponse(c, err1, nil)
		return
	}

	api.SendResponse(c, nil, filterRecordsResp{page, pageSize, count, resp})
}

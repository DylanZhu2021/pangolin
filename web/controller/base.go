package controller

import (
	"github.com/gin-gonic/gin"
	"pangolin/core/model"
)

func Welcome(c *gin.Context) {
	ResponseWithSuccessMsg(c, "welcome to pangolin")
}

// Query 查询
func Query(c *gin.Context) {
	var request = &model.SearchRequest{}
	if err := c.ShouldBind(&request); err != nil {
		ResponseWithErrorMsg(c, err.Error())
		return
	}
	//调用搜索
	r := srv.Base.Query(request)
	ResponseSuccessWithData(c, r)
}

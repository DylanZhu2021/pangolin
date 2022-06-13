package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回信息的文件

// ResponseData 返回数据结构体
type ResponseData struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

// ResponseSuccessWithData 返回成功，带数据
func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Message: "success",
		Status:  true,
		Data:    data,
	})
}

// ResponseWithSuccessMsg 返回成功，不带数据
func ResponseWithSuccessMsg(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &ResponseData{
		Message: message,
		Status:  true,
		Data:    nil,
	})
}

// ResponseWithErrorMsg 返回失败
func ResponseWithErrorMsg(c *gin.Context, message string) {
	c.JSON(http.StatusOK, &ResponseData{
		Message: message,
		Status:  false,
		Data:    nil,
	})
}

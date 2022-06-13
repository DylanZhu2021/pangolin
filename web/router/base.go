package router

import (
	"github.com/gin-gonic/gin"
	"pangolin/web/controller"
)

// InitRouter 基础管理路由
func InitRouter(router *gin.RouterGroup) {

	BaseRouter := router.Group("")
	{
		BaseRouter.GET("/", controller.Welcome)
		BaseRouter.POST("query", controller.Query)
	}
}

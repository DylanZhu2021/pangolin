package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

// SetupRouter 路由管理
func SetupRouter() *gin.Engine {

	router := gin.Default()

	var handlers []gin.HandlerFunc

	group := router.Group("/api", handlers...)
	{
		InitRouter(group) // 基础管理
	}
	log.Printf("API Url: \t http://%s/api", "127.0.0.0:8080")
	return router
}

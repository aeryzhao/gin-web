package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iszhaoxg/gin-web/pkg/setting"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return engine
}

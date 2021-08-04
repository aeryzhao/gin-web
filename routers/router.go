package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/iszhaoxg/gin-web/api/v1"
	"github.com/iszhaoxg/gin-web/pkg/setting"
	"github.com/iszhaoxg/gin-web/routers/api"
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

	engine.GET("/auth", api.GetAuth)

	apiV1 := engine.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTags)
		apiV1.PUT("/tags/:id", v1.EditTags)
		apiV1.DELETE("/tags/:id", v1.DeleteTags)
	}

	return engine
}

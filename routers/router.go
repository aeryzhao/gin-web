package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/iszhaoxg/gin-web/api/v1"
	"github.com/iszhaoxg/gin-web/middleware/jwt"
	"github.com/iszhaoxg/gin-web/pkg/setting"
	"github.com/iszhaoxg/gin-web/routers/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTags)
		apiV1.PUT("/tags/:id", v1.EditTags)
		apiV1.DELETE("/tags/:id", v1.DeleteTags)
	}

	articleApi := r.Group("/api/v1")
	{
		articleApi.GET("/articles", v1.GetArticles)
		articleApi.GET("/articles/:id", v1.GetArticle)
		articleApi.POST("/articles", v1.AddArticle)
		articleApi.PUT("/articles/:id", v1.EditArticle)
		articleApi.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}

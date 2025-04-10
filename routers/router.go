package routers

import (
	"github.com/aeryzhao/gin-web/api"
	v1 "github.com/aeryzhao/gin-web/api/v1"
	"github.com/aeryzhao/gin-web/middleware/jwt"
	"github.com/aeryzhao/gin-web/pkg/setting"
	"github.com/gin-gonic/gin"
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

	tagsV1 := r.Group("/api/v1/tags")
	tagsV1.Use(jwt.JWT())
	{
		tagsV1.GET("", v1.GetTags)
		tagsV1.POST("", v1.AddTags)
		tagsV1.PUT("/:id", v1.EditTags)
		tagsV1.DELETE("/:id", v1.DeleteTags)
	}

	articlesV1 := r.Group("/api/v1/articles")
	{
		articlesV1.GET("", v1.GetArticles)
		articlesV1.GET("/:id", v1.GetArticle)
		articlesV1.POST("", v1.AddArticle)
		articlesV1.PUT("/:id", v1.EditArticle)
		articlesV1.DELETE("/:id", v1.DeleteArticle)
	}

	return r
}

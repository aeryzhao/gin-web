package main

import (
	"fmt"
	"github.com/aeryzhao/gin-web/pkg/setting"
	"github.com/aeryzhao/gin-web/routers"
	"net/http"
)

// @title                      简单博客服务
// @version                    1.0
// @description                简单的 web 服务
// @host                       localhost:8000
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                JWT Token (格式: `Bearer <token>`)
func main() {
	router := routers.InitRouter()

	server := http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

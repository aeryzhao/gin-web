module github.com/aeryzhao/gin-web

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect; indirectggo
	github.com/jinzhu/gorm v1.9.16
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.3
	github.com/swaggo/swag v1.8.5
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220909162455-aba9fc2a8ff2 // indirect
	golang.org/x/tools v0.1.12 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	github.com/aeryzhao/gin-web/conf => ./gin-web/pkg/conf
	github.com/aeryzhao/gin-web/middleware => ./gin-web/middleware
	github.com/aeryzhao/gin-web/models => ./gin-web/models
	github.com/aeryzhao/gin-web/pkg/error => ./gin-web/pkg/error
	github.com/aeryzhao/gin-web/pkg/setting => ./gin-web/pkg/setting
	github.com/aeryzhao/gin-web/routers => ./gin-web/routers
)

# gin-web
本项目使用到的依赖
- gorm
- gin
- swag-go
- go-ini

## gin
[官方文档](https://gin-gonic.com/zh-cn/docs/)

## gorm
[官方文档](https://gorm.io/zh_CN/docs/index.html)

## swag-go
[官方文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

### 安装
要使用swag-go,首先需要安装`swag cli`

```shell
$ go get -u github.com/swaggo/swag/cmd/swag

# 1.16 及以上版本
$ go install github.com/swaggo/swag/cmd/swag@latest
```

### 使用
还需要再项目中引用的包

```shell
# gin-swagger 中间件
go get github.com/swaggo/gin-swagger
# swagger 内置文件
go get github.com/swaggo/gin-swagger/swaggerFiles
```

### 注释

#### API操作
| 注释        | 描述                                                         |
| ----------- | ------------------------------------------------------------ |
| description | 操作行为的详细说明。                                         |
| id          | 用于标识操作的唯一字符串。在所有API操作中必须唯一。          |
| tags        | 每个API操作的标签列表，以逗号分隔。                          |
| summary     | 该操作的简短摘要。                                           |
| produce     | API可以生成的MIME类型的列表。值必须如“[Mime类型](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime-types)”中所述。 |
| param       | 用空格分隔的参数。`param name`,`param type`,`data type`,`is mandatory?`,`comment` `attribute(optional)` |
| security    | 每个API操作的[安全性](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#安全性)。 |
| success     | 以空格分隔的成功响应。`return code`,`{param type}`,`data type`,`comment` |
| failure     | 以空格分隔的故障响应。`return code`,`{param type}`,`data type`,`comment` |
| response    | 与success、failure作用相同                                   |
| header      | 以空格分隔的头字段。 `return code`,`{param type}`,`data type`,`comment` |
| router      | 以空格分隔的路径定义。 `path`,`[httpMethod]`                 |

### 生成

进入项目根目录中，执行初始化命令
```shell
swag init
```

会在项目目录中生成如下：
```shell
docs/
├── docs.go
└── swagger.json
└── swagger.yaml
```
启动项目访问`ip:port/swagger/index.html`

## Docker 镜像

1.编写 Dockerfile 文件

```dockerfile
FROM scratch

WORKDIR $GOPATH/src/github.com/iszhaoxg/gin-web
COPY . $GOPATH/src/github.com/iszhaoxg/gin-web

EXPOSE 8000
CMD ["./go-gin-example"]
```

2.编译可执行文件
```shell
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .
```

3.构建镜像
```shell
docker build -t gin-blog-docker-scratch .
```

4.运行
```shell
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker-scratch
```

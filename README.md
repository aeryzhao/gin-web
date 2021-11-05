# gin-web
使用到的依赖
- gorm
- gin
- swaggo
- go-ini
## gin


## gorm

## swaggo
### 安装
要使用swaggo,首先需要安装`swag cli`
> go get -u github.com/swaggo/swag/cmd/swag

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
> swag init

会在项目目录中生成如下：
```shell
docs/
├── docs.go
└── swagger.json
└── swagger.yaml
```
启动项目访问`ip:port/swagger/index.html`
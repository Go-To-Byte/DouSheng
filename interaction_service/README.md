# DouSheng
第五届青训营项目——抖声

## 启动说明

### 1、项目配置文件
需要将`etc/dousheng.toml.template`配置模板文件复制粘贴为：`etc/dousheng.toml`，并且修改为自己的配置，如：
```toml
[app]
name = "douyin"

[app.http]
host = "127.0.0.1"
port = "8080"

[app.grpc]
host = "127.0.0.1"
port = "8505"

[mysql]
host = "127.0.0.1"
port = "3306"
username = "root"
password = "root"
database = "dou_sheng"

[log]
level = "debug"
dir = "logs"
format = "text"
to = "stdout"
```

### 2、项目启动
本项目采用了Makefile工程化便于开发，电脑需要额外安装`make`指令。
* windows推荐安装教程：[推荐第二种方式，别忘记配置环境变量](https://tehub.com/a/aCYp1uw0tG)
* 命令行键入`make run`启动项目
* 当然，不想安装make指令，也可以直接键入`go run main.go start`启动项目


## 项目结构

* 目录结构概览[解读]
```text
DouSheng            # 极简版抖音 APP
├── apps            # 所有服务模块[其中的每一个模块，都可单独拆分出来做成微服务]
│   ├── all         # 统一管理所有模块实例的注册[驱动加载的方式]
│   ├── comment     # 评论模块
│   │   ├── api
│   │   ├── impl
│   │   └── pb
│   ├── user        # 用户模块
│   │   ├── api     
│   │   ├── impl
│   │   └── pb
│   └── video       # 视频模块
│       ├── api
│       ├── impl
│       └── pb
├── cmd             # CLI
├── common.pb       # 放置公共的protobuf文件[可抽离]
├── conf            # 项目配置对象
├── docs            # 项目相关文档
├── etc             # 项目具体配置
├── ioc             # IoC容器[可抽离]
├── protocol        # 提供协议
├── utils           # 工具包
└── version         # 版本信息
```

* 部分主要文件概览[解读]

```text
├── apps                            # 所有的业务模块
│   ├── all                         # 驱动注册所有的IOC容器实例
│   │   └── auto_register.go
│   ├── user                        # 以用户模块举例
│   │   ├── api                     # 提供的 API 接口
│   │   │   ├── http.go             # 使用 HTTP 的方式暴露 控制层逻辑
│   │   │   └── user.go             # user服务模块暴露的方法
│   │   ├── app.go                  # user模块的结构体方法
│   │   ├── impl                    # user.ServerService 的实现
│   │   │   ├── dao.go              # 可以看作是 持久层逻辑
│   │   │   ├── impl.go             # 可以看作是 业务层逻辑
│   │   │   ├── user.go             # user.ServerService 接口方法的实现
│   │   │   └── user_test.go        # 此模块测试用例【注：必写，一般用于测试本模块CURD的功能】
│   │   ├── pb                      # 此模块的protobuf文件，里面有（接口方法、请求model、响应model、本模块model）
│   │   │   └── user.proto      
│   │   ├── README.md               # 本模块说明
│   │   ├── user.pb.go              # 利用 protoc 生成（结构体）
│   │   └── user_grpc.pb.go         # 利用 protoc 生成（接口）
├── cmd                             # 用于启动项目
│   ├── root.go                     
│   └── start.go                    # 启动逻辑在这
├── common                          # 定义的公共的protobuf文件，可抽离
│   ├── common.pb.go
│   └── pb
│       └── common.proto
├── conf                            # 项目配置对象
│   ├── app.go                      # 此项目的配置
│   ├── config.go                   # 统一配置
│   ├── config_test.go              
│   ├── load.go                     # 加载所有配置
│   ├── log.go                      # 日志相关配置
│   └── mysql.go                    # mysql相关配置
├── etc
│   ├── dousheng.toml               # 项目配置文件位置【可换成其他的，用其他库解析】[禁止上传github]
│   └── dousheng.toml.template      # 配置文件模板[可上传github]
├── ioc                             # IoC容器
│   ├── all.go                      # 统一所有容器
│   ├── gin.go                      # Gin HTTP 服务容器
│   ├── grpc.go                     # GRPC 服务容器
│   └── internal.go                 # 内部服务容器
├── Makefile                        # 利用Makefile管理项目[相当于一个脚手架]
├── utils                           # 放置一些通用的工具
│   └── md5.go  
```

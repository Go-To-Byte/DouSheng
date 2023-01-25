# DouSheng
第五届青训营项目——抖声

## 启动说明

### 1、项目配置文件
需要将`etc/dousheng.toml.template`配置模板文件复制粘贴为：`etc/dousheng.toml`，并且修改为自己的配置，如：
```toml
[app]
name = "DouSheng"
host = "0.0.0.0"
port = "8050"

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

* 大致目录结构
```text
DouSheng
├── apps            # 项目模块
│   ├── all         # 统一注册服务到IOC容器中
│   └── user        # 用户模块
│       ├── http    # Controller
│       └── impl    # ServiceImpl
├── cmd             # 通过CLI启动项目
├── conf            # 项目配置
├── dist            # 构建的项目路径
├── docs.example    # 相关文档
├── etc             # 环境变量
├── protocol        # 暴露的协议
└── version         # 项目版本相关
```

* 部分文件结构
```text
├── apps
│   ├── all         
│   │   └── auto_register.go    # IOC容器统一注册位置【驱动加载】
│   ├── app.go                  # IOC容器
│   └── user                    
│       ├── app.go              # 用户模块的名称【用于注入服务至IOC】
│       ├── http
│       │   ├── user.go         # Controller控制层的方法
│       │   └── http.go         # HTTP服务相关
│       ├── impl
│       │   ├── mapper.go       # Dao持久层
│       │   ├── mysql.go        
│       │   └── user.go         # ServiceImpl，用户模块业务接口的实现
│       ├── interface.go        # Service 业务层，用户模块的接口
│       └── model.go            # user 模块相关的 model
├── cmd
│   └── start.go                # CLI启动入口
├── conf
│   ├── config.go               # 配置文件对象【APP、MYSQL、Log】
│   └── load.go                 # 用于外部加载配置文件
├── etc
│   ├── dousheng.toml           # 配置文件【不提交Github】
│   └── dousheng.toml.template  # 配置文件模板【这里采用轻量级的toml】
├── Makefile                    # 使用工程化管理项目【方便开发】
├── protocol                    
│   ├── grpc.go                 # 对外提供GRPC服务，给内部使用【还未实现】
│   └── http.go                 # 对外提供HTTP服务，给外部使用
```

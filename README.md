# DouSheng
第五届青训营项目——抖声


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

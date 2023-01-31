# DouSheng
第五届青训营项目——抖声

```text
├─apps                                  # grpc服务
│  ├─comment                            # 评论模块
│  │  ├─.idea                           
│  │  ├─config                          # 配置文件
│  │  ├─dao                             # 持久化
│  │  │  ├─dal                          # gen 生成文件
│  │  │  │  ├─model                     # 数据表模型
│  │  │  │  └─query                     # 数据库操作
│  │  │  └─generator                    # 使用 gen
│  │  ├─init                            # 初始化
│  │  ├─middle                          # 中间层
│  │  ├─models                          # 服务所需的模型
│  │  ├─proto                           # protobuf 文件，及其生成文件
│  │  ├─service                         # grpc 服务逻辑代码
│  │  └─test                            # 测试
│  ├─favorite                           # 点赞模块
│  │  ├─.idea
│  │  ├─config
│  │  ├─dao
│  │  │  ├─dal
│  │  │  │  ├─model
│  │  │  │  └─query
│  │  │  └─generator
│  │  ├─init
│  │  ├─middle
│  │  ├─models
│  │  ├─proto
│  │  ├─service
│  │  └─test
│  ├─feed                               # feed 流模块
│  ├─message                            # 用户聊天模块
│  │  ├─.idea
│  │  ├─config
│  │  ├─dao
│  │  │  ├─dal
│  │  │  │  ├─model
│  │  │  │  └─query
│  │  │  └─generator
│  │  ├─init
│  │  ├─middle
│  │  ├─models
│  │  ├─proto
│  │  ├─service
│  │  └─test
│  ├─relation                           # 用户关系模块
│  │  ├─.idea
│  │  ├─config
│  │  ├─dao
│  │  │  ├─dal
│  │  │  │  ├─model
│  │  │  │  └─query
│  │  │  └─generator
│  │  ├─init
│  │  ├─middle
│  │  ├─models
│  │  ├─proto
│  │  ├─service
│  │  └─test
│  ├─user                               # 用户注册、登录模块
│  │  ├─.idea
│  │  ├─config
│  │  ├─dao
│  │  │  ├─dal
│  │  │  │  ├─model
│  │  │  │  └─query
│  │  │  └─generator
│  │  ├─init
│  │  ├─middle
│  │  ├─models
│  │  ├─proto
│  │  ├─service
│  │  └─test
│  └─video                              # 视频模块
│      ├─.idea
│      ├─config
│      ├─dao
│      │  ├─dal
│      │  │  ├─model
│      │  │  └─query
│      │  └─generator
│      ├─init
│      ├─middle
│      ├─models
│      ├─proto
│      ├─service
│      └─test
└─network                               # 提供 http 服务
```
# 视频互动服务
- 评论模块
  - 添加评论
  - 删除评论
  - 获取评论列表
- 点赞模块
  - 视频点赞
  - 取消点赞
  - 获取喜欢视频列表
## 启动说明

### 1、项目配置文件
需要将`etc/config.toml.template`配置模板文件复制粘贴为：`etc/config.toml`，并且修改为自己的配置，如：
```toml
[app.grpc]
host = ""
port = 8507

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

[consul.register]
registry_name = "interaction_service"
host = "127.0.0.1"
port = 8500
tags = ["互动服务", "interaction_service"]

[consul.discovers.api_rooter]
discover_name = "api_rooter"
address = "127.0.0.1:8500"

[consul.discovers.user_center]
discover_name = "user_center"
address = "127.0.0.1:8500"

[consul.discovers.video_service]
discover_name = "video_service"
address = "127.0.0.1:8500"
```

### 2、项目启动
本项目采用了Makefile工程化便于开发，电脑需要额外安装`make`指令。
* windows推荐安装教程：[推荐第二种方式，别忘记配置环境变量](https://tehub.com/a/aCYp1uw0tG)
* 命令行键入`make run`启动项目
* 当然，不想安装make指令，也可以直接键入`go run main.go start`启动项目


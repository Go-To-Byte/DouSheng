---
title: 拿到1个项目后，一般该如何运行？
author: Ciusyan
date: '2023-2-19'
---

# 拿到1个项目后，一般该如何运行？


先回答标题，当然，这里所说的是一般情况~
1. 提供配置文件
2. 安装项目所需环境以及创建数据库模型
3. 运行项目

那我们来看看，该如何启动`Dousheng`，这里面用到的一些工具，在这里不详细展开，请查看后面的章节，这里我们仅仅使用~

## 一、利用工作区模式方便本地开发🥅

* 拉取项目
```shell
# 拉取项目
git clone git@github.com:Go-To-Byte/DouSheng.git

cd DouSheng
```

* 下载服务所需依赖
```shell
# 以下载ApiRooter的依赖为例
cd api_rooter/

go mod tidy
```

* 初始化`go.work`文件
```shell
# 初始化 工作区
go work init dou_kit/ user_center/ video_service/
```
执行过后，你的`go.work`文件内容如下：

```text
go 1.19

use (
   user_center
   dou_kit
   video_service
)
```



当你初始化工作区之后，本身是三个独立的服务，在本地开发调试时，就可以直接当作是一个项目了（文件路由引用）

注：这里所说的是**当作**，实际上还是独立的服务。此文件不要推送至Git仓库


## 二、添加服务的配置文件📄

当新打开一个项目时，当然需要查看此项目是如何为程序提供配置的呐，应该很少有人把配置硬编码到代码中吧！！！


因为要保证相对的安全，敏感资源不被泄露，所以真正的配置文件是不会被推送至代码仓库的。

但是我们提供了配置文件的模板。按照模板，填写你的配置即可~

又因为几个服务是独立的，所以按理来说，每个服务都会有自己的配置文件，那我们先配置，让程序跑起来再说！


* 添加Api-Rooter的配置文件(根据模板填写你自己的配置)

```shell
cd api_rooter/etc/

cp config.toml.template config.toml
```

* 添加视频服务配置文件(根据模板填写你自己的配置)
```shell
cd video_service/etc/

cp config.toml.template config.toml
```

* 添加用户中心配置文件(根据模板填写你自己的配置)
```shell
cd user_center/etc/

cp config.toml.template config.toml
```

* 依次类推，待会需要启动什么服务，就对什么服务进行配置....



至此，你已经为你的服务提供了配置文件，来看看具体如何启动吧！


## 三、根据配置文件，安装所需环境、创建数据库

在启动项目前，我们需要有这些环境。(在你为对应服务提供配置文件的时候，应该已经注意到了)

1. mysql环境			 [用于数据存储]
2. mongodb环境           [用于token鉴权]
3. consul环境            [微服务的注册中心]



这里就不搬运大自然的产物了，可自行Google。但为了方便，推荐使用 docker



当然，肯定得先建立Mysql数据库所需要的表结构呐，这里不保姆式的教你创建数据库、创建表了。因为"咱们"都是：高级程序员~



创建对应的数据库，并且执行去执行sql脚本。

如果你觉得麻烦：可以把表创建在一个数据库中，总的脚本文件在项目中的：`dou_kit/docs/sql/tables.sql` 文件中。所有服务Mysql的配置就可以一致了

否则，你就需要切换到每一个服务，这里以`user_center`为例

```shell
cd user_center/

# 初始化操作
go run main.go init
```

当你运行了上面的命令，就会去读取`docs/sql/tables.sql`目录下的sql脚本。

也是一样的，你的mysql配置是什么，就会在你配置的数据库中执行脚本，若想将服务的数据库分开，那就可以使用多个数据库的配置。


## 四、利用Makefile工程化管理项目🤝



如果**安装好环境**，并且添加了**配置文件**，那么你就可以使用如下命令启动服务：



* 传统方式启动

```shell
# 以启动ApiRooter为例
cd api_rooter/

# 启动服务
go run main.go start
```



启动其他服务也是类似的做法，这里就不一一列举了。



* 工程化管理方式启动

上面这种方式我们需要手动执行脚本，我们可以利用预定义在`Makefile`文件中的脚本，来对项目进行工程化管理。

当然呐，也有使用的前置条件：[安装make指令(推荐第二种)](https://tehub.com/a/aCYp1uw0tG)



若安装完成后，可以利用如下指令，启动项目：
```shell
# 进入api_rooter项目中
cd api_rooter/

# 利用工程化的方式启动服务
make run
```

当然了，Makefile文件下，有很多指令，比如：`dep、init、gen...`可以先自行研究，在后面的章节再介绍。


至此，相信你已经启动好项目了。若还想继续了解咱们`Go-To-Byte的Dousheng项目`，咱们接着往下看~


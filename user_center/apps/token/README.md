# Token 认证模块

首先说明，为什么要单独拿出来做一个模块：
+ 服务大起来后，都需要来这里获取Token，而需要Token的可能有
  + 内部服务用户、SDK用户、HTTP用户...
+ 而且可以有很多种Token的机制：
  + 用户名密码、Access_Token...

所以考虑之后的扩展性，将其单独做成了一个模块。

这里会采用 mongodb + redis 的认证方式

## 安装mongodb

* 使用docker简单安装，docker自行安装
```shell
docker pull mongo
```

* 运行并进入客户端
```shell
docker run -itd -p 27017:27017 mongo
```

* 添加管理员
```shell
use admin # 使用数据库
db.createUser({user:"admin",pwd:"root",roles:["root"]}) # 创建认证用户
db.auth("admin", "root") # 验证用户
```

* 添加认证数据库用户
```shell
use usercenter
db.createUser({user: "usercenter", pwd: "123456", roles: [{ role: "dbOwner", db: "usercenter" }]})
db.auth("usercenter", "123456")
```

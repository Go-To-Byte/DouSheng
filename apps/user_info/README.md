# User 服务模块

* 注册接口

## IMPL

这个模块写完后，User Service 的具体实现，上层业务就基于 Service 进行编程，面向接口编程


```text
http
 |
Host Service（接口定义）
 |
impl（基于Mysql实现）
```
User Service 定义并把接口实现，有很多使用的方式：
* 用户内部模块，基于它封装更高一层的业务逻辑，比如发布服务
* user Service 对外暴露：HTTP协议（暴露给用户）
* User Service 对外暴露：GRPC（暴露给内部服务）

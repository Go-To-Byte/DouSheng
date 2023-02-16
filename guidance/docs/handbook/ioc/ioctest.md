---
title: 这是测试页面，对吧！！！！
author: Ciusyan
date: '2023-2-16'
---

# 这是测试页面，对吧！！！！

# Go 项目相关



## 测试相关



* 断言库：`assert`

```
github.com/stretchr/testify
```



* 处理Sql的连接池连接**[mysql](https://github.com/go-sql-driver/mysql)**



```go
// 连接池：driverConn是具体的一个连接对象，它维护着一个Tcp链接
// pool：[]*driverConn，要维护 pool里面的链接都是可用的，需定期检查我们的conn健康情况
// 若某一个 driverConn已经失效了，需要将 driverConn.Reset()，清空该结构体的数据，
// 再次Reconnection重新获取一个连接，让该conn借壳存活
// 避免 driverConn 结构体的内存申请和释放的开销成本，达到复用的效果

func (m *MySQL) getDBConn() (*sql.DB, error) {
   
   dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
      m.UserName, m.Password, m.Host, m.Port, m.Database)

   db, err := sql.Open("mysql", dsn)

   if err != nil {
      return nil, fmt.Errorf("链接mysql：%s, error：%s", dsn, err.Error())
   }

   db.SetMaxOpenConns(m.MaxOpenConn)
   db.SetMaxIdleConns(m.MaxIdleConn)
   db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
   db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
   
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   
   defer cancel()

   if err := db.PingContext(ctx); err != nil {
      return nil, fmt.Errorf("ping mysql：%s error, %s", dsn, err.Error())
   }
   
   return db, nil
}
```



* 顺带Mysql连接池的大致原理
    * // 连接池：driverConn是具体的一个连接对象，它维护着一个Tcp链接
    * // pool：[]*driverConn，要维护 pool里面的链接都是可用的，需定期检查我们的conn健康情况
    * // 若某一个 driverConn已经失效了，需要将 driverConn.Reset()，清空该结构体的数据，
    * // 再次Reconnection重新获取一个连接，让该conn借壳存活
    * // 避免 driverConn 结构体的内存申请和释放的开销成本，达到复用的效果





1. 定义接口、数据类型 【领域 + 接口】
2. 接口的具体实现
3. HTTP对外暴露接口
4. 可以通过CLI串起来



* 定义的结构体里面的字段尽量都是用户需要的。不要在结构体内部放置一些用户不知道的字段，会给用户使用写的接口的时候造成一定的困扰





* channel 的细节



```go
ch := make(chan os.Signal, 1)
// channel 是一种复合数据结构，可以当作一个容器，自定义的struct make(chan int, 1000)
// 如果没有close， gc是不会自动回收的
defer close(ch)

// Go 为了 并发编程设计的（CSP），依赖Channel作为数据通信的信道
// 所以这里出现了一个思维模式的转变：
//        单兵作战（只有一个Goroutine） ---> 团队作战（多个Goroutine 采用Channel 来通信）
//        比如main{ for range channel } 这个时候的channel仅仅是当作一个缓存，在单兵作战，必须要选择带缓冲区的channel，而且用完最好手动关闭
//        我们这里是团队作战：
//           signal.Notify 当作一个 Goroutine， g1
//           go manager.WaitStop(ch) 当作第二个 Goroutine， g2
//           g1 <-- ch --> g2 ch就类似于一个电话，g1 和 g2 用这个电话来团队作战

signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
// 后台挂起
go manager.WaitStop(ch)
```







## ORM



```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```





## TODO







###  2、使用CLI执行SQL生成数据库的表





### 6、为视频流添加分页器`pager`



### 7、添加oss云存储



## Docker



* [Windows安装](https://bbs.huaweicloud.com/blogs/329843)
* [安装教程2](https://www.ruyut.com/2022/09/windows-11-install-docker.html)



* [入门文章](https://mp.weixin.qq.com/s?__biz=MjM5MDc4MzgxNA==&mid=2458469748&idx=1&sn=f21e3b317dd65777ebe47f345afe5071&chksm=b1c2655d86b5ec4b80092b6df10d11c1ddae89b0a07ec7550c228a78358e3a68c1e1a6b993c9&mpshare=1&scene=1&srcid=0130dehf67sAZXveB6aevF35&sharer_sharetime=1675037991612&sharer_shareid=0cca238e877ef1606f383ca72a9c08f3&key=f4b4db11521317443db31fede02e87ace3fbef03cb06b6c4f74457b22e5628bbb117c4620876f1933ddeef663e4394cc6d52a64ca875f3ab4d0c379faddd146cdb00b70633a49c9aa76fdcc9d73f7c73aabebe25658871f178d56b360fb0fae027d62c3c03dc4bcb00463ba16946ee08207e6986793dff5db815b049aa8b15a3&ascene=1&uin=MzU5MDkzNjAx&devicetype=Windows+11+x64&version=6308011a&lang=zh_CN&exportkey=n_ChQIAhIQwJfbYYkKnKcNvWOusxPGsxL1AQIE97dBBAEAAAAAAFaeOOSk0OMAAAAOpnltbLcz9gKNyK89dVj0N2KMMIEdY4QbnraOXo7p6%2FVzWD8bp7QiRzQQx2aI5qiRo6dfIfHPnmrm3uYS5Pz5t851DUauRdDOHpbrIXI33f6G4bgbsKkP2k5o09B0bOXhLSIOexXur3B7x4jfUdDJYz6P%2BwwqJAuUgoeeV7bDufmHv43C1r%2FkdtK2P8NqX8hBuNhJmHy7OQcnzM07YRvk%2FNYd9cvtVvwCMbsJZAspNkSr%2FJH2%2F1g%2BBC78kMJb7H%2BL4QQvhucPmqIoXh%2F3AyU%2BEwi87TE3J0CdX9%2BOWc8N&acctmode=0&pass_ticket=HzmdwLGqRIyvz4AjJIA%2B7MYpmjalcc7objVNll%2BBDJNaXXhtNmsPkWe6LlkPLW65n79krQgSc8LYU9d%2FVJsvwg%3D%3D&wx_header=1&fontgear=2)










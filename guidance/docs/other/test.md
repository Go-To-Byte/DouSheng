---
title: 一份简单的测试报告
author:
date: '2023-2-24'
---

# 一份简单的测试报告
测试关系着系统的质量，系统质量又决定了线上的稳定性

测试：测试是避免事故的最后一道屏障

只要做好完备的测试，就可以避免大部分事故的发生

## 基准测试

* 测试电脑配置：
    * goos: windows
    * goarch: amd64
    * cpu: AMD Ryzen 7 5800H with Radeon Graphics


* 测试数据

| 基准测试名称                          | 重复次数 | 单次时间     | 堆内存     | 单词平均内存分配次数 |
| :------------------------------------ | :------: | :----------: | :--------: | :------------------: |
| BenchmarkVideoServiceImpl_FeedVideos  | 1287     | 954124 ns/op | 17791 B/op | 345 allocs/op        |
| BenchmarkVideoServiceImpl_PublishList | 1339     | 917430 ns/op | 15043 B/op | 272 allocs/op        |
| BenchmarkUserServiceImpl_UserInfo     | 4657     | 244547 ns/op | 4923 B/op  | 69 allocs/op         |
| BenchmarkUserServiceImpl_UserMap      | 4663     | 253973 ns/op | 6749 B/op  | 104 allocs/op        |


## 单元测试

主要包括：输入、测试单元、输出，以及与期望输出的校对，最后通过输出和期望值做校对，来反映我们的代码的运行结果是否和预期相符，是否正确。

每次编写新代码，加入了单元测试，一方面保证了新功能本身的正确性，由于历史代码也有相关的单元测试，如果整体的单元测试跑通了，又标明了新的代码没有影响原有代码的正确性

### 发布评论

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209263-167720906439412.png)

### 删除评论

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209202-16772090643761.png)

### 获取视频评论列表

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209207-16772090643762.png)

### 视频点赞

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209212-16772090643773.png)

### 取消点赞

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209217-16772090643774.png)

### 获取喜欢视频列表

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209221-16772090643775.png)

评论和点赞模块，测试文件命名符合规范，测试结果符合预期

## 接口测试

### 发布评论

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209226-16772090643776.png)

### 删除评论

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209232-16772090643777.png)

### 获取视频评论列表

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209239-16772090643778.png)

### 视频点赞

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209246-16772090643789.png)

### 取消点赞

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209252-167720906437810.png)

### 获取喜欢视频列表

![img](https://yczbest.cn/wp-content/uploads/2023/02/1677209257-167720906437811.png)

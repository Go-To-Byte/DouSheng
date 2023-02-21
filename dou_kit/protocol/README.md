# 对外暴露协议的方案

## HTTP服务

1. 启动方法

核心就是去监听HTTP服务的端口

```go

// Start 开启服务
func (s *HttpService) Start() error {
    
    // 调用执行前的逻辑
    if err := s.before.Before(s); err != nil {
        s.L.Errorf("Start前的逻辑执行失败")
        return err
    }
    
    s.L.Infof("[HTTP] 服务监听地址：%s", s.C.App.HTTP.Addr())
    if err := s.server.ListenAndServe(); err != nil {
        // 如果错误是正常关闭，则不报错
        if err == http.ErrServerClosed {
            s.L.Infof("服务 stop 成功")
            return nil
        }
	    return fmt.Errorf("开启 [HTTP] 服务异常：%s", err.Error())
    }
    
    return nil
}
```

2. 优雅关闭

什么时优雅关闭？为什么需要优雅关闭？

优雅关闭就是当程序监听到一些信号，会导致程序进程退出时，会像优雅的谦谦君子一样，做好退出的提示和善后

至于为什么要优雅关闭。除了友好的提示。最主要的还是因为：

比如有几个串行访问的请求，他们已经发送出来了，但是这个时候程序遇到了一些紧急关闭信号。

如果没有优雅关闭，已经发出来的请求，会得不到处理，程序就会退出。程序就会崩溃。很像不负责的渣男！！！

但如果有优雅关闭，当接收到关闭信号时，它大致会做两件事情：

+ 关闭新请求的通道，保证不会有新的请求进来
+ 处理完已发送的请求，做一个负责的男人



```go
func (s *HttpService) Stop() error {
	s.L.Infof("服务开始 stop")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.L.Warnf("关闭服务异常：%s", err)
		return err
	}
	return nil
}

```

## GRPC服务

1. 启动（监听GRPC服务所需的TCP端口）、优雅关闭

```go
func (s *GRPCService) Start() {
	// =====
	// 1、注册IOC中所有的GRPC服务
	// =====
	ioc.RegistryGrpc(s.server)

	// =====
	// 2、启动GRPC服务
	// =====
	addr := s.cfg.App.GRPC.Addr()
	listener, err := net.Listen("tcp", addr)
	s.l.Infof("[GRPC] 服务监听地址：%s", addr)
	if err != nil {
		s.l.Errorf("启动GRPC服务错误：%s", err.Error())
		return
	}

	// 理论上我们需要等待GRPC服务启动成功后，才注册此服务到注册中心
	// 但是 GRPC Server 并没有什么成功的回调通知

	// 所以我们只能假设GRPC Server 1秒后启动成功
	time.AfterFunc(1*time.Second, s.Register)

	// 注册服务开启健康检查
	grpc_health_v1.RegisterHealthServer(s.server, health.NewServer())

	// =====
	// 3、监听GRPC服务
	// =====
	if err = s.server.Serve(listener); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Infof("[GRPC] 服务关闭成功")
		}
		s.l.Errorf("开启 [GRPC] 服务异常：%s", err.Error())
		return
	}

}

// Stop 优雅关闭GRPC服务
func (s *GRPCService) Stop() error {
	s.l.Infof("优雅关闭 [GRPC] 服务")
	// 注销在consul中的 GRPC 服务
	s.DeRegister()
	s.server.GracefulStop()
	return nil
}
```

2. 服务注册到注册中心

既然是微服务，需要走RPC调用，肯定需要服务发现和服务注册的注册中心，(我们这里是Consul)

那么，是如何注册服务到Consul的呢、当退出时，又是如何注销服务的呢？

核心就是通过Consul的配置文件，拿到Consul的客户端。用Client提供的SDK，管理注册和退出。

注册时需要：实例化健康检查对象、注册对象

注销时需要：提供待注销服务注册时的ID

```go
func (s *GRPCService) Register() {

	// 生成对应grpc的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           s.cfg.App.GRPC.Addr(),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	consul := s.cfg.Consul

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = consul.Register.RegistryName
	registration.ID = xid.New().String()
	registration.Port = s.cfg.App.GRPC.Port
	registration.Tags = consul.Register.Tags
	registration.Address = s.cfg.App.GRPC.Host
	registration.Check = check

	err := s.client.Agent().ServiceRegister(registration)
	if err != nil {
		s.l.Errorf("注册GRPC服务到consul失败：%s", err)
		return
	}

	s.registration = registration
	s.l.Info("成功注册GRPC服务到consul")
}

// 服务注销
func (s *GRPCService) DeRegister() {
    if s.registration != nil {
        err := s.client.Agent().ServiceDeregister(s.registration.ID)
        if err != nil {
            s.l.Errorf("注销实例失败：%s", err)
        } else {
        s.l.Info("注销实例成功")
        }
    }
}

```

## 怎么利用AOP（面向切面编程）的思想抽取成公共代码了？

因为抽取到公共库里面了。为了使其他逻辑通用。利用了切面的方式，可供外界在启动前执行一些逻辑。

当然，你还可以定义执行后的切面...

比如：HTTP在启动前，我这里做了一个启动前的切面，并且提供了默认切面：

```go
    // 调用执行前的逻辑
	if err := s.before.Before(s); err != nil {
		s.L.Errorf("Start前的逻辑执行失败")
		return err
	}

// DefaultHttpStartBefore 返回一个默认的切面：HTTPStartBefore
func DefaultHttpStartBefore() StartFuncAop {
    return func(s *HttpService) error {
        // 1、将所有的Gin服务对象注册到IOC中
        option := ioc.NewGinOption(s.R, "/"+s.C.App.Name)
        ioc.RegistryGin(option)
        return nil
    }
}
```

在外界使用时：若无特殊处理逻辑，直接使用默认的Before切面即可。
但若有特殊逻辑，如video_service：添加了部分中间件，

```go
	cmd.HttpStartAop = func(s *protocol.HttpService) error {
		// 1、获取中间件
		mids, err := getMiddle()
		if err != nil {
			return err
		}
		// 2、将所有的Gin服务对象注册到IOC中
		option := ioc.NewGinOption(s.R, "/"+s.C.App.Name, mids...)
		option.NotVersion = true
		option.NotName = true
		ioc.RegistryGin(option)
		return nil
	}
```


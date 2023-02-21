# 程序启动交由CLI管理

## 为什么要用命令行来管理程序？

1. 简化main方法[自己编写的代码的启动入口]
2. 便于配合Makefile，实现工程化管理

## 如何使用CLI

原生的 flag 用起来不太好使。
推荐使用[A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)


两个核心点：

1. 定义命令

```go
var RootCmd = &cobra.Command{
	Use:     "dousheng",
	Long:    "极简版抖音Api",
	Short:   "doushengApi",
	Example: "go run main.go [Commands] [Flags]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
		}
		return cmd.Help()
	},
}
```

添加子命令某命令的子集命令：将StartCmd添加到RootCmd的子集命令：

```go
RootCmd.AddCommand(StartCmd)
```

2. 定义flags
```go
    f := RootCmd.PersistentFlags()
	f.BoolVarP(&vers, "version", "v", false, "用户中心的版本信息")
```

## 编写执行逻辑

这里拿启动逻辑距离，RunE函数中的逻辑，本该写在Main方法内部的。

```go
var StartCmd = &cobra.Command{
    Use:     "start",
    Long:    "启动 API服务",
    Short:   "启动 API服务",
    Example: "go run main start",
    RunE: func(cmd *cobra.Command, args []string) error {
        // ========
        // 1、加载配置文件&全局Logger对象
        // ========
        
        if err := conf.LoadConfigFromToml(configFile); err != nil {
            return err
        }
        if err := conf.LoadGlobalLogger(); err != nil {
            return err
        }
        
        // ========
        // 2、初始化IOC容器中的所有服务
        // ========
        
        // 类似于Mysql注入驱动的方式加载UserServiceImpl的 init方法，将依赖注入IOC
        // _ "github.com/Go-To-Byte/DouSheng/apps/user/impl"【User模块的ServiceImpl服务注入IOC】
        if err := ioc.InitAllDependencies(); err != nil {
			return err
        }
    
        // ========
        // 3、使用服务管理者来处理服务的关闭和开启
        // ========
        return managerStartAndStop()
    },
}
```


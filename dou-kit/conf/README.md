# 如何使用配置对象

## 先说使用

前提说明：此目录下的文件，load.go用于加载配置，其余文件均为配置对象。config.go为总配置对象

1. 使用配置文件（真实环境使用）

这里选择了：[轻量级的toml](https://toml.io/cn/v1.0.0)，规则少且配置强大，更易学习。
[toml解析库](https://github.com/BurntSushi/toml)

封装了`LoadConfigFromToml()`方法，用于加载全局配置对象，需要传入配置的`文件路径`
```go
func LoadConfigFromToml(filePath string) error {
    // 初始化全局对象
    cfg := NewDefaultConfig()
    _, err := toml.DecodeFile(filePath, cfg)
    if err != nil {
        return fmt.Errorf("load config file error，path：%s，%s", filePath, err)
    }
    
    return cfg.LoadGlobal()
}
```

2. 使用环境变量（一般测试使用）

这里选择的：[环境变量映工具](https://github.com/caarlos0/env)

封装了`LoadConfigFromEnv()`方法。
值得一提的是，如果使用的是.env的文件来防止环境变量。

若是GoLand，需要下载插件：`envfile`，启动时选择启动配置：选择envfile文件路径

若是vscode。可自行百度设置 debug 环境

```go
func LoadConfigFromEnv() error {
	config := NewDefaultConfig()
	if err := env.Parse(config); err != nil {
		return err
	}
	return config.LoadGlobal()
}
```

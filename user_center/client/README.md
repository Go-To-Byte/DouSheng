# 用户中心[user_center]SDK

使用此SDK的方式如下测试代码：

```go
// @Author: Ciusyan 2023/2/9
func TestUserCenter(t *testing.T) {
	should := assert.New(t)

	// 获取用户中心的客户端[从环境变量中获取配置]
	// 获取的配置去执行 kit 库中的 client.NewConfig(consulCfg, discoverName)
	userCenter, err := rpc.NewUserCenterFromEnv()

	if should.NoError(err) {
		tokenReq := token.NewValidateTokenRequest("xxx")
		// 这里主要是为了获取 用户ID
		validatedToken, err := userCenter.TokenService().ValidateToken(context.Background(), tokenReq)
		if should.NoError(err) {
			t.Log(validatedToken)
		}
	}
}

func TestUserCenter_GinAuthHandlerFunc(t *testing.T) {
    r := gin.New()
    // 使用 auth 中间件
    group := r.Group("/v1", userCenter.GinAuthHandlerFunc())
    
    group.GET("/", func(c *gin.Context) {
        c.String(200, "Get")
    })
	
    r.Run()
}

```

详细使用方式请看 client_test.go 文件
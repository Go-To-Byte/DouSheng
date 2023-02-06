# 用户中心[user_center]SDK

```go

// user_center 客户端
// 需要配置注册中心的[地址、服务名称]
// 利用注册中心 获取user_center的客户端
func TestToken(t *testing.T) {
	should := assert.New(t)

	// 配置Consul[地址、服务名称]
	consulConf := conf.NewDefaultConsul()
	consulConf.Host = os.Getenv("CONSUL_HOST")
	consulConf.Port, _ = strconv.Atoi(os.Getenv("CONSUL_PORT"))
	consulConf.Name = os.Getenv("CONSUL_NAME")

	// 根据注册中心的配置，获取用户中心的客户端
	client, err := rpc.NewClientSet(consulConf)
	
	// 下面就可以使用user_center提供的SDK了
	if should.NoError(err) {
		req := user.NewLoginAndRegisterRequest()
		req.Username = "ciusyan"
		req.Password = "111"

		serviceClient := client.Token()

		request := token.NewValidateTokenRequest("xxx")

		resp, err := serviceClient.ValidateToken(context.Background(), request)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	}
}

```
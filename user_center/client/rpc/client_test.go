// @Author: Ciusyan 2023/2/5
package rpc_test

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"
	"github.com/Go-To-Byte/DouSheng/user_center/client/rpc"
	"github.com/Go-To-Byte/DouSheng/user_center/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	should := assert.New(t)

	client, err := rpc.NewClientSet()

	if should.NoError(err) {
		req := user.NewLoginAndRegisterRequest()
		req.Username = "ciusyan"
		req.Password = "111"

		serviceClient := client.User()

		resp, err := serviceClient.Login(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	}
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
}

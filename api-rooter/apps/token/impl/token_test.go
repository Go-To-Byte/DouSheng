// @Author: Ciusyan 2023/2/6
package impl_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Go-To-Byte/DouSheng/dou_kit/conf"
	"github.com/Go-To-Byte/DouSheng/dou_kit/ioc"
	"github.com/Go-To-Byte/DouSheng/user_center/apps/user"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
	// 驱动加载所有需要放入IOC的实例
	_ "github.com/Go-To-Byte/DouSheng/api_rooter/common/all"
)

var (
	service token.ServiceServer
)

func TestIssueToken(t *testing.T) {
	should := assert.New(t)

	po := user.NewDefaultUserPo()
	po.Username = "xxxx"
	// 颁发
	req := token.NewIssueTokenRequest(po)
	iTk, err := service.IssueToken(context.Background(), req)

	if should.NoError(err) {
		t.Log(iTk)

		// 验证
		validateReq := token.NewValidateTokenRequest(iTk.AccessToken)
		vTk, err := service.ValidateToken(context.Background(), validateReq)
		if should.NoError(err) {
			t.Log(vTk)
		}
	}
}

func TestValidateToken(t *testing.T) {
	should := assert.New(t)

	req := token.NewValidateTokenRequest("d2Zh3HoEQXdm9yntFBFH3oOd")
	vTk, err := service.ValidateToken(context.Background(), req)

	if should.NoError(err) {
		t.Log(vTk)
	}
}

func init() {

	// 从环境变量种加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 初始化全局Logger
	if err := zap.DevelopmentSetup(); err != nil {
		panic(err)
	}

	// 初始化IOC
	if err := ioc.InitAllDependencies(); err != nil {
		panic(err)
	}

	// 获取依赖
	service = ioc.GetGrpcDependency(token.AppName).(token.ServiceServer)

}

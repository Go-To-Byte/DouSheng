// @Author: Ciusyan 2023/1/23
package impl

import (
	"context"
	"fmt"
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/bwmarrin/snowflake"
)

func (i *UserServiceImpl) CreateUser(ctx context.Context, request *user.LoginAndRegisterRequest) (*user.Token, error) {

	// 1、校验参数
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("参数校验失败：%s", err.Error())
	}

	newUser := user.NewUser()
	newUser.Name = request.Username
	newUser.Password = request.Password

	// 2、插入数据
	db := i.db.WithContext(ctx)
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, fmt.Errorf("自动迁移表失败：%s", err.Error())
	}
	result := db.Create(&newUser)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("插入数据失败：%s", err.Error())
	}
	// TODO：将 node 设置为单例，传参由配置文件提供
	node, err := snowflake.NewNode(1)
	id := node.Generate()
	token := id.String()
	return user.NewToken(newUser.ID, token), nil
}

func (i *UserServiceImpl) Login(ctx context.Context, request *user.LoginAndRegisterRequest) (*user.Token, error) {
	return nil, nil
}

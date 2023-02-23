// Author: BeYoung
// Date: 2023/2/21 21:11
// Software: GoLand

package impl

import (
	"context"
	"github.com/Go-To-Byte/DouSheng/user/apps/user"
	"github.com/Go-To-Byte/DouSheng/user/apps/user/impl/dal/model"
	_ "github.com/Go-To-Byte/DouSheng/user/apps/user/impl/init"
	"github.com/Go-To-Byte/DouSheng/user/apps/user/impl/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
)

func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (*user.InfoResponse, error) {
	u := model.User{
		ID:       req.UserId,
		Username: "",
		Passwd:   "",
	}

	// 查询用户
	zap.S().Debugf("user: %v", u.ID)
	r := s.UserFindById(u)
	if len(r) > 0 {
		userInfo := user.User{
			Id:            r[0].ID,
			Name:          r[0].Username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		return &user.InfoResponse{
			StatusCode: 0,
			StatusMsg:  "ok",
			User:       &userInfo,
		}, nil
	}

	return &user.InfoResponse{
		StatusCode: 1,
		StatusMsg:  "failed get user info",
		User:       nil,
	}, errors.New("failed get user info")
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	u := model.User{
		ID:       0,
		Username: req.Username,
		Passwd:   req.Password,
	}

	// 查询用户是否存在
	results := s.UserFindByName(u)
	if results == nil || len(results) > 1 {
		return &user.LoginResponse{
			StatusCode: 6,
			StatusMsg:  "user not found",
			UserId:     results[0].ID,
			Token:      "",
		}, nil
	}

	// 密码匹配，TODO: 使用加密算法匹配
	result := results[0]
	if u.Passwd != result.Passwd {
		return &user.LoginResponse{
			StatusCode: 6,
			StatusMsg:  "passwd mismatch",
			UserId:     results[0].ID,
			Token:      "",
		}, nil
	}

	return &user.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserId:     result.ID,
		Token:      strconv.FormatInt(result.ID, 10),
	}, nil
}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	u := model.User{
		ID:       models.Node.Generate().Int64(),
		Username: req.Username,
		Passwd:   req.Password,
	}

	// 查询用户是否存在
	results := s.UserFindByName(u)
	if results == nil || len(results) > 0 {
		return &user.RegisterResponse{
			StatusCode: 6,
			StatusMsg:  "registered",
			UserId:     results[0].ID,
			Token:      "",
		}, nil
	}

	// 添加用户, TODO: 密码加密
	zap.S().Debugf("add user: %+v", u)
	err := s.Add(u)
	if err != nil {
		return nil, err
	}

	// TODO: 生成token
	return &user.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "ok",
		UserId:     u.ID,
		Token:      strconv.FormatInt(u.ID, 10),
	}, nil
}

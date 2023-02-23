// @Author: Ciusyan 2023/2/6
package impl

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/Go-To-Byte/DouSheng/dou_kit/constant"

	"github.com/Go-To-Byte/DouSheng/api_rooter/apps/token"
)

var (
	DefaultTokenDuration = 10 * time.Minute
)

func (s *tokenServiceImpl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	newToken := token.NewToken(req, DefaultTokenDuration)
	_, err := s.col.InsertOne(ctx, newToken)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "颁发Token失败：%s", err.Error())
	}
	return newToken, nil
}

func (s *tokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {

	// 获取 Token
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		// TODO 调试信息
		s.log.Errorf("token: AccessToken：%s", req.AccessToken)

		s.log.Errorf("token: ValidateToken出现错误：%s", err.Error())
		return nil, status.Error(codes.Unknown,
			constant.Code2Msg(constant.ERROR_TOKEN_VALIDATE))
	}

	// 校验过期时间[采取双Token的机制]
	if !tk.IsExpired(tk.AccessTokenExpiredAt) {
		return tk, nil
	}

	// 来到这里说明 Access_Token 过期了，再看看 Refresh_Token
	if tk.IsExpired(tk.RefreshTokenExpiredAt) {
		return nil, status.Error(codes.Unauthenticated,
			constant.Code2Msg(constant.WRONG_TOKEN_EXPIRED))
	}

	// 来到这里说明 Access_Token 过期了，Refresh_Token 没过期

	if err = s.update(ctx, tk.Extend(DefaultTokenDuration)); err != nil {
		s.log.Errorf("token: ValidateToken出现错误：%s", err.Error())
		return nil, status.Error(codes.Unknown,
			constant.Code2Msg(constant.ERROR_TOKEN_VALIDATE))
	}

	return tk, nil
}

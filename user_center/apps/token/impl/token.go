// @Author: Ciusyan 2023/2/6
package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
)

var (
	DefaultTokenDuration = 10 * time.Minute
)

func (s *tokenServiceImpl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	newToken := token.NewToken(req, DefaultTokenDuration)
	_, err := s.col.InsertOne(ctx, newToken)
	if err != nil {
		return nil, fmt.Errorf("颁发Token失败：%s", err.Error())
	}
	return newToken, nil
}

func (s *tokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {

	// 获取 Token
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 校验过期时间[采取双Token的机制]
	if !tk.IsExpired(tk.AccessTokenExpiredAt) {
		return tk, nil
	}

	// 来到这里说明 Access_Token 过期了，再看看 Refresh_Token
	if tk.IsExpired(tk.RefreshTokenExpiredAt) {
		return nil, fmt.Errorf("token已过期")
	}

	// 来到这里说明 Access_Token 过期了，Refresh_Token 没过期

	if err = s.update(ctx, tk.Extend(DefaultTokenDuration)); err != nil {
		return nil, err
	}

	return tk, nil
}

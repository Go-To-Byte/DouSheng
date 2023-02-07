// @Author: Ciusyan 2023/2/6
package token

import (
	"github.com/Go-To-Byte/DouSheng/user_center/common/utils"
	"time"
)

const (
	AppName = "token"
)

func NewToken(req *IssueTokenRequest, expiredDuration time.Duration) *Token {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 5)

	return &Token{
		AccessToken:           utils.MakeBearer(24),
		IssueAt:               now.UnixMilli(),
		IssueBy:               req.Username,
		AccessTokenExpiredAt:  expired.UnixMilli(),
		RefreshToken:          utils.MakeBearer(32),
		RefreshTokenExpiredAt: refresh.UnixMilli(),
	}
}

func NewDefaultToken() *Token {
	return &Token{}
}

// IsExpired 判断AccessToken、RefreshToken有没有过期
func (s *Token) IsExpired(expiredAt int64) bool {
	return time.Now().UnixMilli() > expiredAt
}

// Extend 续约Token
func (t *Token) Extend(expiredDuration time.Duration) *Token {
	now := time.Now()
	// 更新过期时间
	t.AccessTokenExpiredAt = now.Add(expiredDuration).UnixMilli()
	t.RefreshTokenExpiredAt = now.Add(expiredDuration * 5).UnixMilli()

	t.UpdateAt = now.UnixMilli()
	t.UpdateBy = t.IssueBy

	return t
}

func NewIssueTokenRequest(username string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
	}
}

func NewValidateTokenRequest(ak string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: ak,
	}
}

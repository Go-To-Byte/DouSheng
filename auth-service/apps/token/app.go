// @Author: Ciusyan 2023/2/6
package token

import (
	"strconv"
	"time"

	"github.com/Go-To-Byte/DouSheng/auth-service/common/utils"
	"github.com/Go-To-Byte/DouSheng/dou-kit/constant"
)

const (
	AppName = "token"
)

func NewToken(req *IssueTokenRequest, expiredDuration time.Duration) *Token {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 5)

	// 额外携带个用户ID
	m := make(map[string]string, 1)
	m[constant.USER_ID] = strconv.Itoa(int(req.UserId))

	return &Token{
		AccessToken:           utils.MakeBearer(24),
		IssueAt:               now.UnixMilli(),
		IssueBy:               req.Username,
		AccessTokenExpiredAt:  expired.UnixMilli(),
		RefreshToken:          utils.MakeBearer(32),
		RefreshTokenExpiredAt: refresh.UnixMilli(),
		Meta:                  m,
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

func NewValidateTokenRequest(ak string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: ak,
	}
}

// GetUserId 从Meta中获取 用户ID
func (t *Token) GetUserId() int64 {
	userId, err := strconv.Atoi(t.Meta[constant.USER_ID])
	// 因为此方法是内部使用，所以基本不会乱传ID的情况
	if err != nil {
		panic(err)
	}
	return int64(userId)
}

func NewUIDResponse() *UIDResponse {
	return &UIDResponse{}
}

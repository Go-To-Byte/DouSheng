// Author: BeYoung
// Date: 2023/2/2 15:51
// Software: GoLand

package milddlewares

import (
	"errors"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var (
	ErrTokenMalformed   = errors.New("token is malformed")
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token is not valid yet")
	ErrTokenInvalidId   = errors.New("token has invalid id")
)

// // DecodePaddingAllowed 将分别切换用于解码JWT的编解码器。
// // 请注意，JWS RFC7515 声明令牌将使用不带填充的 Base64url 编码。
// // 不幸的是，一些实现 的 JWT 正在生成非标准令牌，因此需要支持解码。
// // 请注意，这是一个全局 变量，更新它将改变包级别的行为，并且也不是例行安全。
// // 要使用不推荐的解码，请在使用此包之前将此布尔值设置为“true”。
// var DecodePaddingAllowed bool

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		// token := c.Request.Header.Get("token")
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"msg": "授权已过期",
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}
		c.Set("token", claims)
		c.Set("userID", claims.ID)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(models.Config.Jwt.Key), // 可以设置过期时间
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(userID int64) (string, error) {
	claims := models.TokenClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go_to_byte",
			Subject:   "DouSheng",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)), // 15 天过期
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*models.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalidId
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalidId

	} else {
		return nil, ErrTokenInvalidId
	}

}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
		return j.CreateToken(claims.ID)
	}
	return "", ErrTokenInvalidId
}

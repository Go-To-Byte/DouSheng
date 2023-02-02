// Author: BeYoung
// Date: 2023/2/2 16:16
// Software: GoLand

package models

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

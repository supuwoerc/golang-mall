package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"server/component/config"
)

type CustomClaims struct {
	Uid      string `json:"uid"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// Sign 签发token的方法
func Sign(token *CustomClaims) (string, error) {
	secret := config.Config.GetString("jwt.secret")
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(secret))
	return tokenString, err
}

// Parse 解析token
func Parse(sign string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	} else {
		if tokenClaims != nil {
			if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
				return claims, nil
			} else {
				return nil, errors.New("tokenClaims is not valid")
			}
		} else {
			return nil, errors.New("tokenClaims is nil")
		}
	}
}

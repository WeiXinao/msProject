package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwter interface {
	GenAccessToken(uid int64, ip string) (string, error)
	GenRefreshToken(uid int64) (string, error)
	ParseToken(token string, claims jwt.Claims) error
	AccessExpire() time.Time
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId int64
	IP string
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	UserId int64
}

package jwtx

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var (
	ErrTokenParseFail = errors.New("token解析失败")
)

type JwtToken struct {
	// 单位s
	accessDDL  time.Time
	accessExp  time.Duration
	refreshExp time.Duration
	atKey      string
	rtKey      string
}

func NewJwtToken(atKey, rtKey string, aExp, rExp time.Duration) Jwter {
	return &JwtToken{
		accessExp:  aExp,
		refreshExp: rExp,
		atKey:      atKey,
		rtKey:      rtKey,
	}
}

func (j *JwtToken) AccessExpire() time.Time {
	return j.accessDDL
}

func (j *JwtToken) ParseToken(token string, claims jwt.Claims) error {
	var (
		tk  *jwt.Token
		err error
	)
	switch c := (claims).(type) {
	case *UserClaims:
		tk, err = jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
			return []byte(j.atKey), nil
		})
		if err != nil {
			return err
		}
	case *RefreshClaims:
		tk, err = jwt.ParseWithClaims(token, c, func(token *jwt.Token) (interface{}, error) {
			return []byte(j.rtKey), nil
		})
		if err != nil {
			return err
		}
	default:
		return ErrTokenParseFail
	}
	if tk == nil || !tk.Valid {
		return ErrTokenParseFail
	}
	return nil
}

func (j *JwtToken) GenAccessToken(uid int64) (string, error) {
	j.accessDDL = time.Now().Add(j.accessExp)
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.accessDDL),
		},
		UserId: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedString, err := token.SignedString([]byte(j.atKey))
	if err != nil {
		logx.Error("[GenAccessToken] ", err)
	}
	return signedString, err
}

func (j *JwtToken) GenRefreshToken(uid int64) (string, error) {
	claims := RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExp)),
		},
		UserId: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedString, err := token.SignedString([]byte(j.rtKey))
	if err != nil {
		logx.Error("[GenRefreshToken] ", err)
	}
	return signedString, err
}

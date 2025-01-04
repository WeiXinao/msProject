package middleware

import (
	"context"
	"errors"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
	"time"
)

var (
	ErrIllegalAuthorizationHeader = errors.New("无效的Authorization头")
	KeyMemberId                   = "memberId"
)

type AuthMiddlewareBuilder struct {
	jwtx.Jwter
}

func NewAuthMiddlewareBuilder(jwter jwtx.Jwter) *AuthMiddlewareBuilder {
	return &AuthMiddlewareBuilder{
		Jwter: jwter,
	}
}

func (a *AuthMiddlewareBuilder) Build(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	1. 从 header 中拿到 token
		token := r.Header.Get("Authorization")
		segments := strings.Split(token, " ")
		if len(segments) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			logx.Error("[Auth]", ErrIllegalAuthorizationHeader)
			return
		}
		token = segments[1]
		//	2. token 认证
		userClaims := jwtx.UserClaims{}
		err := a.Jwter.ParseToken(token, &userClaims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logx.Error("[Auth]", err)
			return
		}

		if userClaims.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//	3. 处理结果，认证通过 将信息放入 gin 的上下文，失败返回未登录
		ctx := context.WithValue(r.Context(), KeyMemberId, userClaims.UserId)
		next(w, r.WithContext(ctx))
	}
}

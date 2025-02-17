package middleware

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/WeiXinao/msProject/account/account"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrIllegalAuthorizationHeader = errors.New("无效的Authorization头")
	ErrIpNotConsistent = errors.New("IP 不一致")
	KeyMemberId                   = "memberId"
	KeyMemberName                 = "memberName"
	KeyOrganizationCode           = "organizationCode"
)

type AuthMiddlewareBuilder struct {
	jwtx.Jwter
	IngoreURLs []string
	UserClient loginservice.LoginService
	AccountClient account.Account
}

func NewAuthMiddlewareBuilder(
	jwter jwtx.Jwter,
	userClient loginservice.LoginService,
	accountClient account.Account,
	) *AuthMiddlewareBuilder {
	return &AuthMiddlewareBuilder{
		Jwter:      jwter,
		UserClient: userClient,
		AccountClient: accountClient,
	}
}

func (a *AuthMiddlewareBuilder) AddIngoreURLs(urls ...string) *AuthMiddlewareBuilder {
	a.IngoreURLs = append(a.IngoreURLs, urls...)
	return a
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

		// 先去查询 node 表，确认不适用登录控制的接口，不做登录认证了

		// 认证 IP
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logx.Error("[Auth]", err)
			return
		}
		if host != userClaims.IP {
			w.WriteHeader(http.StatusUnauthorized)
			logx.Error("[Auth]", ErrIpNotConsistent)
			return
		}

		if userClaims.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//	3. 处理结果，认证通过 将信息放入 gin 的上下文，失败返回未登录
		rsp, err := a.UserClient.MemberInfo(r.Context(), &userv1.MemberInfoRequest{
			Id: userClaims.UserId,
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			logx.Error("[Auth]", err)
			return
		}
		ctx := context.WithValue(r.Context(), KeyMemberId, userClaims.UserId)
		ctx = context.WithValue(ctx, KeyMemberName, rsp.Member.Name)
		ctx = context.WithValue(ctx, KeyOrganizationCode, rsp.Member.OrganizationCode)

		uri := strings.TrimPrefix(r.RequestURI, `/`) 

		if slice.Contains[string](a.IngoreURLs, uri) {
			next(w, r.WithContext(ctx))
			return
		}

		authNodesRsp, err := a.AccountClient.AuthNodesByMemberId(ctx, &account.AuthNodesByMemberIdRequest{
			MemberId: rsp.GetMember().GetId(),
		})
		if err != nil {
			err = respx.FromStatusErr(err)
			logx.Error("[Auth]", err)
			resp := respx.Fail(err.(*respx.Error))
			encoder := sonic.ConfigDefault.NewEncoder(w)
			encoder.Encode(resp)
			return 
		}
		
		if !slice.Contains(authNodesRsp.List, uri) {
			encoder := sonic.ConfigDefault.NewEncoder(w)
			encoder.Encode(respx.Fail(respx.ErrNotHasAuthority))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r.WithContext(ctx))
	}
}

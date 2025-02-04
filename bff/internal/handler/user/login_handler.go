package user

import (
	"context"
	"net"
	"net/http"

	"github.com/WeiXinao/msProject/pkg/respx"

	"github.com/WeiXinao/msProject/bff/internal/logic/user"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取 ip
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			logx.WithContext(r.Context()).Error("[handler LoginHandler] %#v", err)
			httpx.ErrorCtx(r.Context(), w, err)
		}
		ctx := context.WithValue(r.Context(), "ip", host) 

		l := user.NewLoginLogic(ctx, svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
		} else {
			httpx.OkJson(w, respx.Success(resp))
		}
	}
}

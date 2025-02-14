package auth

import (
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/auth"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewAuthListLogic(r.Context(), svcCtx)
		resp, err := l.AuthList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

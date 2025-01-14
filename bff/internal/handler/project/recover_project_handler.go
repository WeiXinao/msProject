package project

import (
	"github.com/WeiXinao/msProject/pkg/respx"
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/project"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecoverProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecycleProjectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewRecoverProjectLogic(r.Context(), svcCtx)
		resp, err := l.RecoverProject(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp.ProjectList))
		}
	}
}

package organization

import (
	"github.com/WeiXinao/msProject/pkg/respx"
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/organization"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MyOrgListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrgListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := organization.NewMyOrgListLogic(r.Context(), svcCtx)
		resp, err := l.MyOrgList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp.OrganizationList))
		}
	}
}

package project

import (
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/project"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProjectMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProjectMemberListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewProjectMemberLogic(r.Context(), svcCtx)
		resp, err := l.ProjectMember(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp))
		}
	}
}

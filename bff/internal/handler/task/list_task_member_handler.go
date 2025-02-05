package task

import (
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/task"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTaskMemberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTaskMemberReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewListTaskMemberLogic(r.Context(), svcCtx)
		resp, err := l.ListTaskMember(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp))
		}
	}
}

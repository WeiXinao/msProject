package user

import (
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/pkg/validatex"
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/user"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if !validatex.VerifyMobile(req.Mobile) {
			httpx.ErrorCtx(r.Context(), w, respx.IllegalMobile)
			return
		}

		if !validatex.VerifyEmailFormat(req.Email) {
			httpx.ErrorCtx(r.Context(), w, respx.IllegalEmail)
			return
		}

		if req.Password != req.Password2 {
			httpx.ErrorCtx(r.Context(), w, respx.InconsistentPwdAndConfirm)
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, respx.Success(resp))
		}
	}
}

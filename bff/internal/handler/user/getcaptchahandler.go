package user

import (
	"github.com/WeiXinao/msProject/pkg/paramx"
	"github.com/WeiXinao/msProject/pkg/respx"
	"net/http"

	"github.com/WeiXinao/msProject/bff/internal/logic/user"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 参数校验
		if !paramx.VerifyMobile(req.Mobile) {
			httpx.ErrorCtx(r.Context(), w, respx.IllegalMobile)
		}

		l := user.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, respx.Success(resp.Captcha))
		}
	}
}

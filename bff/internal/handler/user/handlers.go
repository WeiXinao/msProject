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

func GetCaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCaptchaReq
		err := httpx.Parse(r, &req)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// 参数校验
		if !validatex.VerifyMobile(req.Mobile) {
			httpx.ErrorCtx(r.Context(), w, respx.IllegalMobile)
		}

		l := user.NewGetCaptchaLogic(r.Context(), ctx)
		resp, err := l.GetCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, respx.Success(resp.Captcha))
		}
	}
}

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
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

		l := user.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, respx.Success(resp))
		}
	}
}

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, respx.Success(resp))
		}
	}
}

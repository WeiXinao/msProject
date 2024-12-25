package user

import (
	"context"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.GetCaptchaReq) (resp *types.GetCaptchaResp, err error) {
	captchaResp, err := l.svcCtx.UserClient.GetCaptcha(l.ctx, &userv1.CaptchaRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetCaptchaResp{
		Captcha: captchaResp.GetCode(),
	}
	return
}

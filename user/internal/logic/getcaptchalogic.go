package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *userv1.CaptchaRequest) (*userv1.CaptchaResponse, error) {
	// 生成随机验证码
	code := "123456"

	// 调用短信平台
	threading.GoSafe(func() {
		time.Sleep(2 * time.Second)
		l.Logger.Info("短信平台调用成功，发送短信")
		// redis 假设后续可能缓存在 mysql 当中，也可能存在 mongo 当中，也可能存在 memoached 当中
		// 存储验证码到 redis 中，过期时间 15 分钟
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		err := l.svcCtx.UserRepo.CacheCaptcha(ctx, req.Mobile, code, 15*time.Minute)
		cancel()
		if err != nil {
			l.Logger.Errorf("存储验证码失败, cause by：%s\n", err.Error())
		} else {
			l.Logger.Infof("将手机号和验证码存入redis成功：REGISTER_%s:%s", req.Mobile, code)
		}
	})

	return &userv1.CaptchaResponse{
		Code: code,
	}, nil
}

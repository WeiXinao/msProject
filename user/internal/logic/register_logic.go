package logic

import (
	"context"
	"errors"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/msProject/user/internal/repo"
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	userRepo := l.svcCtx.UserRepo
	// 1. 校验验证码
	err := userRepo.VerifyCaptcha(l.ctx, in.Mobile, in.Captcha)
	if errors.Is(err, repo.ErrVerifyCaptchaFail) {
		return nil, respx.ToStatusErr(respx.ErrVerifyCaptchaFail)
	}
	if err != nil {
		l.Error("[VerifyCaptcha] redis get error: ", err)
		return nil, err
	}

	// 2. 校验业务逻辑（邮箱是否被注册 账号是否被注册 手机号是否被注册）
	_, err = userRepo.GetMemberByEmail(l.ctx, in.GetEmail())
	if err == nil {
		return nil, respx.ToStatusErr(respx.ErrEmailDuplicated)
	}
	if !errors.Is(err, repo.ErrRecordNotFound) {
		l.Error("[GetMemberByEmail] db get error:", err)
		return nil, err
	}

	_, err = userRepo.GetMemberByAccount(l.ctx, in.GetName())
	if err == nil {
		return nil, respx.ToStatusErr(respx.ErrAccountDuplicated)
	}
	if !errors.Is(err, repo.ErrRecordNotFound) {
		l.Error("[GetMemberByAccount] db get error:", err)
		return nil, err
	}

	_, err = userRepo.GetMemberByMobile(l.ctx, in.GetMobile())
	if err == nil {
		return nil, respx.ToStatusErr(respx.ErrMobileDuplicated)
	}
	if !errors.Is(err, repo.ErrRecordNotFound) {
		l.Error("[GetMemberByMobile] db get error:", err)
		return nil, err
	}

	// 3. 执行业务 将数据存入 member 表，生成一个数据，存入组织表 organization
	pwd := encrypts.Md5(in.Password)

	err = l.svcCtx.UserRepo.CreateMemberAndOrganization(
		l.ctx,
		domain.Member{
			Account:       in.Name,
			Password:      pwd,
			Name:          in.Name,
			Mobile:        in.Mobile,
			Email:         in.Email,
			CreateTime:    time.Now(),
			LastLoginTime: time.Now(),
			Status:        domain.StatusMemberNormal,
		}, domain.Organization{
			Name:     in.Name + "个人项目",
			CTime:    time.Now(),
			Personal: domain.StatusOrganizationPersonal,
			Avatar:   "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fc-ssl.dtstatic.com%2Fuploads%2Fblog%2F202103%2F31%2F20210331160001_9a852.thumb.1000_0.jpg&refer=http%3A%2F%2Fc-ssl.dtstatic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1673017724&t=ced22fc74624e6940fd6a89a21d30cc5",
		})
	if err != nil {
		l.Error("[CreateMember] db create user error:", err)
		return nil, err
	}

	// 生成一个账户，账户的授权角色是成员，新生成一个角色（如果没有），同时将此角色的授权 node 生成
	return &userv1.RegisterResponse{}, nil
}

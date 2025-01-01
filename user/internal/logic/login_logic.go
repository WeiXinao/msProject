package logic

import (
	"context"
	"errors"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/msProject/user/internal/repo"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	// 1. 去数据库查询 账号密码是否正确
	pwd := encrypts.Md5(in.Password)
	m, err := l.svcCtx.UserRepo.GetMemberByAccountAndPwd(l.ctx, in.Account, pwd)
	if errors.Is(err, repo.ErrRecordNotFound) {
		return nil, respx.ToStatusErr(respx.ErrAccountOrPasswordIncorrect)
	}
	if err != nil {
		return nil, err
	}
	memberMsg := &userv1.MemberMessage{}
	err = copier.Copy(memberMsg, &m)
	if err != nil {
		return nil, err
	}
	memberMsg.LastLoginTime = m.LastLoginTime.UnixMilli()
	memberMsg.Code, _ = l.svcCtx.Encrypter.EncryptInt64(m.Id)

	// 2. 根据用户id查询组织
	orgs, err := l.svcCtx.UserRepo.GetOrganizationByMemId(l.ctx, m.Id)
	if err != nil {
		return nil, err
	}
	orgsMsg := slice.Map(orgs, func(idx int, src domain.Organization) *userv1.OrganizationMessage {
		orgMsg := &userv1.OrganizationMessage{}
		er := copier.Copy(orgMsg, &src)
		if er != nil {
			l.Error("[Login] ", er)
		}
		orgMsg.CreateTime = src.CTime.UnixMilli()
		orgMsg.Code, _ = l.svcCtx.Encrypter.EncryptInt64(src.Id)
		return orgMsg
	})

	// 3. 用 jwt 生成 token
	aToken, err := l.svcCtx.Jwter.GenAccessToken(m.Id)
	if err != nil {
		return nil, err
	}
	rToken, err := l.svcCtx.Jwter.GenRefreshToken(m.Id)
	if err != nil {
		return nil, err
	}
	tokenList := &userv1.TokenMessage{
		AccessToken:    aToken,
		RefreshToken:   rToken,
		TokenType:      "bearer",
		AccessTokenExp: l.svcCtx.Jwter.AccessExpire().Unix(),
	}
	return &userv1.LoginResponse{
		Member:           memberMsg,
		OrganizationList: orgsMsg,
		TokenList:        tokenList,
	}, nil
}

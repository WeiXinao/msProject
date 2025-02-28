// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: user.proto

package server

import (
	"context"

	"github.com/WeiXinao/msProject/user/internal/logic"
	"github.com/WeiXinao/msProject/user/internal/svc"
	v1_userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
)

type LoginServiceServer struct {
	svcCtx *svc.ServiceContext
	v1_userv1.UnimplementedLoginServiceServer
}

func NewLoginServiceServer(svcCtx *svc.ServiceContext) *LoginServiceServer {
	return &LoginServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServiceServer) GetCaptcha(ctx context.Context, in *v1_userv1.CaptchaRequest) (*v1_userv1.CaptchaResponse, error) {
	l := logic.NewGetCaptchaLogic(ctx, s.svcCtx)
	return l.GetCaptcha(in)
}

func (s *LoginServiceServer) Register(ctx context.Context, in *v1_userv1.RegisterRequest) (*v1_userv1.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *LoginServiceServer) Login(ctx context.Context, in *v1_userv1.LoginRequest) (*v1_userv1.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *LoginServiceServer) MyOrgList(ctx context.Context, in *v1_userv1.MyOrgListRequest) (*v1_userv1.MyOrgListResponse, error) {
	l := logic.NewMyOrgListLogic(ctx, s.svcCtx)
	return l.MyOrgList(in)
}

func (s *LoginServiceServer) MemberInfo(ctx context.Context, in *v1_userv1.MemberInfoRequest) (*v1_userv1.MemberInfoResponse, error) {
	l := logic.NewMemberInfoLogic(ctx, s.svcCtx)
	return l.MemberInfo(in)
}

func (s *LoginServiceServer) MemberInfosById(ctx context.Context, in *v1_userv1.MemberInfosByIdRequest) (*v1_userv1.MemberInfosByIdResponse, error) {
	l := logic.NewMemberInfosByIdLogic(ctx, s.svcCtx)
	return l.MemberInfosById(in)
}

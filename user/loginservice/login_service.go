// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package loginservice

import (
	"context"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CaptchaRequest      = userv1.CaptchaRequest
	CaptchaResponse     = userv1.CaptchaResponse
	LoginRequest        = userv1.LoginRequest
	LoginResponse       = userv1.LoginResponse
	MemberMessage       = userv1.MemberMessage
	OrganizationMessage = userv1.OrganizationMessage
	RegisterRequest     = userv1.RegisterRequest
	RegisterResponse    = userv1.RegisterResponse
	TokenMessage        = userv1.TokenMessage

	LoginService interface {
		GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	}

	defaultLoginService struct {
		cli zrpc.Client
	}
)

func NewLoginService(cli zrpc.Client) LoginService {
	return &defaultLoginService{
		cli: cli,
	}
}

func (m *defaultLoginService) GetCaptcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error) {
	client := userv1.NewLoginServiceClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}

func (m *defaultLoginService) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := userv1.NewLoginServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultLoginService) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := userv1.NewLoginServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
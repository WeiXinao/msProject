package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthNodesByMemberIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthNodesByMemberIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthNodesByMemberIdLogic {
	return &AuthNodesByMemberIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthNodesByMemberIdLogic) AuthNodesByMemberId(in *v1.AuthNodesByMemberIdRequest) (*v1.AuthNodesResponse, error) {
	return gozerox.HttpLogicWrapper(l.ctx, l, in, func(methodName string, logic *AuthNodesByMemberIdLogic, req *v1.AuthNodesByMemberIdRequest) (*v1.AuthNodesResponse, error) {
		nodes, err := l.svcCtx.AccoutRepo.FindAuthNodeStringListByMemberId(l.ctx, in.MemberId)
		if err != nil {
			return nil, err
		}
		return &v1.AuthNodesResponse{
			List: nodes,
		}, nil
	})
}

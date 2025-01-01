package project

import (
	"context"
	projectv1 "github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"
	"strings"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.IndexReq) (resp *types.IndexRsp, err error) {
	segments := strings.Split(req.Token, " ")
	if len(segments) != 2 {
		return nil, respx.ErrIllegalInput
	}
	projectRsp, err := l.svcCtx.ProjectClient.Index(l.ctx, &projectv1.IndexRequest{
		Token: segments[1],
	})
	if err != nil {
		return nil, err
	}
	resp = &types.IndexRsp{}
	err = copier.CopyWithOption(resp, &projectRsp, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}
	return
}

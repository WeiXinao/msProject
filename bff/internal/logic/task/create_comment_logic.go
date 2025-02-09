package task

import (
	"context"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/taskservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentRsp, err error) {
	_, err = l.svcCtx.TaskClient.CreateComment(l.ctx, &taskservice.CreateCommentRequest{
		TaskCode: req.TaskCode,
		CommentContent: req.Comment,
		Mentions: req.Mentions,
		MemberId: l.ctx.Value("memberId").(int64),
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic CreateComment] %#v", err)
		return
	}
	resp = &types.CreateCommentRsp{Success: true}
	return
}

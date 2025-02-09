package logic

import (
	"context"
	"time"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/msProject/task/internal/domain"
	"github.com/WeiXinao/msProject/task/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	taskCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetTaskCode())
	if err != nil {
		l.Errorf("[logic CreateComment] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	pl := domain.ProjectLog{
		MemberCode: in.GetMemberId(),
		Content: in.GetCommentContent(),
		Remark: in.GetCommentContent(),
		Type: "comment",
		CreateTime: time.Now().UnixMilli(),
		SourceCode: taskCode,
		ActionType: "task",
		ToMemberCode: 0,
		IsComment: domain.Comment,
		Icon: "plus",
		IsRobot: 0,
	}
	err = l.svcCtx.TaskRepo.SaveComment(l.ctx, pl)
	if err != nil {
		l.Errorf("[logic CreateComment] %#v", err)
		return nil, respx.ToStatusErr(respx.ErrInternalServer)
	}
	return &v1.CreateCommentResponse{}, nil
}

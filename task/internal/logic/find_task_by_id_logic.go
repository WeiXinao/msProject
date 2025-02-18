package logic

import (
	"context"

	"github.com/WeiXinao/msProject/api/proto/gen/task/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/WeiXinao/msProject/task/internal/svc"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTaskByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTaskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTaskByIdLogic {
	return &FindTaskByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindTaskByIdLogic) FindTaskById(in *v1.FindTaskByIdRequest) (*v1.TaskMessage, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *FindTaskByIdLogic, req *v1.FindTaskByIdRequest) (*v1.TaskMessage, error) {
		t, err := l.svcCtx.TaskRepo.FindTaskById(logic.ctx, in.GetTaskId())
		if err != nil {
			return nil, err
		}
		taskMsg := &v1.TaskMessage{}
		err = copier.Copy(&taskMsg, t.ToTaskDisplay(l.svcCtx.Encrypter))
		if err != nil {
			return nil, err
		}
		return taskMsg, nil
	})

}

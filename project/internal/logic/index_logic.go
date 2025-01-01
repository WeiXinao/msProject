package logic

import (
	"context"
	"github.com/WeiXinao/msProject/project/internal/domain"
	"github.com/jinzhu/copier"

	"github.com/WeiXinao/msProject/api/proto/gen/project/v1"
	"github.com/WeiXinao/msProject/project/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IndexLogic) Index(in *v1.IndexRequest) (*v1.IndexResponse, error) {
	menus, err := l.svcCtx.ProjectRepo.GetMenus(l.ctx)
	childs := domain.CovertChild(menus)
	mms := []*v1.MenuMessage{}
	err = copier.CopyWithOption(&mms, &childs, copier.Option{DeepCopy: true})
	return &v1.IndexResponse{
		Menus: mms,
	}, err
}

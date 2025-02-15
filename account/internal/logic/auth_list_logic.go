package logic

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthListLogic {
	return &AuthListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthListLogic) AuthList(in *v1.AuthListRequest) (*v1.AuthListResponse, error) {
	return gozerox.RpcLogicWrapper(l.ctx, l, in, func(methodName string, logic *AuthListLogic, req *v1.AuthListRequest) (*v1.AuthListResponse, error) {
		organCode, err := l.svcCtx.Encrypter.DecryptInt64(req.OrganizationCode)
		if err != nil {
			return nil, err
		}
		auths, total, err := l.svcCtx.AccoutRepo.FindAuthListByOrganizaitonCodePagination(logic.ctx, organCode, req.GetPage(), req.GetPageSize())
		if err != nil {
			return nil, err
		}
		displays := slice.Map(auths, func(idx int, src *domain.ProjectAuth) *domain.ProjectAuthDisplay {
			return src.ToDisplay(l.svcCtx.Encrypter)
		})
		list := []*v1.ProjectAuth{}
		err = copier.CopyWithOption(&list, displays, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, err
		}
		return &v1.AuthListResponse{
			List: list,	
			Total: total,
		}, nil
	})
}

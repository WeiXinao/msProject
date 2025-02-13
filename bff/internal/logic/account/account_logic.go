package account

import (
	"context"

	v1 "github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/bff/internal/middleware"
	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogic {
	return &AccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountLogic) Account(req *types.AccountReq) (resp *types.AccountRsp, err error) {
	accountRsp, err := l.svcCtx.AccountClient.Account(l.ctx, &v1.AccountRequest{
		MemberId: l.ctx.Value(middleware.KeyMemberId).(int64),
		OrganizationCode: l.ctx.Value(middleware.KeyOrganizationCode).(string),
		Page: int64(req.Page),
		PageSize: int64(req.PageSize),
		SearchType: int32(req.SearchType),
		DepartmentCode: req.DepartmentCode,
	})	
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic Account] %#v", err)
		return
	}
	resp = &types.AccountRsp{}
	err = copier.CopyWithOption(&resp, accountRsp, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic Account] %#v", err)
		err = respx.ErrInternalServer
		return 
	}
	if resp.AccountList == nil {
		resp.AccountList = []*types.MemberAccount{}
	}
	if resp.AuthList == nil {
		resp.AuthList = []*types.ProjectAuth{}
	}
	resp.Page = int64(req.Page)
	return
}

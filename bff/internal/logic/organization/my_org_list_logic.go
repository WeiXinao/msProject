package organization

import (
	"context"
	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/bff/internal/logic"
	"github.com/WeiXinao/msProject/pkg/respx"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/WeiXinao/msProject/bff/internal/svc"
	"github.com/WeiXinao/msProject/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyOrgListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyOrgListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyOrgListLogic {
	return &MyOrgListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyOrgListLogic) MyOrgList(req *types.GetOrgListReq) (resp *types.GetOrgListRsp, err error) {
	resp = &types.GetOrgListRsp{OrganizationList: []types.OrganizationList{}}
	memId, ok := l.ctx.Value("memberId").(int64)
	if !ok {
		l.Error("[logic MyOrgList]", logic.ErrTypeAssertFail)
		return resp, respx.ErrInternalServer
	}
	orgs, err := l.svcCtx.UserClient.MyOrgList(l.ctx, &userv1.MyOrgListRequest{MemId: memId})
	if err != nil {
		l.Error("[logic MyOrgList]", err)
		return resp, respx.FromStatusErr(err)
	}
	resp.OrganizationList = slice.Map(orgs.GetOrganizationList(),
		func(idx int, src *userv1.OrganizationMessage) types.OrganizationList {
			orgList := types.OrganizationList{}
			err = copier.Copy(&orgList, &src)
			if err != nil {
				l.Error("[logic MyOrgList]", err)
			}
			return orgList
		})
	return
}

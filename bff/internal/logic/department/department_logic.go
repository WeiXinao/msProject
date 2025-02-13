package department

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

type DepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentLogic {
	return &DepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DepartmentLogic) Department(req *types.DepartmentReq) (resp *types.DepartmentRsp, err error) {
	listDepartmentsResp, err := l.svcCtx.AccountClient.ListDepartments(l.ctx, &v1.ListDepartmentsReqeust{
		Page: req.Page,
		PageSize: req.PageSize,
		ParentDepartmentCode: req.Pcode,	
		OrganizationCode: l.ctx.Value(middleware.KeyOrganizationCode).(string),
	})
	if err != nil {
		l.Errorf("[logic Department] %#v", err)
		err = respx.FromStatusErr(err)
		return 
	}

	resp = &types.DepartmentRsp{}
	err = copier.CopyWithOption(&resp, listDepartmentsResp, copier.Option{DeepCopy: true})
	if err != nil {
		l.Errorf("[logic Department] %#v", err)
		err = respx.ErrInternalServer
		return
	}

	if resp.List == nil {
		resp.List = []*types.Department{}
	}

	resp.Page = req.Page
	return
}

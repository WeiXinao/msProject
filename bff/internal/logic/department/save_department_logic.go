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

type SaveDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveDepartmentLogic {
	return &SaveDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveDepartmentLogic) SaveDepartment(req *types.SaveDepartmentReq) (resp *types.Department, err error) {
	saveDeptRsp, err := l.svcCtx.AccountClient.SaveDepartment(l.ctx, &v1.SaveDepartmentRequest{
		Name: req.Name,
		DepartmentCode: req.DepartmentCode,
		ParentDepartmentCode: req.ParentDepartmentCode,
		OrganizationCode: l.ctx.Value(middleware.KeyOrganizationCode).(string),
	})
	if err != nil {
		err = respx.FromStatusErr(err)
		l.Errorf("[logic SaveDepartment] %#v", err)
		return
	}
	resp = &types.Department{}
	err = copier.Copy(&resp, saveDeptRsp)
	if err != nil {
		l.Errorf("[logic SaveDepartment] %#v", err)
		err = respx.ErrInternalServer
		return
	}
	return 
}

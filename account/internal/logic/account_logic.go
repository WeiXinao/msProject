package logic

import (
	"context"
	"time"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/svc"
	"github.com/WeiXinao/msProject/api/proto/gen/account/v1"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountLogic {
	return &AccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountLogic) Account(in *v1.AccountRequest) (*v1.AccountResponse, error) {
		// 1. 去 account 表查询 account
			organizationId, err := l.svcCtx.Encrypter.DecryptInt64(in.OrganizationCode)
			if err != nil {
				l.Errorf("[logic %s] %#v", "Account", err)
				return nil, err
			}
			departmentId := int64(0)
			if in.DepartmentCode != "" {
				departmentId, err = l.svcCtx.Encrypter.DecryptInt64(in.DepartmentCode)
				if err != nil {
					l.Errorf("[logic %s] %#v", "Account", err)
					return nil, err
				}
			}
			mas, total, err := l.svcCtx.AccoutRepo.FindMemberAccountList(l.ctx, in.GetSearchType(), organizationId, departmentId, in.GetPage(), in.GetPageSize())
			if err != nil {
				return nil, err
			}
			mids := slice.Map(mas, func(idx int, src *domain.MemberAccount) int64 {
				return src.MemberCode
			})
			memberInfosRsp, err := l.svcCtx.UserClient.MemberInfosById(l.ctx, &loginservice.MemberInfosByIdRequest{
				MIds: mids,
			})
			idToMemberInfoMap := slice.ToMap(memberInfosRsp.List, func(element *loginservice.MemberMessage) int64 {
				return element.GetId()
			})

		memberAccountDisplay := slice.Map(mas, func(idx int, src *domain.MemberAccount) *domain.MemberAccountDisplay {
				display := src.ToDisplay(l.svcCtx.Encrypter)
				if err != nil {
					l.Errorf("[logic %s] %#v", "Account", err)
				}
				memberMsg := idToMemberInfoMap[src.MemberCode]
				display.Avatar = memberMsg.GetAvatar()
				if src.DepartmentCode > 0 {
					department, err := l.svcCtx.AccoutRepo.FindDepartmentById(l.ctx, src.DepartmentCode)
					if err != nil {
						l.Errorf("[logic %s] %#v", "Account", err)
					}
					display.Departments = department.Name
				}
				return display
			})
			memberAccountMsgs := make([]*v1.MemberAccount, 0)
			err = copier.Copy(&memberAccountMsgs, memberAccountDisplay)
			if err != nil {
				l.Errorf("[logic %s] %#v", "Account", err)
				return nil, err
			}

		// 2. 去 auth 表查询 authList
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			orgCode, err := l.svcCtx.Encrypter.DecryptInt64(in.GetOrganizationCode()) 	
			if err != nil {
				l.Errorf("[logic %s] %#v", "Account", err)
				return nil, err
			}
			pas, err := l.svcCtx.AccoutRepo.FindAuthListByOrganizaitonCode(ctx, orgCode)
			if err != nil {
				l.Errorf("[logic %s] %#v", "Account", err)
				return nil, err
			}
		 	display :=	slice.Map(pas, func(idx int, src *domain.ProjectAuth) *domain.ProjectAuthDisplay {
				return src.ToDisplay(l.svcCtx.Encrypter)
			})
			msgs := make([]*v1.ProjectAuth, 0)
			err = copier.CopyWithOption(&msgs, display, copier.Option{DeepCopy: true})
			if err != nil {
				l.Errorf("[logic %s] %#v", "Account", err)
				return nil, err
			}

		if err != nil {
			l.Errorf("[logic %s] %#v", "Account", err)
			return nil, err
		}

		return &v1.AccountResponse{
			AuthList: msgs,	
			AccountList: memberAccountMsgs,
			Total: total,
		}, nil
}

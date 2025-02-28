package repo

import (
	"context"

	"github.com/WeiXinao/msProject/account/internal/domain"
	"github.com/WeiXinao/msProject/account/internal/repo/dao"
	"github.com/jinzhu/copier"
)

type AccountRepo interface {
	FindAuthListByOrganizaitonCode(ctx context.Context, orgCode int64) ([]*domain.ProjectAuth, error)
	FindAuthListByOrganizaitonCodePagination(ctx context.Context, orgCode int64, page int64, pageSize int64) ([]*domain.ProjectAuth, int64, error)
	FindMemberAccountList(ctx context.Context, searchType int32, organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*domain.MemberAccount, int64, error)

	FindDepartmentById(ctx context.Context, departmentId int64) (domain.Department, error)
	ListDepartments(ctx context.Context, orgCode int64, parentDeptCode int64, page int64, pageSize int64) ([]*domain.Department, int64, error)
	SaveDepartment(ctx context.Context, dept domain.Department) (domain.Department, error)

	GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error)

	FindAllProjectNodes(ctx context.Context) ([]*domain.ProjectNode, error)
	FindAuthNodeStringList(ctx context.Context, authId int64) ([]string, error)
	UpdateAuthNodes(ctx context.Context, authId int64, nodes []string) error
	FindAuthNodeStringListByMemberId(ctx context.Context, memberId int64) ([]string, error)
}

type accountRepo struct {
	dao dao.AccountDao
}

// FindAuthNodeStringListByMemberId implements AccountRepo.
func (a *accountRepo) FindAuthNodeStringListByMemberId(ctx context.Context, memberId int64) ([]string, error) {
	return a.dao.FindAuthNodeStringListByMemberId(ctx, memberId)
}

// UpdateAuthNodes implements AccountRepo.
func (a *accountRepo) UpdateAuthNodes(ctx context.Context, authId int64, nodes []string) error {
	return a.dao.UpdateAuthNodes(ctx, authId, nodes)
}

// FindAuthNodeStringList implements AccountRepo.
func (a *accountRepo) FindAuthNodeStringList(ctx context.Context, authId int64) ([]string, error) {
	return a.dao.FindAuthNodeStringList(ctx, authId)
}

// FindAllProjectNodes implements AccountRepo.
func (a *accountRepo) FindAllProjectNodes(ctx context.Context) ([]*domain.ProjectNode, error) {
	pns, err := a.dao.FindAllProjectNodes(ctx)
	if err != nil {
		return nil, err
	}
	pnds := make([]*domain.ProjectNode, 0)
	err = copier.Copy(&pnds, pns)
	if err != nil {
		return nil, err
	}
	return pnds, nil
}

func (a *accountRepo) GetMenus(ctx context.Context) ([]*domain.ProjectMenu, error) {
	menus, err := a.dao.GetMenus(ctx)
	if err != nil {
		return nil, err
	}
	menuDos := []*domain.ProjectMenu{}
	err = copier.Copy(&menuDos, &menus)
	return menuDos, err
}

// FindAuthListByOrganizaitonCodePagination implements AccountRepo.
func (a *accountRepo) FindAuthListByOrganizaitonCodePagination(ctx context.Context, orgCode int64, page int64, pageSize int64) ([]*domain.ProjectAuth, int64, error) {
	pa, total, err := a.dao.FindAuthListByOrganizaitonCodePagination(ctx, orgCode, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	pas := make([]*domain.ProjectAuth, 0)
	err = copier.CopyWithOption(&pas, pa, copier.Option{DeepCopy: true})
	return pas, total, err
}

// SaveDepartment implements AccountRepo.
func (a *accountRepo) SaveDepartment(ctx context.Context, dept domain.Department) (domain.Department, error) {
	deptEty := dao.Department{}
	err := copier.Copy(&deptEty, dept)
	if err != nil {
		return domain.Department{}, err
	}
	d, err := a.dao.SaveDepartment(ctx, deptEty)
	if err != nil {
		return domain.Department{}, err
	}
	deptDmn := domain.Department{}
	err = copier.Copy(&deptDmn, d)
	if err != nil {
		return domain.Department{}, err
	}
	return deptDmn, nil
}

// ListDepartments implements AccountRepo.
func (a *accountRepo) ListDepartments(ctx context.Context, orgCode int64, parentDeptCode int64, page int64, pageSize int64) ([]*domain.Department, int64, error) {
	d, total, err := a.dao.ListDepartments(ctx, orgCode, parentDeptCode, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	depts := make([]*domain.Department, 0)
	err = copier.CopyWithOption(&depts, d, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, 0, err
	}
	return depts, total, nil
}

// FindDepartmentById implements AccountRepo.
func (a *accountRepo) FindDepartmentById(ctx context.Context, departmentId int64) (domain.Department, error) {
	d, err := a.dao.FindDepartmentById(ctx, departmentId)
	if err != nil {
		return domain.Department{}, err
	}
	department := domain.Department{}
	err = copier.Copy(&department, d)
	return department, err
}

// FindMemberAccountList implements AccountRepo.
func (a *accountRepo) FindMemberAccountList(ctx context.Context, searchType int32, organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*domain.MemberAccount, int64, error) {
	mas, total, err := a.dao.FindMemberAccountList(ctx, searchType, organizationCode, departmentCode, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	memberAccounts := make([]*domain.MemberAccount, 0)
	err = copier.CopyWithOption(&memberAccounts, mas, copier.Option{DeepCopy: true})
	return memberAccounts, total, err
}

// FindAuthListByOrganizaitonCode implements AccountRepo.
func (a *accountRepo) FindAuthListByOrganizaitonCode(ctx context.Context, orgCode int64) ([]*domain.ProjectAuth, error) {
	pa, err := a.dao.FindAuthListByOrganizaitonCode(ctx, orgCode)
	if err != nil {
		return nil, err
	}
	pas := make([]*domain.ProjectAuth, 0)
	err = copier.CopyWithOption(&pas, pa, copier.Option{DeepCopy: true})
	return pas, err
}

func NewAccountRepo(dao dao.AccountDao) AccountRepo {
	return &accountRepo{
		dao: dao,
	}
}

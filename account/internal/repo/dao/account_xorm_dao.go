package dao

import (
	"context"
	"fmt"

	"github.com/WeiXinao/msProject/pkg/ormx"
	"github.com/WeiXinao/xkit/slice"
	"xorm.io/xorm"
)

type accountXormDao struct {
	db *xorm.Engine
}

// FindAuthNodeStringListByMemberId implements AccountDao.
func (a *accountXormDao) FindAuthNodeStringListByMemberId(ctx context.Context, memberId int64) ([]string, error) {
	nodes := make([]string, 0)
	err := a.db.Context(ctx).Table("ms_member_account").
		Join("inner", "ms_project_auth_node", "ms_member_account.authorize = ms_project_auth_node.auth").
		Where("ms_member_account.member_code = ?", memberId).Cols("node").Find(&nodes)
	return nodes, err
}

// UpdateAuthNodes implements AccountDao.
func (a *accountXormDao) UpdateAuthNodes(ctx context.Context, authId int64, nodes []string) error {
	return ormx.NewTxSession(a.db.Context(ctx)).EngineTx(func(engine *xorm.Engine) error {
		_, err := engine.Where("auth = ?", authId).Delete(new(ProjectAuthNode))
		if err != nil {
			return err
		}
		_, err = engine.Insert(slice.Map(nodes, func(idx int, src string) any {
			return ProjectAuthNode{
				Auth: authId,
				Node: src,
			}
		})...)
		return err
	})
}

// FindAuthNodeStringList implements AccountDao.
func (a *accountXormDao) FindAuthNodeStringList(ctx context.Context, authId int64) ([]string, error) {
	nodes := make([]string, 0)
	err := a.db.Context(ctx).Table("ms_project_auth_node").Where("auth = ?", authId).Cols("node").Find(&nodes)
	return nodes, err
}

// FindAllProjectNodes implements AccountDao.
func (a *accountXormDao) FindAllProjectNodes(ctx context.Context) ([]*ProjectNode, error) {
	pns := make([]*ProjectNode, 0)
	err := a.db.Find(&pns)
	if err != nil {
		return nil, err
	}
	return pns, nil
}

func (a *accountXormDao) GetMenus(ctx context.Context) ([]*ProjectMenu, error) {
	meuns := []*ProjectMenu{}
	err := a.db.Context(ctx).OrderBy("pid, sort ASC, id ASC").Find(&meuns)
	return meuns, err
}

// FindAuthListByOrganizaitonCodePagination implements AccountDao.
func (a *accountXormDao) FindAuthListByOrganizaitonCodePagination(ctx context.Context, orgCode int64, page int64, pageSize int64) ([]*ProjectAuth, int64, error) {
	projectAuths := make([]*ProjectAuth, 0)
	whereCond, args := "organization_code = ?", []any{orgCode}
	err := a.db.Context(ctx).
		Where(whereCond, args...).
		Limit(int(pageSize), int((page-1)*pageSize)).
		Find(&projectAuths)
	if err != nil {
		return nil, 0, err
	}
	total, err := a.db.Context(ctx).Where(whereCond, args...).Count(new(ProjectAuth))
	return projectAuths, total, err
}

// SaveDepartment implements AccountDao.
func (a *accountXormDao) SaveDepartment(ctx context.Context, dept Department) (Department, error) {
	d := Department{}
	whereCond, args := "organization_code = ? AND pcode = ? AND name = ?", []any{dept.OrganizationCode, dept.Pcode, dept.Name}
	if dept.Pcode == 0 {
		whereCond, args = "organization_code = ? AND name = ?", []any{dept.OrganizationCode, dept.Name}
	}
	has, err := a.db.Context(ctx).
		Where(whereCond, args...).
		Get(&d)
	if err != nil {
		return Department{}, err
	}
	if has {
		return d, nil
	}
	_, err = a.db.InsertOne(&dept)
	if err != nil {
		return Department{}, err
	}
	return dept, nil
}

// ListDepartments implements AccountDao.
func (a *accountXormDao) ListDepartments(ctx context.Context, orgCode int64, parentDeptCode int64, page int64, pageSize int64) ([]*Department, int64, error) {
	whereCond, args := "organization_code = ?", []any{orgCode}
	if parentDeptCode > 0 {
		whereCond, args = "organization_code = ? AND pcode = ?", []any{orgCode, parentDeptCode}
	}
	total, err := a.db.Context(ctx).Where(whereCond, args...).Count(new(Department))
	if err != nil {
		return nil, 0, err
	}
	depts := make([]*Department, 0)
	err = a.db.Context(ctx).
		Where(whereCond, args...).
		Limit(int(pageSize), int((page-1)*pageSize)).
		Find(&depts)
	return depts, total, err
}

// FindDepartmentById implements AccountDao.
func (a *accountXormDao) FindDepartmentById(ctx context.Context, departmentId int64) (Department, error) {
	d := Department{}
	_, err := a.db.ID(departmentId).Get(&d)
	return d, err
}

// FindMemberAccountList implements AccountDao.
func (a *accountXormDao) FindMemberAccountList(ctx context.Context, searchType int32, organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*MemberAccount, int64, error) {
	whereCond := "organization_code = ? AND "
	switch searchType {
	case 2:
		whereCond += "department_code IS NULL"
	case 3:
		whereCond += "status = 0"
	case 4:
		whereCond += fmt.Sprintf("status = 1 AND department_code = %d", departmentCode)
	default:
		whereCond += "status = 1"
	}
	ma := make([]*MemberAccount, 0)
	offset := (page - 1) * pageSize
	err := a.db.Context(ctx).
		Where(whereCond, organizationCode).
		Limit(int(pageSize), int(offset)).
		Find(&ma)
	if err != nil {
		return nil, 0, err
	}
	total, err := a.db.Context(ctx).Where(whereCond, organizationCode).Count(new(MemberAccount))
	return ma, total, err
}

// FindAuthListByOrganizaitonCode implements AccountDao.
func (a *accountXormDao) FindAuthListByOrganizaitonCode(ctx context.Context, orgCode int64) ([]*ProjectAuth, error) {
	projectAuths := make([]*ProjectAuth, 0)
	err := a.db.Context(ctx).Where("organization_code = ?", orgCode).Find(&projectAuths)
	return projectAuths, err
}

func NewAccountXormDao(db *xorm.Engine) AccountDao {
	db.Sync(
		new(MemberAccount),
		new(ProjectAuth),
		new(ProjectAuthNode),
		new(ProjectNode),
		new(Department),
		new(DepartmentMember),
	)
	return &accountXormDao{
		db: db,
	}
}

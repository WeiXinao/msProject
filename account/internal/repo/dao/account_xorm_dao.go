package dao

import (
	"context"
	"fmt"

	"xorm.io/xorm"
)

type accountXormDao struct {
	db *xorm.Engine
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

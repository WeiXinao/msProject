package dao

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/ormx"
	"xorm.io/xorm"
)

type XormUserDao struct {
	db *xorm.Engine
}

func (x *XormUserDao) GetOrganizationByMemId(ctx context.Context, memId int64) ([]Organization, error) {
	orgs := make([]Organization, 0)
	err := x.db.Context(ctx).Where("member_id = ?", memId).Find(&orgs)
	return orgs, err
}

func (x *XormUserDao) GetMemberByAccountAndPwd(ctx context.Context, account string, pwd string) (Member, error) {
	m := Member{}
	has, err := x.db.Context(ctx).Where("account = ? and password = ?", account, pwd).Get(&m)
	if err != nil {
		return Member{}, err
	}
	if !has {
		return Member{}, ErrRecordNotFound
	}
	return m, nil
}

func (x *XormUserDao) CreateMemberAndOrganization(ctx context.Context, member Member, organization Organization) error {
	oldDB := x.db
	defer func() {
		x.db = oldDB
	}()
	return ormx.NewTxSession(x.db.Context(ctx)).Tx(func(session any) error {
		var ok bool
		sess, ok := session.(*xorm.Session)
		if !ok {
			return ErrTypeConvert
		}
		x.db = sess.Engine()

		memId, err := x.CreateMember(ctx, member)
		if err != nil {
			return err
		}
		organization.MemberId = memId
		_, err = x.CreateOrganization(ctx, organization)
		return err
	})
}

func (x *XormUserDao) CreateMember(ctx context.Context, member Member) (int64, error) {
	_, err := x.db.Context(ctx).Insert(&member)
	return member.Id, err
}

func (x *XormUserDao) CreateOrganization(ctx context.Context, organization Organization) (int64, error) {
	_, err := x.db.Context(ctx).Insert(&organization)
	return organization.Id, err
}

func (x *XormUserDao) GetMemberByMobile(ctx context.Context, mobile string) (Member, error) {
	member := Member{}
	has, err := x.db.Context(ctx).Where("mobile = ?", mobile).Get(&member)
	if err != nil {
		return Member{}, err
	}
	if !has {
		return Member{}, ErrRecordNotFound
	}
	return member, nil
}

func (x *XormUserDao) GetMemberByAccount(ctx context.Context, account string) (Member, error) {
	member := Member{}
	has, err := x.db.Context(ctx).Where("account = ?", account).Get(&member)
	if err != nil {
		return Member{}, err
	}
	if !has {
		return Member{}, ErrRecordNotFound
	}
	return member, nil
}

func (x *XormUserDao) GetMemberByEmail(ctx context.Context, email string) (Member, error) {
	member := Member{}
	has, err := x.db.Context(ctx).Where("email = ?", email).Get(&member)
	if err != nil {
		return Member{}, err
	}
	if !has {
		return Member{}, ErrRecordNotFound
	}
	return member, nil
}

func NewXormUserDao(db *xorm.Engine) UserDao {
	return &XormUserDao{
		db: db,
	}
}

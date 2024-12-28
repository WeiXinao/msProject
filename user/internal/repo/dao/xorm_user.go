package dao

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/ormx"
	"xorm.io/xorm"
)

type XormUserDao struct {
	db xorm.Interface
}

func (x *XormUserDao) CreateMemberAndOrganization(ctx context.Context, member Member, organization Organization) error {
	oldDB := x.db
	defer func() {
		x.db = oldDB
	}()
	eg, ok := x.db.(*xorm.Engine)
	if !ok {
		return ErrTypeConvert
	}
	return ormx.NewTxSession(eg).Tx(func(session any) error {
		var ok bool
		x.db, ok = session.(xorm.Interface)
		if !ok {
			return ErrTypeConvert
		}
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
	_, err := x.db.Insert(&member)
	return member.Id, err
}

func (x *XormUserDao) CreateOrganization(ctx context.Context, organization Organization) (int64, error) {
	_, err := x.db.Insert(&organization)
	return organization.Id, err
}

func (x *XormUserDao) GetMemberByMobile(ctx context.Context, mobile string) (Member, error) {
	member := Member{}
	has, err := x.db.Where("mobile = ?", mobile).Get(&member)
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
	has, err := x.db.Where("account = ?", account).Get(&member)
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
	has, err := x.db.Where("email = ?", email).Get(&member)
	if err != nil {
		return Member{}, err
	}
	if !has {
		return Member{}, ErrRecordNotFound
	}
	return member, nil
}

func NewXormUserDao(db xorm.Interface) UserDao {
	return &XormUserDao{
		db: db,
	}
}

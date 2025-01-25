package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"github.com/WeiXinao/msProject/user/internal/repo/dao"
	"github.com/WeiXinao/xkit/slice"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"time"
)

type UserRepo interface {
	CacheCaptcha(ctx context.Context, mobile, captcha string, expire time.Duration) error
	VerifyCaptcha(ctx context.Context, mobile, expectCaptcha string) error
	GetMemberById(ctx context.Context, memId int64) (domain.Member, error)
	GetMemberByIds(ctx context.Context, memIds []int64) ([]*domain.Member, error)
	GetMemberByEmail(ctx context.Context, email string) (domain.Member, error)
	GetMemberByAccount(ctx context.Context, account string) (domain.Member, error)
	GetMemberByMobile(ctx context.Context, mobile string) (domain.Member, error)
	CreateMember(ctx context.Context, member domain.Member) (int64, error)
	CreateOrganization(ctx context.Context, organization domain.Organization) (int64, error)
	CreateMemberAndOrganization(ctx context.Context, member domain.Member, organization domain.Organization) error
	GetMemberByAccountAndPwd(ctx context.Context, account string, pwd string) (domain.Member, error)
	GetOrganizationByMemId(ctx context.Context, memId int64) ([]domain.Organization, error)
}

type userRepo struct {
	cache     cachex.Cache
	userCache *cache.UserCache
	userDao   dao.UserDao
	accessExp time.Duration
}

func (u *userRepo) GetMemberByIds(ctx context.Context, memIds []int64) ([]*domain.Member, error) {
	member, err := u.userDao.GetMemberByIds(ctx, memIds)
	if err != nil {
		return nil, err
	}
	memberDmns := make([]*domain.Member, 0)
	err = copier.Copy(&memberDmns, member)
	return memberDmns, err
}

func (u *userRepo) GetMemberById(ctx context.Context, memId int64) (domain.Member, error) {
	memStr, err := u.cache.Get(ctx, u.userCache.MemberKey(memId))
	memDmn := domain.Member{}
	if err == nil {
		err = json.Unmarshal([]byte(memStr), &memDmn)
		if err == nil {
			return memDmn, nil
		}
	}
	logx.WithContext(ctx).Error("[repo GetMemberById]", err)

	member, err := u.userDao.GetMemberById(ctx, memId)
	if err != nil {
		return domain.Member{}, err
	}
	memDmn, err = u.ToDomain(member)
	if err != nil {
		return domain.Member{}, err
	}
	threading.GoSafe(func() {
		er := u.cache.Put(ctx, u.userCache.MemberKey(memDmn.Id), memDmn, u.accessExp)
		if er != nil {
			logx.WithContext(ctx).Error("[repo GetMemberByAccountAndPwd]", er)
		}
	})
	return u.ToDomain(member)
}

func (u *userRepo) GetOrganizationByMemId(ctx context.Context, memId int64) ([]domain.Organization, error) {
	orgDmns := make([]domain.Organization, 0)
	orgsStr, err := u.cache.Get(ctx, u.userCache.MemberOrganizationKey(memId))
	if err == nil {
		if err == nil {
			err = json.Unmarshal([]byte(orgsStr), &orgDmns)
			if err == nil {
				return orgDmns, nil
			}
		}
	}

	if err != redis.Nil {
		logx.WithContext(ctx).Error("[GetOrganizationByMemId] repo ", err)
	}

	orgs, err := u.userDao.GetOrganizationByMemId(ctx, memId)
	if err != nil {
		return nil, err
	}
	orgDmns = slice.Map(orgs, func(idx int, src dao.Organization) domain.Organization {
		org := domain.Organization{}
		er := copier.Copy(&org, &src)
		if er != nil {
			logx.WithContext(ctx).Errorf("[GetOrganizationByMemId] repo: %w", ErrTypeConvert)
		}
		return org
	})
	threading.GoSafe(func() {
		er := u.cache.Put(ctx, u.userCache.MemberOrganizationKey(memId), orgDmns, u.accessExp)
		if err != nil {
			logx.WithContext(ctx).Error("[repo GetOrganizationByMemId]", er)
		}
	})
	return orgDmns, nil
}

func (u *userRepo) GetMemberByAccountAndPwd(ctx context.Context, account string, pwd string) (domain.Member, error) {
	memEntity, err := u.userDao.GetMemberByAccountAndPwd(ctx, account, pwd)
	if err != nil {
		return domain.Member{}, err
	}
	memDomain, err := u.ToDomain(memEntity)
	if err != nil {
		return domain.Member{}, err
	}
	threading.GoSafe(func() {
		er := u.cache.Put(ctx, u.userCache.MemberKey(memDomain.Id), memDomain, u.accessExp)
		if er != nil {
			logx.WithContext(ctx).Error("[repo GetMemberByAccountAndPwd]", er)
		}
	})
	return memDomain, nil
}

func (u *userRepo) CreateMemberAndOrganization(ctx context.Context, member domain.Member, organization domain.Organization) error {
	organEntity := dao.Organization{}
	err := copier.Copy(&organEntity, organization)
	if err != nil {
		return err
	}
	memberEntity, err := u.ToEntity(member)
	if err != nil {
		return err
	}
	return u.userDao.CreateMemberAndOrganization(ctx, memberEntity, organEntity)
}

func (u *userRepo) CreateOrganization(ctx context.Context, organization domain.Organization) (int64, error) {
	organEntity := dao.Organization{}
	err := copier.Copy(&organEntity, &organization)
	if err != nil {
		return 0, err
	}
	return u.userDao.CreateOrganization(ctx, organEntity)
}

func (u *userRepo) CreateMember(ctx context.Context, member domain.Member) (int64, error) {
	entity, err := u.ToEntity(member)
	if err != nil {
		return 0, err
	}
	return u.userDao.CreateMember(ctx, entity)
}

func (u *userRepo) GetMemberByAccount(ctx context.Context, account string) (domain.Member, error) {
	member, err := u.userDao.GetMemberByAccount(ctx, account)
	if err != nil {
		return domain.Member{}, err
	}
	memberDomain, err := u.ToDomain(member)
	if err != nil {
		return domain.Member{}, err
	}
	return memberDomain, nil
}

func (u *userRepo) GetMemberByMobile(ctx context.Context, mobile string) (domain.Member, error) {
	member, err := u.userDao.GetMemberByMobile(ctx, mobile)
	if err != nil {
		return domain.Member{}, err
	}
	memberDomain, err := u.ToDomain(member)
	if err != nil {
		return domain.Member{}, err
	}
	return memberDomain, nil
}

func (u *userRepo) GetMemberByEmail(ctx context.Context, email string) (domain.Member, error) {
	member, err := u.userDao.GetMemberByEmail(ctx, email)
	if err != nil {
		return domain.Member{}, err
	}
	memberDomain, err := u.ToDomain(member)
	if err != nil {
		return domain.Member{}, err
	}
	return memberDomain, nil
}

func (u *userRepo) VerifyCaptcha(ctx context.Context, mobile, expectCaptcha string) error {
	captcha, err := u.cache.Get(ctx, u.userCache.RegisterCaptchaKey(mobile))
	if errors.Is(err, redis.Nil) {
		return ErrVerifyCaptchaFail
	}
	if err != nil {
		return err
	}
	if captcha != expectCaptcha {
		return ErrVerifyCaptchaFail
	}
	return nil
}

func (u *userRepo) CacheCaptcha(ctx context.Context, mobile, captcha string, expire time.Duration) error {
	return u.cache.Put(ctx, u.userCache.RegisterCaptchaKey(mobile), captcha, expire)
}

func (u *userRepo) ToDomain(member dao.Member) (domain.Member, error) {
	memberDomain := domain.Member{}
	err := copier.Copy(&memberDomain, &member)
	if err != nil {
		return domain.Member{}, err
	}
	memberDomain.CreateTime = time.UnixMilli(member.CreateTime)
	memberDomain.LastLoginTime = time.UnixMilli(member.LastLoginTime)
	return memberDomain, nil
}

func (u *userRepo) ToEntity(member domain.Member) (dao.Member, error) {
	memberDao := dao.Member{}
	err := copier.Copy(&memberDao, &member)
	if err != nil {
		return dao.Member{}, err
	}
	memberDao.CreateTime = member.CreateTime.UnixMilli()
	memberDao.LastLoginTime = member.LastLoginTime.UnixMilli()
	return memberDao, nil
}

func NewUserRepo(dao dao.UserDao, cache cachex.Cache, userCache *cache.UserCache, accessExp time.Duration) UserRepo {
	return &userRepo{
		cache:     cache,
		userCache: userCache,
		userDao:   dao,
	}
}

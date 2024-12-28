package repo

import (
	"context"
	"errors"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/user/internal/domain"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"github.com/WeiXinao/msProject/user/internal/repo/dao"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type UserRepo interface {
	CacheCaptcha(ctx context.Context, mobile, captcha string, expire time.Duration) error
	VerifyCaptcha(ctx context.Context, mobile, expectCaptcha string) error
	GetMemberByEmail(ctx context.Context, email string) (domain.Member, error)
	GetMemberByAccount(ctx context.Context, account string) (domain.Member, error)
	GetMemberByMobile(ctx context.Context, mobile string) (domain.Member, error)
	CreateMember(ctx context.Context, member domain.Member) (int64, error)
	CreateOrganization(ctx context.Context, organization domain.Organization) (int64, error)
	CreateMemberAndOrganization(ctx context.Context, member domain.Member, organization domain.Organization) error
}

type userRepo struct {
	cache     cachex.Cache
	userCache *cache.UserCache
	userDao   dao.UserDao
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

func NewUserRepo(dao dao.UserDao, cache cachex.Cache, userCache *cache.UserCache) UserRepo {
	return &userRepo{
		cache:     cache,
		userCache: userCache,
		userDao:   dao,
	}
}

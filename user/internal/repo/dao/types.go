package dao

import (
	"context"
	"time"
)

type UserDao interface {
	GetMemberByEmail(ctx context.Context, email string) (Member, error)
	GetMemberByAccount(ctx context.Context, account string) (Member, error)
	GetMemberByMobile(ctx context.Context, mobile string) (Member, error)
	CreateMember(ctx context.Context, member Member) (int64, error)
	CreateOrganization(ctx context.Context, organization Organization) (int64, error)
	CreateMemberAndOrganization(ctx context.Context, member Member, organization Organization) error
	GetMemberByAccountAndPwd(ctx context.Context, account string, pwd string) (Member, error)
	GetOrganizationByMemId(ctx context.Context, memId int64) ([]Organization, error)
	GetMemberById(ctx context.Context, memId int64) (Member, error)
}

type Member struct {
	Id              int64  `xorm:"'id' pk notnull autoincr comment('系统前台用户表')"`
	Account         string `xorm:"'account' notnull default('') comment('用户登陆账号')"`
	Password        string `xorm:"'password' null default('') comment('登陆密码')"`
	Name            string `xorm:"'name' null default('') comment('用户昵称')"`
	Mobile          string `xorm:"'mobile' null default(null) comment('手机')"`
	Realname        string `xorm:"'realname' null default(null) comment('真实姓名')"`
	CreateTime      int64  `xorm:"'create_time' null default(null) comment('创建时间')" copier:"-"`
	Status          int    `xorm:"'status' null default(0) comment('状态')"`
	LastLoginTime   int64  `xorm:"'last_login_time' null default(null) comment('上次登陆时间')" copier:"-"`
	Sex             int    `xorm:"'sex' null default(0) comment('性别')"`
	Avatar          string `xorm:"'avatar' null text comment('头像')"`
	Idcard          string `xorm:"'idcard' null default(null) comment('身份证')"`
	Province        int    `xorm:"'province' null default(0) comment('省')"`
	City            int    `xorm:"'city' null default(0) comment('市')"`
	Area            int    `xorm:"'area' null default(0) comment('区')"`
	Address         string `xorm:"'address' null default(null) comment('所在地址')"`
	Description     string `xorm:"'description' null comment('备注')"`
	Email           string `xorm:"'email' null default(null) comment('邮箱')"`
	DingtalkOpenid  string `xorm:"'dingtalk_openid' null default(null) comment('钉钉openid')"`
	DingtalkUnionid string `xorm:"'dingtalk_unionid' null default(null) comment('钉钉unionid')"`
	DingtalkUserid  string `xorm:"'dingtalk_userid' null default(null) comment('钉钉用户id')"`
}

func (*Member) TableName() string {
	return "ms_member"
}

type Organization struct {
	Id          int64  `xorm:"'id' pk notnull autoincr"`
	Name        string `xorm:"'name' null default(null) comment('名称')"`
	Avatar      string `xorm:"'avatar' null text comment('头像')"`
	Description string `xorm:"'description' null default(null) comment('描述')"`
	MemberId    int64  `xorm:"'member_id' null default(null) comment('拥有者')"`
	CreateTime  int64  `xorm:"'create_time' null default(null) comment('创建时间')"`
	Personal    int32  `xorm:"'personal' null default(0) comment('是否个人项目')"`
	Address     string `xorm:"'address' null default(null) comment('地址')"`
	Province    int32  `xorm:"'province' null default(0) comment('省')"`
	City        int32  `xorm:"'city' null default(0) comment('市')"`
	Area        int32  `xorm:"'area' null default(0) comment('区')"`
}

func (o *Organization) CTime(time time.Time) {
	o.CreateTime = time.UnixMilli()
}

func (*Organization) TableName() string {
	return "ms_organization"
}

package dao

import "context"

type AccountDao interface {
    FindAuthListByOrganizaitonCode(ctx context.Context, orgCode int64) ([]*ProjectAuth, error)
	FindMemberAccountList(ctx context.Context, searchType int32, organizationCode int64, departmentCode int64, page int64, pageSize int64) ([]*MemberAccount, int64, error)

	FindDepartmentById(ctx context.Context, departmentId int64) (Department, error)
}

type MemberAccount struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    MemberCode int64 `xorm:"'member_code' bigint null
		default(null) comment('所属账号id')"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('所属组织')"`
    DepartmentCode int64 `xorm:"'department_code' bigint null default(null) comment('部门编号')"`
    Authorize string `xorm:"'authorize' varchar(255) null default(null) comment('角色')"`
    IsOwner int `xorm:"'is_owner' tinyint(1) null default(0) comment('是否主账号')"`
    Name string `xorm:"'name' varchar(255) null default(null) comment('姓名')"`
    Mobile string `xorm:"'mobile' varchar(12) null default(null) comment('手机号码')"`
    Email string `xorm:"'email' varchar(255) null default(null) comment('邮件')"`
    CreateTime int64 `xorm:"'create_time' bigint null   default(null) comment('创建时间')"`
    LastLoginTime int64 `xorm:"'last_login_time' bigint null default(null) comment('上次登录时间')"`
    Status int `xorm:"'status' tinyint(1) null default(null) comment('状态0禁用 1使用中')"`
    Description string `xorm:"'description' varchar(255) null default(null) comment('描述')"`
    Avatar string `xorm:"'avatar' varchar(255) null default(null) comment('头像')"`
    Position string `xorm:"'position' varchar(255) null default(null) comment('职位')"`
    Department string `xorm:"'department' varchar(255) null default(null) comment('部门')"`
}

func (*MemberAccount) TableName() string {
    return "ms_member_account"
}

type ProjectAuth struct {
    Id int64 `xorm:"'id' bigint unsigned autoincr pk notnull"`
    Title string `xorm:"'title' varchar(20) notnull comment('权限名称')"`
    Status int `xorm:"'status' tinyint unsigned null default(1) comment('状态(0:禁用,1:启用)')"`
    Sort int `xorm:"'sort' smallint unsigned null default(0) comment('排序权重')"`
    Desc string `xorm:"'desc' varchar(255) null default(null) comment('备注说明')"`
    CreateBy int64 `xorm:"'create_by' bigint unsigned null default(0) comment('创建人')"`
    CreateAt int64 `xorm:"'create_at' bigint null  default(null) comment('创建时间')"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('所属组织')"`
    IsDefault int `xorm:"'is_default' tinyint(1) null default(0) comment('是否默认')"`
    Type string `xorm:"'type' varchar(255) null default(null) comment('权限类型')"`
}

func (*ProjectAuth) TableName() string {
    return "ms_project_auth"
}

type ProjectAuthNode struct {
    Id int64 `xorm:"'id' bigint unsigned autoincr pk notnull"`
    Auth int64 `xorm:"'auth' bigint unsigned null default(null) comment('角色ID')"`
    Node string `xorm:"'node' varchar(200) null default(null) comment('节点路径')"`
}

func (*ProjectAuthNode) TableName() string {
    return "ms_project_auth_node"
}

type ProjectNode struct {
    Id int `xorm:"'id' int unsigned autoincr pk notnull"`
    Node string `xorm:"'node' varchar(100) null default(null) comment('节点代码')"`
    Title string `xorm:"'title' varchar(500) null default(null) comment('节点标题')"`
    IsMenu int `xorm:"'is_menu' tinyint unsigned null default(0) comment('是否可设置为菜单')"`
    IsAuth int `xorm:"'is_auth' tinyint unsigned null default(1) comment('是否启动RBAC权限控制')"`
    IsLogin int `xorm:"'is_login' tinyint unsigned null default(1) comment('是否启动登录控制')"`
    CreateAt int64 `xorm:"'create_at' bigint null default(null) comment('创建时间')"`
}

func (*ProjectNode) TableName() string {
    return "ms_project_node"
}

type Department struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('组织编号')"`
    Name string `xorm:"'name' varchar(30) null default(null) comment('名称')"`
    Sort int `xorm:"'sort' int null default(0) comment('排序')"`
    Pcode int64 `xorm:"'pcode' bigint null default(null) comment('上级编号')"`
    Icon string `xorm:"'icon' varchar(20) null default(null) comment('图标')"`
    CreateTime int64 `xorm:"'create_time' bigint null  default(null) comment('创建时间')"`
    Path string `xorm:"'path' text null comment('上级路径')"`
}

func (*Department) TableName() string {
    return "ms_department"
}

type DepartmentMember struct {
    Id int64 `xorm:"'id' bigint autoincr pk notnull"`
    DepartmentCode int64 `xorm:"'department_code' bigint null default(null) comment('部门id')"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('组织id')"`
    AccountCode int64 `xorm:"'account_code' bigint null default(null)  comment('成员id')"`
    JoinTime int64 `xorm:"'join_time' bigint null default(null) comment('加入时间')"`
    IsPrincipal int `xorm:"'is_principal' tinyint(1) null default(null) comment('是否负责人')"`
    IsOwner int `xorm:"'is_owner' tinyint(1) null default(0) comment('拥有者')"`
    Authorize string `xorm:"'authorize' varchar(255) null default(null) comment('角色')"`
}

func (*DepartmentMember) TableName() string {
    return "ms_department_member"
}
package dao

import (
	"context"
)

type ProjectDao interface {
	GetMenus(ctx context.Context) ([]*ProjectMenu, error)
	FindProjectByMemId(ctx context.Context, selectBy string, memId int64, page int64,
		size int64) ([]*ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, memId int64, page int64,
		size int64) ([]*ProjectAndMember, int64, error)

	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]ProjectTemplate, int64, error)
	FindInProTemIds(ctx context.Context, ids []int) ([]MsTaskStagesTemplate, error)

	SaveProject(ctx context.Context, project Project, projectMember ProjectMember) (Project, error)
	SaveProjectMember(ctx context.Context, projectMember *ProjectMember) error
	GetProjectAndMemberByPidAndMid(ctx context.Context, pid int64, mid int64) (ProjectAndMember, error)
	IsCollectedByPidAndMid(ctx context.Context, pid int64, mid int64) (bool, error)
}

type ProjectMenu struct {
	Id         int64  `xorm:"'id' pk notnull autoincr"`
	Pid        int64  `xorm:"'pid' notnull default(0) comment('父id')"`
	Title      string `xorm:"'title' notnull default('') comment('名称')"`
	Icon       string `xorm:"'icon' notnull default('') comment('菜单图标')"`
	Url        string `xorm:"'url' notnull default('') varchar(400) comment('链接')"`
	FilePath   string `xorm:"'file_path' null default(null) comment('文件路径')"`
	Params     string `xorm:"'params' null default('') comment('链接参数')"`
	Node       string `xorm:"'node' null default('#') comment('权限节点')"`
	Sort       int    `xorm:"'sort' null default(0) comment('菜单排序')"`
	Status     int    `xorm:"'status' null default(1) comment('状态(0:禁用,1:启用)')"`
	CreateBy   int64  `xorm:"'create_by' notnull default(0) comment('创建人')"`
	IsInner    int    `xorm:"'is_inner' null default(0) comment('是否内页')"`
	Values     string `xorm:"'values' null default(null) comment('参数默认值')"`
	ShowSlider int    `xorm:"'show_slider' null default(1) comment('是否显示侧栏')"`
}

func (*ProjectMenu) TableName() string {
	return "ms_project_menu"
}

type Project struct {
	Id                 int64   `xorm:"'id' bigint unsigned autoincr pk notnull"`
	Cover              string  `xorm:"'cover' varchar(255) null default(null) comment('封面')"`
	Name               string  `xorm:"'name' varchar(90) null default(null) comment('名称')"`
	Description        string  `xorm:"'description' text null default(null) comment('描述')"`
	AccessControlType  int     `xorm:"'access_control_type' tinyint null default(0) comment('访问控制l类型')"`
	WhiteList          string  `xorm:"'white_list' varchar(255) null default(null) comment('可以访问项目的权限组（白名单）')"`
	Order              int     `xorm:"'order' int unsigned null default(0) comment('排序')"`
	Deleted            int     `xorm:"'deleted' tinyint(1) null default(0) comment('删除标记')"`
	TemplateCode       int     `xorm:"'template_code' int null default(null) comment('项目类型')"`
	Schedule           float64 `xorm:"'schedule' double(5,2) null default(0.00) comment('进度')"`
	CreateTime         int64   `xorm:"'create_time' varchar(255) null default(null) comment('创建时间')"`
	OrganizationCode   int64   `xorm:"'organization_code' bigint null default(null) comment('组织id')"`
	DeletedTime        string  `xorm:"'deleted_time' varchar(30) null default(null) comment('删除时间')"`
	Private            int     `xorm:"'private' tinyint(1) null default(1) comment('是否私有')"`
	Prefix             string  `xorm:"'prefix' varchar(10) null default(null) comment('项目前缀')"`
	OpenPrefix         int     `xorm:"'open_prefix' tinyint(1) null default(0) comment('是否开启项目前缀')"`
	Archive            int     `xorm:"'archive' tinyint(1) null default(0) comment('是否归档')"`
	ArchiveTime        int64   `xorm:"'archive_time' bigint null default(null) comment('归档时间')"`
	OpenBeginTime      int     `xorm:"'open_begin_time' tinyint(1) null default(0) comment('是否开启任务开始时间')"`
	OpenTaskPrivate    int     `xorm:"'open_task_private' tinyint(1) null default(0) comment('是否开启新任务默认开启隐私模式')"`
	TaskBoardTheme     string  `xorm:"'task_board_theme' varchar(255) null default('default') comment('看板风格')"`
	BeginTime          int64   `xorm:"'begin_time' bigint null default(null) comment('项目开始日期')"`
	EndTime            int64   `xorm:"'end_time' bigint null default(null) comment('项目截止日期')"`
	AutoUpdateSchedule int     `xorm:"'auto_update_schedule' tinyint(1) null default(0) comment('自动更新项目进度')"`
}

func (*Project) TableName() string {
	return "ms_project"
}

type ProjectMember struct {
	Id          int64  `xorm:"'id' bigint autoincr pk notnull"`
	ProjectCode int64  `xorm:"'project_code' bigint null default(null) comment('项目id')"`
	MemberCode  int64  `xorm:"'member_code' bigint null default(null) comment('成员id')"`
	JoinTime    int64  `xorm:"'join_time' bigint null default(null) comment('加入时间')"`
	IsOwner     int64  `xorm:"'is_owner' bigint null default(0) comment('拥有者')"`
	Authorize   string `xorm:"'authorize' varchar(255) null default(null) comment('角色')"`
}

func (*ProjectMember) TableName() string {
	return "ms_project_member"
}

type ProjectAndMember struct {
	Project       `xorm:"extends"`
	ProjectMember `xorm:"extends"`
}

type ProjectCollection struct {
	Id          int64 `xorm:"'id' bigint autoincr pk notnull"`
	ProjectCode int64 `xorm:"'project_code' bigint null default(0) comment('项目id')"`
	MemberCode  int64 `xorm:"'member_code' bigint null default(0) comment('成员id')"`
	CreateTime  int64 `xorm:"'create_time' bigint null default(0) comment('加入时间')"`
}

func (*ProjectCollection) TableName() string {
	return "ms_project_collection"
}

type ProjectTemplate struct {
	Id               int    `xorm:"'id' int autoincr pk notnull"`
	Name             string `xorm:"'name' varchar(255) null default(null) comment('类型名称')"`
	Description      string `xorm:"'description' text null comment('备注')"`
	Sort             int    `xorm:"'sort' tinyint null default(0)"`
	CreateTime       int64  `xorm:"'create_time' bigint null default(0)"`
	OrganizationCode int64  `xorm:"'organization_code' bigint null default(null) comment('组织id')"`
	Cover            string `xorm:"'cover' varchar(511) null default(null) comment('封面')"`
	MemberCode       int64  `xorm:"'member_code' bigint null default(null) comment('创建人')"`
	IsSystem         int    `xorm:"'is_system' tinyint(1) null default(0) comment('系统默认')"`
}

func (*ProjectTemplate) TableName() string {
	return "ms_project_template"
}

type MsTaskStagesTemplate struct {
	Id                  int    `xorm:"'id' int autoincr pk notnull"`
	Name                string `xorm:"'name' varchar(255) null default(null) comment('类型名称')"`
	ProjectTemplateCode int    `xorm:"'project_template_code' int null default(0) comment('项目id')"`
	CreateTime          int64  `xorm:"'create_time' bigint null default(null)"`
	Sort                int    `xorm:"'sort' int null default(0)"`
}

func (*MsTaskStagesTemplate) TableName() string {
	return "ms_task_stages_template"
}

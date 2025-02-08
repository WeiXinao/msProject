package dao

import "context"


type FileDao interface {
    SaveFileAndSourceLink(ctx context.Context, file File, sourceLink SourceLink) error
    SaveFile(ctx context.Context, file *File) error
    FindByIds(ctx context.Context, ids []int64) (list []*File, err error)
    SaveSourceLink(ctx context.Context, link *SourceLink) error
	FindByTaskCode(ctx context.Context, taskCode int64) (list []*SourceLink, err error)
}

type File struct {
    Id int `xorm:"'id' int unsigned autoincr pk notnull"`
    PathName string `xorm:"'path_name' varchar(200) null default(null) comment('相对路径')"`
    Title  string `xorm:"'title' char(90) null default(null) comment('名称')"`
    Extension  string `xorm:"'extension' char(30) null  default(null) comment('扩展名')"`
    Size int `xorm:"'size' int unsigned null default(0) comment('文件大小')"`
    ObjectType string  `xorm:"'object_type' char(30) null default(null) comment('对象类型')"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('组织编码')"`
    TaskCode int64 `xorm:"'task_code' bigint null default(null) comment('任务编码')"`
    ProjectCode int64 `xorm:"'project_code' bigint null default(null) comment('项目编码')"`
    CreateBy int64 `xorm:"'create_by' bigint null default(null) comment('上传人')"`
    CreateTime int64 `xorm:"'create_time' bigint null   default(null) comment('创建时间')"`
    Downloads int `xorm:"'downloads' mediumint unsigned null default(0) comment('下载次数')"`
    Extra string `xorm:"'extra' varchar(255) null default(null) comment('额外信息')"`
    Deleted int `xorm:"'deleted' tinyint(1) null default(0) comment('删除标记')"`
    FileUrl string `xorm:"'file_url' text null comment('完整地址')"`
    FileType string `xorm:"'file_type' varchar(255) null default(null) comment('文件类型')"`
    DeletedTime int64 `xorm:"'deleted_time' bigint null default(null) comment('删除时间')"`
}

func (*File) TableName() string {
    return "ms_file"
}
type SourceLink struct {
    Id int `xorm:"'id' int autoincr pk notnull"`
    SourceType string `xorm:"'source_type' char(20) null default(null) comment('资源类型')"`
    SourceCode int64 `xorm:"'source_code' bigint null default(null) comment('资源编号')"`
    LinkType string  `xorm:"'link_type' char(20) null  default(null) comment('关联类型')"`
    LinkCode int64 `xorm:"'link_code' bigint null default(null) comment('关联编号')"`
    OrganizationCode int64 `xorm:"'organization_code' bigint null default(null) comment('组织编码')"`
    CreateBy int64 `xorm:"'create_by' bigint null default(null) comment('创建人')"`
    CreateTime int64 `xorm:"'create_time' bigint null default(null) comment('创建时间')"`
    Sort int `xorm:"'sort' int null default(0) comment('排序')"`
}

func (*SourceLink) TableName() string {
    return "ms_source_link"
}
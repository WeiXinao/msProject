package dao

import "context"

type ProjectDao interface {
	GetMenus(ctx context.Context) ([]*ProjectMenu, error)
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

package domain

type Project struct {
	Id                 int64
	Cover              string
	Name               string
	Description        string
	AccessControlType  int
	WhiteList          string
	Order              int
	Deleted            int
	TemplateCode       int
	Schedule           float64
	CreateTime         int64
	OrganizationCode   int64
	DeletedTime        string
	Private            int
	Prefix             string
	OpenPrefix         int
	Archive            int
	ArchiveTime        int64
	OpenBeginTime      int
	OpenTaskPrivate    int
	TaskBoardTheme     string
	BeginTime          int64
	EndTime            int64
	AutoUpdateSchedule int
}

func (m *Project) GetAccessControlType() string {
	return []string{"open", "private", "custom"}[m.AccessControlType]
}

const (
	Undeleted = iota
	Deleted
)

const (
	Unarchived = iota
	Archived
)

const (
	AccessControlTypeOpen = iota
	AccessControlTypePrivate
)

const (
	TaskBoardThemeDefault = "default"
	TaskBoardThemeSimple  = "simple"
)

package domain

import (
	"time"

	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/jinzhu/copier"
)

const (
	Wait = iota
	Doing
	Done
	Pause
	Cancel
	Closed
)

const (
	NoExecutor = iota
	IsExecutor
)

const (
	NoOwner = iota
	Owner
)

const (
	NoCanRead = iota
	CanRead
)

type Task struct {
	Id            int64
	ProjectCode   int64
	Name          string
	Pri           int
	ExecuteStatus int
	Description   string
	CreateBy      int64
	DoneBy        int64
	DoneTime      int64
	CreateTime    int64
	AssignTo      int64
	Deleted       int
	StageCode     int
	TaskTag       string
	Done          int
	BeginTime     int64
	EndTime       int64
	RemindTime    int64
	Pcode         int64
	Sort          int
	Like          int
	Star          int
	DeletedTime   int64
	Private       int
	IdNum         int
	Path          string
	Schedule      int
	VersionCode   int64
	FeaturesCode  int64
	WorkTime      int
	Status        int
}

type TaskMember struct {
	Id         int64
	TaskCode   int64
	IsExecutor int
	MemberCode int64
	JoinTime   int64
	IsOwner    int
}

func (t *Task) GetExecuteStatusStr() string {
	status := t.ExecuteStatus
	if status == Wait {
		return "wait"
	}
	if status == Doing {
		return "doing"
	}
	if status == Done {
		return "done"
	}
	if status == Pause {
		return "pause"
	}
	if status == Cancel {
		return "cancel"
	}
	if status == Closed {
		return "closed"
	}
	return ""
}

func (t *Task) ToTaskDisplay(encrypter encrypts.Encrypter) *TaskDisplay {
	td := &TaskDisplay{}
	copier.Copy(td, t)
	td.CreateTime = time.UnixMilli(t.CreateTime).Format(time.DateTime)
	td.DoneTime = time.UnixMilli(t.DoneTime).Format(time.DateTime)
	td.BeginTime = time.UnixMilli(t.BeginTime).Format(time.DateTime)
	td.EndTime = time.UnixMilli(t.EndTime).Format(time.DateTime)
	td.RemindTime = time.UnixMilli(t.RemindTime).Format(time.DateTime) 
	td.DeletedTime = time.UnixMilli(t.DeletedTime).Format(time.DateTime)
	td.CreateBy, _ = encrypter.EncryptInt64(t.CreateBy)
	td.ProjectCode, _ = encrypter.EncryptInt64(t.ProjectCode)
	td.DoneBy, _ = encrypter.EncryptInt64(t.DoneBy)
	td.AssignTo, _ = encrypter.EncryptInt64(t.AssignTo)
	td.StageCode, _ = encrypter.EncryptInt64(int64(t.StageCode))
	td.Pcode, _ = encrypter.EncryptInt64(t.Pcode)
	td.VersionCode, _ = encrypter.EncryptInt64(t.VersionCode)
	td.FeaturesCode, _ = encrypter.EncryptInt64(t.FeaturesCode)
	td.ExecuteStatus = t.GetExecuteStatusStr()
	td.Code, _ = encrypter.EncryptInt64(t.Id)
	td.CanRead = 1
	return td
}


type TaskDisplay struct {
	Id            int64
	ProjectCode   string
	Name          string
	Pri           int
	ExecuteStatus string
	Description   string
	CreateBy      string
	DoneBy        string
	DoneTime      string
	CreateTime    string
	AssignTo      string
	Deleted       int
	StageCode     string
	TaskTag       string
	Done          int
	BeginTime     string
	EndTime       string
	RemindTime    string
	Pcode         string
	Sort          int
	Like          int
	Star          int
	DeletedTime   string
	Private       int
	IdNum         int
	Path          string
	Schedule      int
	VersionCode   string
	FeaturesCode  string
	WorkTime      int
	Status        int
	Code          string
	CanRead       int
	Executor Executor
}

type Executor struct {
	Name string
	Avatar string
}
package domain

import (

	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/jinzhu/copier"
)

const (
	NotComment = iota
	Comment
)

type ProjectLog struct {
	Id           int64
	MemberCode   int64
	Content      string
	Remark       string
	Type         string
	CreateTime   int64
	SourceCode   int64
	ActionType   string
	ToMemberCode int64
	IsComment    int
	ProjectCode  int64
	Icon         string
	IsRobot      int
}

type ProjectLogDisplay struct {
	Id           int64
	MemberCode   string
	Content      string
	Remark       string
	Type         string
	CreateTime   string
	SourceCode   string
	ActionType   string
	ToMemberCode string
	IsComment    int
	ProjectCode  string
	Icon         string
	IsRobot      int
	Member       Member
}

func (l *ProjectLog) ToDisplay(encrypter encrypts.Encrypter) *ProjectLogDisplay {
	pd := &ProjectLogDisplay{}
	copier.Copy(pd, l)
	pd.MemberCode, _ = encrypter.EncryptInt64(l.MemberCode)
	pd.ToMemberCode, _ = encrypter.EncryptInt64(l.ToMemberCode)
	pd.ProjectCode, _ = encrypter.EncryptInt64(l.ProjectCode)
	pd.CreateTime = formatx.ToDateTimeString(l.CreateTime)
	pd.SourceCode, _ = encrypter.EncryptInt64(l.SourceCode)
	return pd
}

type IndexProjectLogDisplay struct {
	Content      string
	Remark       string
	CreateTime   string
	SourceCode   string
	IsComment    int
	ProjectCode  string
	MemberAvatar string
	MemberName   string
	ProjectName  string
	TaskName     string
}

func (l *ProjectLog) ToIndexDisplay(encrypter encrypts.Encrypter) *IndexProjectLogDisplay {
	pd := &IndexProjectLogDisplay{}
	copier.Copy(pd, l)
	pd.ProjectCode, _ = encrypter.EncryptInt64(l.ProjectCode)
	pd.CreateTime = formatx.ToDateTimeString(l.CreateTime)
	pd.SourceCode, _ = encrypter.EncryptInt64(l.SourceCode)
	return pd
}

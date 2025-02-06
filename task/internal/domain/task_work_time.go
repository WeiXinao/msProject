package domain

import (
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/jinzhu/copier"
)

type TaskWorkTime struct {
	Id         int64
	TaskCode   int64
	MemberCode int64
	CreateTime int64
	Content    string
	BeginTime  int64
	Num        int
}

type Member struct {
	Id int64
	Name string
	Avatar string
	Code string
}

type TaskWorkTimeDisplay struct {
	Id         int64
	TaskCode   string
	MemberCode string
	CreateTime string
	Content    string
	BeginTime  string
	Num        int
	Member     Member
}

func (t *TaskWorkTime) ToDisplay(encrypter encrypts.Encrypter) *TaskWorkTimeDisplay {
	td := &TaskWorkTimeDisplay{}
	copier.Copy(td, t)
	td.MemberCode, _ = encrypter.EncryptInt64(t.MemberCode)
	td.TaskCode, _ = encrypter.EncryptInt64(t.TaskCode)
	td.CreateTime = formatx.ToDateTimeString(t.CreateTime)
	td.BeginTime = formatx.ToDateTimeString(t.BeginTime)
	return td
}
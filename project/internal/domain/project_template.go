package domain

import (
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"time"
)

type ProjectTemplate struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       int64
	OrganizationCode int64
	Cover            string
	MemberCode       int64
	IsSystem         int
}

type ProjectTemplateAll struct {
	Id               int
	Name             string
	Description      string
	Sort             int
	CreateTime       string
	OrganizationCode string
	Cover            string
	MemberCode       string
	IsSystem         int
	TaskStages       []*TaskStagesOnlyName
	Code             string
}

func (pt ProjectTemplate) Convert(encrypter encrypts.Encrypter, taskStages []*TaskStagesOnlyName) *ProjectTemplateAll {
	organizationCode, _ := encrypter.EncryptInt64(pt.OrganizationCode)
	memberCode, _ := encrypter.EncryptInt64(pt.MemberCode)
	code, _ := encrypter.EncryptInt64(int64(pt.Id))
	pta := &ProjectTemplateAll{
		Id:               pt.Id,
		Name:             pt.Name,
		Description:      pt.Description,
		Sort:             pt.Sort,
		CreateTime:       time.UnixMilli(pt.CreateTime).Format(time.DateTime),
		OrganizationCode: organizationCode,
		Cover:            pt.Cover,
		MemberCode:       memberCode,
		IsSystem:         pt.IsSystem,
		TaskStages:       taskStages,
		Code:             code,
	}
	return pta
}

func ToProjectTemplateIds(pts []ProjectTemplate) []int {
	var ids []int
	for _, v := range pts {
		ids = append(ids, v.Id)
	}
	return ids
}

type MsTaskStagesTemplate struct {
	Id                  int
	Name                string
	ProjectTemplateCode int
	CreateTime          int64
	Sort                int
}

type TaskStagesOnlyName struct {
	Name string
}

// CovertProjectMap 模板 id -> 任务步骤
func CovertProjectMap(tsts []MsTaskStagesTemplate) map[int][]*TaskStagesOnlyName {
	var tss = make(map[int][]*TaskStagesOnlyName)
	for _, v := range tsts {
		ts := &TaskStagesOnlyName{}
		ts.Name = v.Name
		tss[v.ProjectTemplateCode] = append(tss[v.ProjectTemplateCode], ts)
	}
	return tss
}

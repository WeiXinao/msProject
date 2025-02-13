package domain

import (
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/jinzhu/copier"
)

type ProjectAuth struct {
	Id               int64  `json:"id"`
	OrganizationCode int64  `json:"organization_code"`
	Title            string `json:"title"`
	CreateAt         int64  `json:"create_at"`
	Sort             int    `json:"sort"`
	Status           int    `json:"status"`
	Desc             string `json:"desc"`
	CreateBy         int64  `json:"create_by"`
	IsDefault        int    `json:"is_default"`
	Type             string `json:"type"`
}

func (a *ProjectAuth) ToDisplay(encrypter encrypts.Encrypter) *ProjectAuthDisplay {
	p := &ProjectAuthDisplay{}
	copier.Copy(p, a)
	p.OrganizationCode, _ = encrypter.EncryptInt64(a.OrganizationCode)
	p.CreateAt = formatx.ToDateTimeString(a.CreateAt)
	if a.Type == "admin" || a.Type == "member" {
		p.CanDelete = 0
	} else {
		p.CanDelete = 1
	}
	return p
}

type ProjectAuthDisplay struct {
	Id               int64  `json:"id"`
	OrganizationCode string `json:"organization_code"`
	Title            string `json:"title"`
	CreateAt         string `json:"create_at"`
	Sort             int    `json:"sort"`
	Status           int    `json:"status"`
	Desc             string `json:"desc"`
	CreateBy         int64  `json:"create_by"`
	IsDefault        int    `json:"is_default"`
	Type             string `json:"type"`
	CanDelete        int    `json:"canDelete"`
}

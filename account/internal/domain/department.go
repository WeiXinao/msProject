package domain

import (
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/formatx"
	"github.com/jinzhu/copier"
)

type Department struct {
	Id               int64
	OrganizationCode int64
	Name             string
	Sort             int
	Pcode            int64
	icon             string
	CreateTime       int64
	Path             string
}

type DepartmentDisplay struct {
	Id               int64
	OrganizationCode string
	Name             string
	Sort             int
	Pcode            string
	icon             string
	CreateTime       string
	Path             string
}

func (d *Department) ToDisplay(encrypter encrypts.Encrypter) *DepartmentDisplay {
	dp := &DepartmentDisplay{}
	copier.Copy(dp, d)
	dp.CreateTime = formatx.ToDateTimeString(d.CreateTime)
	dp.OrganizationCode, _ = encrypter.EncryptInt64(d.OrganizationCode)
	if d.Pcode > 0 {
		dp.Pcode, _ = encrypter.EncryptInt64(d.Pcode)
	} else {
		dp.Pcode = ""
	}
	return dp
}

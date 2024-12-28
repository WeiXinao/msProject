package domain

import "time"

type Organization struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
	MemberId    int64
	CTime       time.Time
	Personal    int32
	Address     string
	Province    int32
	City        int32
	Area        int32
}

func (o *Organization) CreateTime(t int64) {
	o.CTime = time.UnixMilli(t)
}

const (
	StatusOrganizationUnknown = iota
	StatusOrganizationPersonal
)

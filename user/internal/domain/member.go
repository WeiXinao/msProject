package domain

import "time"

type Member struct {
	Id              int64
	Account         string
	Password        string
	Name            string
	Mobile          string
	Realname        string
	CreateTime      time.Time `copier:"-"`
	Status          int
	LastLoginTime   time.Time `copier:"-"`
	Sex             int
	Avatar          string
	Idcard          string
	Province        int
	City            int
	Area            int
	Address         string
	Description     string
	Email           string
	DingtalkOpenid  string
	DingtalkUnionid string
	DingtalkUserid  string
}

const (
	StatusMemberUnknown = iota
	StatusMemberNormal
)

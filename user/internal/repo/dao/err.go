package dao

import "errors"

var (
	ErrRecordNotFound = errors.New("记录不存在")
	ErrTypeConvert    = errors.New("类型转换错误")
)

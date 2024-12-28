package repo

import (
	"errors"
	"github.com/WeiXinao/msProject/user/internal/repo/dao"
)

var (
	ErrVerifyCaptchaFail = errors.New("校验验证码失败")
	ErrRecordNotFound    = dao.ErrRecordNotFound
)

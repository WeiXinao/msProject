package validatex

import (
	"regexp"
	"strings"
)

func VerifyMobile(mobile string) bool {
	if len(strings.TrimSpace(mobile)) == 0 {
		return false
	}

	mobilePattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	mobileExp := regexp.MustCompile(mobilePattern)
	return mobileExp.MatchString(mobile)
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)@\w+([-.]\w+).\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z].){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

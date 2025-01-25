package cache

import "fmt"

type UserCache struct {
}

func (*UserCache) RegisterCaptchaKey(mobile string) string {
	return fmt.Sprintf("REGISTER:%s", mobile)
}

func (*UserCache) MemberKey(memId int64) string {
	return fmt.Sprintf("MEMBER:%d", memId)
}

func (*UserCache) MemberOrganizationKey(memId int64) string {
	return fmt.Sprintf("MEMBER_ORGANIZATION:%d", memId)
}

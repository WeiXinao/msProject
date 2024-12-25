package cache

import "fmt"

type UserCache struct {
}

func (*UserCache) RegisterCaptchaKey(mobile string) string {
	return fmt.Sprintf("REGISTER:%s", mobile)
}

package encrypts

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(str string) string {
	hasher := md5.New()
	_, _ = io.WriteString(hasher, str)
	return hex.EncodeToString(hasher.Sum(nil))
}

package encrypts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

func Md5(str string) string {
	hasher := md5.New()
	_, _ = io.WriteString(hasher, str)
	return hex.EncodeToString(hasher.Sum(nil))
}

type encrypter struct {
	commonIV []byte
	keyText  string
}

// DecryptInt64 implements Encrypter.
func (e *encrypter) DecryptInt64(cipherStr string) (int64, error) {
	plainText, err := e.Decrypt(cipherStr)
	if err != nil {
		return 0, err
	}
	res, err :=  strconv.ParseInt(plainText, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func NewEncrypter(keyText string) Encrypter {
	return &encrypter{
		commonIV: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
		keyText:  keyText,
	}
}

func (e *encrypter) EncryptInt64(id int64) (cipherStr string, err error) {
	idStr := strconv.FormatInt(id, 10)
	return e.Encrypt(idStr)
}

func (e *encrypter) Encrypt(plainText string) (cipherStr string, err error) {
	// 转换成字节数据, 方便加密
	plainByte := []byte(plainText)
	keyByte := []byte(e.keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, e.commonIV)
	cipherByte := make([]byte, len(plainByte))
	cfb.XORKeyStream(cipherByte, plainByte)
	cipherStr = hex.EncodeToString(cipherByte)
	return
}

func (e *encrypter) Decrypt(cipherStr string) (plainText string, err error) {
	// 转换成字节数据, 方便加密
	keyByte := []byte(e.keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, e.commonIV)
	cipherByte, _ := hex.DecodeString(cipherStr)
	plainByte := make([]byte, len(cipherByte))
	cfbdec.XORKeyStream(plainByte, cipherByte)
	plainText = string(plainByte)
	return
}

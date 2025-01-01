package encrypts

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plain := "100123213123"
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	e := NewEncrypter(key)
	// 加密
	cipherByte, err := e.Encrypt(plain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", plain, cipherByte)
	// 解密
	plainText, err := e.Decrypt(cipherByte)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x ==> %s\n", cipherByte, plainText)
}

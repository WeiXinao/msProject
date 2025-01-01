package encrypts

type Encrypter interface {
	EncryptInt64(id int64) (cipherStr string, err error)
	Encrypt(plainText string) (cipherStr string, err error)
	Decrypt(cipherStr string) (plainText string, err error)
}

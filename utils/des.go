package utils

import (
	"crypto/des"
)

var des_key = []byte("hello123")

func DesEncoding(src []byte) ([]byte, error) {
	cipher, err := des.NewCipher(des_key)
	if err != nil {
		return nil, err
	}
	src = padPwd(src, cipher.BlockSize())
	dest := make([]byte, len(src))
	cipher.Encrypt(dest, src)
	return dest, nil
}

func DesDecoding(pwd []byte) ([]byte, error) {
	cipher, err := des.NewCipher(des_key)
	if err != nil {
		return nil, err
	}
	dest := make([]byte, len(pwd))
	cipher.Decrypt(dest, pwd)
	dest = unPadPwd(dest, cipher.BlockSize())
	return dest, nil
}

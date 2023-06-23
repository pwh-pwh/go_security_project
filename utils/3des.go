package utils

import (
	"crypto/des"
)

var des3_key = []byte("hello123hello123hello123")

func Des3Encoding(src []byte) ([]byte, error) {
	cipher, err := des.NewTripleDESCipher(des3_key)
	if err != nil {
		return nil, err
	}
	src = padPwd(src, cipher.BlockSize())
	dest := make([]byte, len(src))
	cipher.Encrypt(dest, src)
	return dest, nil
}

func Des3Decoding(pwd []byte) ([]byte, error) {
	cipher, err := des.NewTripleDESCipher(des3_key)
	if err != nil {
		return nil, err
	}
	dest := make([]byte, len(pwd))
	cipher.Decrypt(dest, pwd)
	dest = unPadPwd(dest, cipher.BlockSize())
	return dest, nil
}

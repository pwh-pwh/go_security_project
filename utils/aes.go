package utils

import (
	"bytes"
	"crypto/aes"
)

var key = []byte("hello1hello1hell")

func padPwd(src []byte, block_size int) []byte {
	pad_num := block_size - len(src)%block_size
	ret := bytes.Repeat([]byte{byte(pad_num)}, pad_num)
	src = append(src, ret...)
	return src
}

func AesEncoding(src []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padPwd(src, cipher.BlockSize())
	dest := make([]byte, len(src))
	cipher.Encrypt(dest, src)
	return dest, nil
}

func AesDecoding(pwd []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dest := make([]byte, len(pwd))
	cipher.Decrypt(dest, pwd)
	return unPadPwd(dest, aes.BlockSize), nil
}

func unPadPwd(pwd []byte, block_size int) []byte {
	pad_num := int(pwd[len(pwd)-1])
	return pwd[:len(pwd)-pad_num]
}

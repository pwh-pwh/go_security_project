package utils

import (
	"encoding/hex"
	"golang.org/x/crypto/md4"
)

func GenMd4(src []byte) string {
	md4Hash := md4.New()
	md4_byte := md4Hash.Sum(src)
	return hex.EncodeToString(md4_byte)
}

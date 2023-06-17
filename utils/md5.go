package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(src []byte) string {
	return hex.EncodeToString(md5.New().Sum(src))
}

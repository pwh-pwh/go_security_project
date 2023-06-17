package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenSha256(src []byte) string {
	sum := sha256.New().Sum(src)
	return hex.EncodeToString(sum)
}

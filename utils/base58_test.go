package utils

import (
	"fmt"
	"testing"
)

func TestBase58(t *testing.T) {
	src := []byte("he")
	encoding := Base58Encoding(src)
	fmt.Printf("encoding:%v\n", encoding)
	decoding := Base58Decoding(encoding)
	fmt.Printf("decoding:%s\n", decoding)
}

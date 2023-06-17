package utils

import "testing"

func TestGenSha256(t *testing.T) {
	println(GenSha256([]byte("hello")))
}

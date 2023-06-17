package utils

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	src := []byte("hello base64")
	encoding := Base64Encoding(src)
	decoding, err := Base64Decoding(encoding)
	if err != nil {

	} else {
		fmt.Printf("encoding:%v \ndecoding:%s", encoding, decoding)
	}
}

func TestBase64Url(t *testing.T) {
	src := []byte("127.0.0.1:8080/hello?name=aaa")
	encoding := Base64UrlEncoding(src)
	decoding, err := Base64UrlDecoding(encoding)
	if err != nil {

	} else {
		fmt.Printf("encoding:%v \ndecoding:%s", encoding, decoding)
	}
}

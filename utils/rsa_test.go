package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestSaveRsaKey(t *testing.T) {
	err := SaveRsaKey(1024)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRsaEncoding(t *testing.T) {
	encoding, err := RsaEncoding([]byte("hello"), "publicKey.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("encoding:%s", encoding)
	decoding, err := RsaDecoding(encoding, "privateKey.pem")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoding:%s", decoding)
}

func TestRsaSign(t *testing.T) {
	sign, err := RsaSign("privateKey.pem", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	err = RsaSignVerifySign("publicKey.pem", []byte("hello"), sign)
	if err != nil {
		fmt.Println("verify fail")
		return
	}
	fmt.Println("verify success")
}

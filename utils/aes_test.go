package utils

import (
	"fmt"
	"testing"
)

func TestAesEncoding(t *testing.T) {
	encoding, err := AesEncoding([]byte("hallenhallenhal"))
	if err != nil {
		fmt.Println(err)
	}
	decoding, err := AesDecoding(encoding)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", decoding)
	fmt.Printf("%s\n", Base64Encoding(encoding))
}

func TestAesDecoding(t *testing.T) {

}

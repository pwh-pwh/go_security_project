package utils

import (
	"fmt"
	"testing"
)

func TestDes3(t *testing.T) {
	encoding, err := Des3Encoding([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", encoding)
	decoding, err := Des3Decoding(encoding)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", decoding)
}

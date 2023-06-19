package utils

import (
	"fmt"
	"testing"
)

func TestDesEncoding(t *testing.T) {
	encoding, err := DesEncoding([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", encoding)
	decoding, err := DesDecoding(encoding)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", decoding)
}

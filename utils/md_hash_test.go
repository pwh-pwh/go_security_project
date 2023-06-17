package utils

import (
	"fmt"
	"testing"
)

func TestGenMd4(t *testing.T) {
	fmt.Printf("%v", GenMd4([]byte("hallen")))
}

func TestGenMd5(t *testing.T) {
	fmt.Printf("%v", GenMd5([]byte("hallen")))
}

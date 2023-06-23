package utils

import (
	"log"
	"testing"
)

func TestSaveRsaKey(t *testing.T) {
	err := SaveRsaKey(32)
	if err != nil {
		log.Fatal(err)
	}
}

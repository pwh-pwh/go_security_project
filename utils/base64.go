package utils

import "encoding/base64"

func Base64Encoding(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decoding(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

func Base64UrlEncoding(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func Base64UrlDecoding(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

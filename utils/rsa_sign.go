package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func RsaSign(path string, src []byte) ([]byte, error) {
	privateKey, err := GetPrivateKey(path)
	if err != nil {
		return nil, err
	}
	hash := sha256.New()
	hash.Write(src)
	sum := hash.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum)
}

//验证
func RsaSignVerifySign(path string, src, sign []byte) error {
	publicKey, err := GetPublicKey(path)
	if err != nil {
		return err
	}
	hash := sha256.New()
	hash.Write(src)
	sum := hash.Sum(nil)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, sum, sign)
}

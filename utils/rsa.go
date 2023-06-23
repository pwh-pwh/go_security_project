package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func SaveRsaKey(bit int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}
	publicKey := privateKey.PublicKey
	pkcs1PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	pkcs1PublicKey := x509.MarshalPKCS1PublicKey(&publicKey)
	block_private := pem.Block{Type: "private key", Bytes: pkcs1PrivateKey}
	block_public := pem.Block{Type: "public key", Bytes: pkcs1PublicKey}
	publicKeyFile, _ := os.Create("publicKey.pem")
	defer publicKeyFile.Close()
	privateKeyFile, _ := os.Create("privateKey.pem")
	defer privateKeyFile.Close()
	err = pem.Encode(publicKeyFile, &block_public)
	if err != nil {
		return err
	}
	err = pem.Encode(privateKeyFile, &block_private)
	if err != nil {
		return err
	}
	return nil
}

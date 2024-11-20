package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA keys:", err)
		return
	}
	publicKey := &privateKey.PublicKey
	fmt.Println("PublicKey", privateKey.PublicKey)

	message := []byte("Hello, RSA Cryptography")

	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		return
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return
	}
	fmt.Println("Decrypted:", string(decrypted))
}

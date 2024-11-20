package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating ECC keys:", err)
		return
	}
	fmt.Printf("Key: (pub: %s, priv: %s)\n", privateKey.PublicKey, privateKey)

	message := "Hello, ECC!"
	hash := sha256.Sum256([]byte(message))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}
	fmt.Printf("Signature: (r: %s, s: %s)\n", r.String(), s.String())

	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("Signature valid:", valid)
}

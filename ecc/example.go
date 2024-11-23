package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating ECC keys:", err)
		return
	}
	// Get the private key in bytes
	privateKeyBytes := privateKey.D.Bytes()

	// Encode the private key to a string (hexadecimal)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Printf("Private Key (Hex): %s\n", privateKeyHex)

	// Extract the public key
	publicKey := privateKey.PublicKey

	// Get the public key as bytes (uncompressed format)
	publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)

	// Encode the public key to a string (hexadecimal)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	fmt.Printf("Public Key (Hex): %s\n", publicKeyHex)

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

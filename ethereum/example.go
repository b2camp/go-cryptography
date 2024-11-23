package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate a private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}

	// Serialize the private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Printf("Private Key (Hex): %s\n", privateKeyHex)

	// Derive the public key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// Serialize the public key to uncompressed format (65 bytes)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)
	fmt.Printf("Public Key (Uncompressed Hex): %s\n", publicKeyHex)

	// Ethereum address is the last 20 bytes of the Keccak-256 hash of the public key
	address := crypto.PubkeyToAddress(*publicKey)
	fmt.Printf("Ethereum Address: %s\n", address.Hex())

	// Alternative: Using common.Address type
	addressBytes := address.Bytes()
	addressHex := hex.EncodeToString(addressBytes)
	fmt.Printf("Ethereum Address (Hex): %s\n", addressHex)
}

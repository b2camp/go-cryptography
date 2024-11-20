package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	message := "Hello"
	hash := sha256.Sum256([]byte(message))

	fmt.Printf("SHA-256 Hash: %x\n", hash)
}

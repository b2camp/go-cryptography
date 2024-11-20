package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("mysecretkey")
	message := "Hello, HMAC!"

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	hash := h.Sum(nil)

	fmt.Printf("HMAC: %s\n", hex.EncodeToString(hash))
}

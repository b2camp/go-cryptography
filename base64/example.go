package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello, Base64!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))
}

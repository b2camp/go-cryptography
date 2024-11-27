package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Define command-line arguments
	keystoreFile := flag.String("keystore", "", "Path to the keystore file")
	password := flag.String("password", "", "Password for the keystore")
	flag.Parse()

	// Validate input arguments
	if *keystoreFile == "" || *password == "" {
		log.Fatal("Both keystore file path and password must be provided")
	}

	// Load the keystore JSON file
	keystoreData, err := ioutil.ReadFile(*keystoreFile)
	if err != nil {
		log.Fatalf("Failed to read keystore file: %v", err)
	}

	// Decrypt the keystore using the provided password
	key, err := keystore.DecryptKey(keystoreData, *password)
	if err != nil {
		log.Fatalf("Failed to decrypt keystore: %v", err)
	}

	// Extract the private key from the decrypted key
	privateKey := key.PrivateKey

	// Display the private key
	fmt.Printf("Private Key: 0x%x\n", privateKey.D.Bytes())

	// Optionally, you can also display the address
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Printf("Address: %s\n", address.Hex())
}

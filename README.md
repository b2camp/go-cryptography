
# Cryptography Examples in Go

This repository contains examples demonstrating various cryptographic algorithms implemented in Go. 
The examples cover encryption, decryption, hashing, and digital signature operations using standard libraries and techniques.

## Table of Contents
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Algorithms and Examples](#algorithms-and-examples)
  - [1. AES Symmetric Encryption](#1-aes-symmetric-encryption)
  - [2. RSA Asymmetric Encryption](#2-rsa-asymmetric-encryption)
  - [3. Elliptic Curve Cryptography (ECC)](#3-elliptic-curve-cryptography-ecc)
  - [4. Hashing with SHA-256](#4-hashing-with-sha-256)
  - [5. HMAC](#5-hmac)
  - [6. Base64 Encoding/Decoding](#6-base64-encodingdecoding)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

Cryptography is the practice of securing communication and data through encoding techniques. 
This repository showcases various cryptographic algorithms and their practical use cases in Go, 
such as encryption, hashing, and signing.

---

## Prerequisites

To run these examples, you need:
- Go 1.17 or later installed on your system.
- [Golang.org crypto package](https://pkg.go.dev/golang.org/x/crypto) for specific algorithms like SHA-3.

---

## Algorithms and Examples

### 1. AES Symmetric Encryption

AES (Advanced Encryption Standard) is a fast and secure symmetric encryption algorithm.

- **Key Size**: 128-bit, 192-bit, or 256-bit.
- **Use Case**: Data encryption where both parties share the same key.

Example: `aes.go`
```bash
go run aes.go
```

---

### 2. RSA Asymmetric Encryption

RSA (Rivest–Shamir–Adleman) is a public-key cryptosystem for secure data transmission.

- **Key Size**: Commonly 2048-bit or 4096-bit.
- **Use Case**: Encrypt messages, secure key exchanges.

Example: `rsa.go`
```bash
go run rsa.go
```

---

### 3. Elliptic Curve Cryptography (ECC)

ECC provides efficient cryptography with smaller keys. It's widely used in blockchain and secure communications.

- **Use Case**: Generate keys, sign messages, verify signatures.

Example: `ecc.go`
```bash
go run ecc.go
```

---

### 4. Hashing with SHA-256

Generate a SHA-256 hash for message integrity verification.

- **Use Case**: Verify data integrity, store secure hashes of passwords.

Example: `sha256.go`
```bash
go run sha256.go
```

---

### 5. HMAC

HMAC (Hash-Based Message Authentication Code) provides message authentication using a shared secret key.

- **Use Case**: Verify the authenticity and integrity of messages.

Example: `hmac.go`
```bash
go run hmac.go
```

---

### 6. Base64 Encoding/Decoding

Base64 encoding is used to safely transmit binary data as text.

- **Use Case**: Encode binary data into text for storage or transmission.

Example: `base64.go`
```bash
go run base64.go
```

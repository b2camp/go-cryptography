
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
  - [7. PBKDF2 for Password Hashing](#7-pbkdf2-for-password-hashing)
  - [8. Digital Signature with RSA](#8-digital-signature-with-rsa)
  - [9. Keccak256 Hash](#9-keccak256-hash)

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

Example: `aes`
```bash
go run aes/example.go
```

---

### 2. RSA Asymmetric Encryption

RSA (Rivest–Shamir–Adleman) is a public-key cryptosystem for secure data transmission.

- **Key Size**: Commonly 2048-bit or 4096-bit.
- **Use Case**: Encrypt messages, secure key exchanges.

Example: `rsa`
```bash
go run rsa/example.go
```

---

### 3. Elliptic Curve Cryptography (ECC)

ECC provides efficient cryptography with smaller keys. It's widely used in blockchain and secure communications.

- **Use Case**: Generate keys, sign messages, verify signatures.

Example: `ecc`
```bash
go run ecc/example.go
```

---

### 4. Hashing with SHA-256

Generate a SHA-256 hash for message integrity verification.

- **Use Case**: Verify data integrity, store secure hashes of passwords.

Example: `sha256`
```bash
go run sha256/example.go
```

---

### 5. HMAC

HMAC (Hash-Based Message Authentication Code) provides message authentication using a shared secret key.

- **Use Case**: Verify the authenticity and integrity of messages.

Example: `hmac`
```bash
go run hmac/example.go
```

---

### 6. Base64 Encoding/Decoding

Base64 encoding is used to safely transmit binary data as text.

- **Use Case**: Encode binary data into text for storage or transmission.

Example: `base64`
```bash
go run base64/example.go
```

---

### 7. PBKDF2 for Password Hashing

PBKDF2 derives secure keys from passwords using a salt.

- **Use Case**: Securely store passwords.

Example: `pbkdf2`
```bash
go run pbkdf2/example.go
```

---

### 8. Digital Signature with RSA

Generate and verify digital signatures using RSA.

- **Use Case**: Authenticate the origin of messages.

Example: `rsa_signature`
```bash
go run rsa_signature/example.go
```

---

### 9. Keccak256 Hash

Keccak256 is used for hashing in Ethereum and other blockchain platforms.

- **Use Case**: Generate hashes for Ethereum addresses, function signatures, etc.

Example: `keccak256`
```bash
go run keccak256/example.go
```

---

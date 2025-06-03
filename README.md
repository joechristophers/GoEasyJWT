
# GoEasyJWT 🔐

A super lightweight and easy-to-use JWT utility package for Go.

## Features

- 🔑 Token generation with custom claims
- ⏳ Token expiration support
- ✅ Token verification and claim extraction
- 📦 Simple, reusable functions for secure apps

---

## Installation

```bash
go get github.com/joechristophers/GoEasyJWT

Usage
Import

import "github.com/joechristophers/GoEasyJWT"

Generate a Token

claims := map[string]interface{}{
    "id": "user123",
    "name": "Joe",
}

key := []byte("your-secret-key")
token, err := GoEasyJWT.GenerateToken(claims, key, 4) // 4-hour expiry
if err != nil {
    log.Fatal("Token generation failed:", err)
}
fmt.Println("JWT:", token)

Verify a Token

parsedClaims, err := GoEasyJWT.VerifyToken(token, key)
if err != nil {
    log.Println("Invalid token:", err)
} else {
    fmt.Println("Claims:", parsedClaims)
}

Example

Check test_jwt.go for a working example.
License

MIT © 2025 Joseph Christopher


---

Let me know if you want to add badges (GoDoc, pkg.go.dev, License, etc.) or instructions for contributing and testing.

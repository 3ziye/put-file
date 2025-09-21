package main

import (
	"fmt"
	"log"

	"github.com/3ziye/put-file/internal/auth"
)

func main() {
	// Test password hashing and verification functionality
	password := "admin123"

	// Hash password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		log.Fatalf("Password hashing failed: %v", err)
	}

	fmt.Printf("Original password: %s\n", password)
	fmt.Printf("Hashed password: %s\n", hashedPassword)

	// Verify correct password
	isValid := auth.CheckPasswordHash(password, hashedPassword)
	fmt.Printf("Correct password verification result: %v\n", isValid)

	// Verify wrong password
	isValid = auth.CheckPasswordHash("wrongpassword", hashedPassword)
	fmt.Printf("Wrong password verification result: %v\n", isValid)

	// Demonstrate Basic auth header creation and parsing
	basicAuth := "Basic " + "YWRtaW46YWRtaW4xMjM=" // Base64 encoding of admin:admin123
	fmt.Printf("Basic auth header: %s\n", basicAuth)
}
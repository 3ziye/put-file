package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Simulate username and password read from environment variables when server starts
// These values should be the same as those actually used by the server
username := "admin"
password := "admin123"

fmt.Println("=== Debug Credentials Info ===")
fmt.Println("Original username:", username)
fmt.Println("Original password:", password)
	
	// Check if password is already bcrypt hashed
isBcrypt := strings.HasPrefix(password, "$2a$") || strings.HasPrefix(password, "$2b$") || strings.HasPrefix(password, "$2y$")
fmt.Println("Is password already bcrypt hashed:", isBcrypt)
	
	// If not hashed, generate hash
var hash string
var err error
if !isBcrypt {
	fmt.Println("Generating password hash...")
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to generate hash:", err)
		os.Exit(1)
	}
	hash = string(hashBytes)
} else {
	hash = password
}

fmt.Println("Final password hash used:", hash)
	
	// Verify if original password matches hash
err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
fmt.Println("Password verification result:", err == nil)
if err != nil {
	fmt.Println("Verification failed reason:", err)
}
	
	// Output commands for testing
fmt.Println("\n=== Test Commands ===")
fmt.Println("Start server with explicit username and password:")
fmt.Printf("USERNAME=\"%s\" PASSWORD=\"%s\" go run cmd/server/main.go\n", username, password)
}
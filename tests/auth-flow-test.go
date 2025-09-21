package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/3ziye/put-file/internal/auth"
	"github.com/3ziye/put-file/internal/models"
)

func main() {
	// Create test user
	username := "admin"
	password := "admin123"

	// Simulate server-side password hashing
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		log.Fatalf("Password hashing failed: %v", err)
	}

	// 创建有效的用户映射
	validUsers := map[string]models.User{
		username: {Username: username, Password: hashedPassword},
	}

	fmt.Printf("=== Authentication Flow Test ===\n")
	fmt.Printf("Original password: %s\n", password)
	fmt.Printf("Hashed password: %s\n", hashedPassword)

	// Simulate frontend creating Basic auth header
	creds := username + ":" + password
	encodedCreds := base64.StdEncoding.EncodeToString([]byte(creds))
	authHeader := "Basic " + encodedCreds
	fmt.Printf("Frontend generated auth header: %s\n", authHeader)

	// Simulate backend parsing Basic auth header
	split := strings.SplitN(authHeader, " ", 2)
	if len(split) != 2 || split[0] != "Basic" {
		log.Fatalf("Invalid auth header format")
	}

	decoded, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		log.Fatalf("Base64 decoding failed: %v", err)
	}

	userPass := strings.SplitN(string(decoded), ":", 2)
	if len(userPass) != 2 {
		log.Fatalf("Invalid username:password format")
	}

	decodedUsername, decodedPassword := userPass[0], userPass[1]
	fmt.Printf("Backend parsed username: %s\n", decodedUsername)
	fmt.Printf("Backend parsed password: %s\n", decodedPassword)

	// Verify credentials
	isValid := auth.ValidateCredentials(decodedUsername, decodedPassword, validUsers)
	fmt.Printf("Credentials verification result: %v\n\n", isValid)

	// Additional test: Simulate http.Request.BasicAuth() method
	fmt.Printf("=== Simulate http.Request.BasicAuth() Test ===\n")

	// Create a test request
	req, err := http.NewRequest("GET", "/api/files", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set auth header
	req.Header.Set("Authorization", authHeader)

	// Using Request.BasicAuth() method
	reqUsername, reqPassword, ok := req.BasicAuth()
	fmt.Printf("BasicAuth() parsing successful: %v\n", ok)
	fmt.Printf("BasicAuth() parsed username: %s\n", reqUsername)
	fmt.Printf("BasicAuth() parsed password: %s\n", reqPassword)

	// Verify using parsed credentials
	isValid = auth.ValidateCredentials(reqUsername, reqPassword, validUsers)
	fmt.Printf("Verification result using BasicAuth() parsed credentials: %v\n", isValid)
}
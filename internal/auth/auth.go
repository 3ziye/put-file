// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package auth

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/3ziye/GoStaticServe/internal/logs"
	"github.com/3ziye/GoStaticServe/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthMiddleware Authentication middleware
type AuthMiddleware struct {
	validUsers map[string]models.User
}

// NewAuthMiddleware Create a new authentication middleware
func NewAuthMiddleware(validUsers map[string]models.User) *AuthMiddleware {
	return &AuthMiddleware{
		validUsers: validUsers,
	}
}

// MiddlewareFunc Return HTTP middleware function
func (am *AuthMiddleware) MiddlewareFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if it's an API path that doesn't require authentication
		if strings.HasPrefix(r.URL.Path, "/api/login") {
			next(w, r)
			return
		}

		// API paths require authentication
		if strings.HasPrefix(r.URL.Path, "/api/") {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.Header().Set("WWW-Authenticate", "Basic realm=\"File Server API\"")
				logs.Error("API authentication failed: missing Authorization header")
				http.Error(w, "Unauthorized access", http.StatusUnauthorized)
				return
			}

			// Parse Basic Auth
			const prefix = "Basic "
			if !strings.HasPrefix(authHeader, prefix) {
				logs.Error("API authentication failed: invalid authorization format")
				http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
				return
			}

			// Validate user credentials
			username, password, ok := r.BasicAuth()
			if !ok {
				logs.Error("API authentication failed: invalid authorization credentials")
				http.Error(w, "Invalid authorization credentials", http.StatusUnauthorized)
				return
			}

			// Verify username and password
			if user, exists := am.validUsers[username]; exists {
				if subtle.ConstantTimeCompare([]byte(password), []byte(user.Password)) == 1 {
					// Authentication successful, continue processing the request
					next(w, r)
					return
				}
			}

			// Authentication failed
			logs.Error("API authentication failed: invalid username or password for user %s", username)
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// Non-API paths pass through directly
		next(w, r)
	}
}

// HashPassword 对密码进行哈希处理
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 验证密码与哈希值是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidateCredentials Check if user credentials are valid
func ValidateCredentials(username, password string, validUsers map[string]models.User) bool {
	if user, exists := validUsers[username]; exists {
		return CheckPasswordHash(password, user.Password)
	}
	return false
}
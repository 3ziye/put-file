// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/3ziye/put-file/internal/auth"
	"github.com/3ziye/put-file/internal/handlers"
	"github.com/3ziye/put-file/internal/models"
)

// Test login functionality
func TestHandleLogin(t *testing.T) {
	// Set up test users
	validUsers := map[string]models.User{
		"admin": {Username: "admin", Password: "admin123"},
	}

	handler := handlers.HandleLogin(validUsers)

	// Test successful login
	t.Run("SuccessfulLogin", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"username":"admin","password":"admin123"}`)
		req, _ := http.NewRequest("POST", "/api/login", reqBody)
		req.Header.Set("Content-Type", "application/json")
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Verify response content
		var response models.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !response.Success {
			t.Error("Expected login to succeed")
		}
	})

	// Test non-existent username
	t.Run("NonExistentUsername", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"username":"notexist","password":"password"}`)
		req, _ := http.NewRequest("POST", "/api/login", reqBody)
		req.Header.Set("Content-Type", "application/json")
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
		}

		var response models.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Success {
			t.Error("Expected login to fail")
		}
	})

	// Test incorrect password
	t.Run("IncorrectPassword", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"username":"admin","password":"wrong"}`)
		req, _ := http.NewRequest("POST", "/api/login", reqBody)
		req.Header.Set("Content-Type", "application/json")
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
		}
	})

	// Test unsupported HTTP method
	t.Run("UnsupportedHTTPMethod", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/login", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}

// Test file upload functionality
func TestHandleAPIUpload(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "test-upload")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	handler := handlers.HandleAPIUpload(tempDir)

	// Test successful upload
	t.Run("SuccessfulUpload", func(t *testing.T) {
		// Create test file
		testFile := bytes.NewBufferString("This is test file content")

		// Create multipart form
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.txt")
		io.Copy(part, testFile)
		writer.Close()

		req, _ := http.NewRequest("POST", "/api/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Verify file exists
		uploadedFilePath := filepath.Join(tempDir, "test.txt")
		if _, err := os.Stat(uploadedFilePath); os.IsNotExist(err) {
			t.Errorf("Uploaded file does not exist: %s", uploadedFilePath)
		}
	})

	// Test unsupported HTTP method
	t.Run("UnsupportedHTTPMethod", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/upload", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}

// TestHandleAPIListFiles tests the file list retrieval functionality
func TestHandleAPIListFiles(t *testing.T) {
	// Create temporary directory and test files
	tempDir, err := os.MkdirTemp("", "test-list")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	os.WriteFile(filepath.Join(tempDir, "test1.txt"), []byte("Test content 1"), 0644)
	os.WriteFile(filepath.Join(tempDir, "test2.txt"), []byte("Test content 2"), 0644)

	handler := handlers.HandleAPIListFiles(tempDir)

	// Test successful file list retrieval
	t.Run("SuccessfulFileListRetrieval", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/files", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Verify response content
		var response models.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !response.Success {
			t.Error("Expected file list retrieval to succeed")
		}

		// Verify file list contains test files
		fileList, ok := response.Data.([]interface{})
		if !ok {
			t.Fatal("Response data is not a file list")
		}

		if len(fileList) < 2 {
			t.Errorf("Expected at least 2 files, got %d", len(fileList))
		}
	})

	// Test unsupported HTTP method
	t.Run("UnsupportedHTTPMethod", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/api/files", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}

// TestHandleAPIDeleteFile tests the file deletion functionality
func TestHandleAPIDeleteFile(t *testing.T) {
	// Create temporary directory and test file
	tempDir, err := os.MkdirTemp("", "test-delete")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test file
	testFilePath := filepath.Join(tempDir, "test.txt")
	os.WriteFile(testFilePath, []byte("Test content"), 0644)

	handler := handlers.HandleAPIDeleteFile(tempDir)

	// Test successful file deletion
	t.Run("SuccessfulFileDeletion", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/files/test.txt", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Verify file was deleted
		if _, err := os.Stat(testFilePath); !os.IsNotExist(err) {
			t.Error("File should be deleted but still exists")
		}
	})

	// Test deletion of non-existent file
	t.Run("DeleteNonExistentFile", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/files/nonexistent.txt", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	// Test unsupported HTTP method
	t.Run("UnsupportedHTTPMethod", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/files/test.txt", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		resp := rw.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}
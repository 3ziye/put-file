// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/3ziye/GoStaticServe/internal/config"
	"github.com/3ziye/GoStaticServe/internal/models"
)

// End-to-end test that simulates a complete user interaction flow
func TestE2E(t *testing.T) {
	// Skip end-to-end tests in short mode
	if testing.Short() {
		t.Skip("Skipping end-to-end test")
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "test-e2e")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create temporary configuration file
	config := config.Config{
		ServerPort: "8081",
		RootDir:    tempDir,
		Username:   "admin",
		Password:   "admin123",
	}

	configData, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to serialize configuration: %v", err)
	}

	configPath := filepath.Join(tempDir, "config.json")
	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		t.Fatalf("Failed to write configuration file: %v", err)
	}

	// Start server
	serverURL := "http://localhost:8081"
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// Start server with custom port and configuration
		args := []string{
			"--port=8081",
			"--root=" + tempDir,
			"--username=admin",
			"--password=admin123",
			"--log-level=ERROR", // Reduce log output
		}

		// Capture stdout and stderr to avoid test output clutter
		oldStdout := os.Stdout
		oldStderr := os.Stderr
		nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
		os.Stdout = nullFile
		os.Stderr = nullFile
		defer func() {
			os.Stdout = oldStdout
			os.Stderr = oldStderr
			nullFile.Close()
		}()

		// Start server (this will block until server stops)
		// Since we need to stop the server before test ends, we start it in a goroutine
		// In actual tests, we'll verify server functionality through HTTP requests
		go func() {
			// Main function will block, so we need a way to stop it
			// In this simplified test, we let the server run for a while then exit automatically
			time.Sleep(10 * time.Second)
		}()
	}()

	// Wait for server to start
	t.Log("Waiting for server to start...")
	time.Sleep(2 * time.Second)

	// Test step 1: API login authentication
	t.Run("APILoginAuthentication", func(t *testing.T) {
		t.Log("Testing API login...")
		reqBody := bytes.NewBufferString(`{"username":"admin","password":"admin123"}`)
		req, err := http.NewRequest("POST", serverURL+"/api/login", reqBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
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

	// Test step 2: API file upload
	t.Run("APIFileUpload", func(t *testing.T) {
		t.Log("Testing API file upload...")
		// Create test file
		testFile := bytes.NewBufferString("This is end-to-end test file content")
		testFileName := "e2e-test.txt"

		// Create multipart form
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", testFileName)
		io.Copy(part, testFile)
		writer.Close()

		req, err := http.NewRequest("POST", serverURL+"/api/upload", body)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.SetBasicAuth("admin", "admin123")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
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
			t.Error("Expected file upload to succeed")
		}

		// Verify file exists on server
		uploadedFilePath := filepath.Join(tempDir, testFileName)
		if _, err := os.Stat(uploadedFilePath); os.IsNotExist(err) {
			t.Errorf("Uploaded file does not exist: %s", uploadedFilePath)
		}
	})

	// Test step 3: API get file list
	t.Run("APIGetFileList", func(t *testing.T) {
		t.Log("Testing API file list retrieval...")
		req, err := http.NewRequest("GET", serverURL+"/api/files", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.SetBasicAuth("admin", "admin123")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
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

		// Verify file list contains our uploaded file
		fileList, ok := response.Data.([]interface{})
		if !ok {
			t.Fatal("Response data is not a file list")
		}

		found := false
		for _, file := range fileList {
			fileMap, ok := file.(map[string]interface{})
			if !ok {
				continue
			}
			if fileName, ok := fileMap["name"].(string); ok && fileName == "e2e-test.txt" {
				found = true
				break
			}
		}

		if !found {
			t.Error("Uploaded test file not found in file list")
		}
	})

	// Test step 4: File download
	t.Run("FileDownload", func(t *testing.T) {
		t.Log("Testing file download...")
		req, err := http.NewRequest("GET", serverURL+"/files/e2e-test.txt", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Verify file content
		content, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read file content: %v", err)
		}

		expectedContent := "This is end-to-end test file content"
		if string(content) != expectedContent {
			t.Errorf("File content mismatch, expected '%s', got '%s'", expectedContent, string(content))
		}
	})

	// Test step 5: API delete file
	t.Run("APIDeleteFile", func(t *testing.T) {
		t.Log("Testing API file deletion...")
		req, err := http.NewRequest("DELETE", serverURL+"/api/files/e2e-test.txt", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.SetBasicAuth("admin", "admin123")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
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
			t.Error("Expected file deletion to succeed")
		}

		// Verify file is deleted
		filePath := filepath.Join(tempDir, "e2e-test.txt")
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			t.Error("File should be deleted but still exists")
		}
	})

	// Test step 6: Verify file is deleted
	t.Run("VerifyFileIsDeleted", func(t *testing.T) {
		t.Log("Verifying file is deleted...")
		req, err := http.NewRequest("GET", serverURL+"/files/e2e-test.txt", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Expect 404 Not Found
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	// Wait for server to shutdown
	t.Log("Waiting for server to shutdown...")
	time.Sleep(1 * time.Second)
	wg.Wait()
}
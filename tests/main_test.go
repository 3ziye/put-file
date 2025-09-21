// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/3ziye/put-file/internal/handlers"
)

// This is a simple test file used to verify the project structure
func TestProjectStructure(t *testing.T) {
	// Check if the handlers package can be imported
	handlerFunc := handlers.HandleUpload("./uploads")
	if handlerFunc == nil {
		t.Fatal("HandleUpload function returns nil")
	}
}

// Simple API response test
func TestSendAPIResponse(t *testing.T) {
	rw := httptest.NewRecorder()
	handlers.SendAPIResponse(rw, true, "Test successful", nil, http.StatusOK)

	resp := rw.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type to contain 'application/json', got '%s'", contentType)
	}
}
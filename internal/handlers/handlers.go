// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/3ziye/GoStaticServe/internal/logs"
	"github.com/3ziye/GoStaticServe/internal/models"
)

// SendAPIResponse Send API response
func SendAPIResponse(w http.ResponseWriter, success bool, message string, data interface{}, statusCode int) {
	response := models.APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	if !success {
		response.Error = message
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logs.Error("Failed to encode response JSON: %v", err)
	}
}
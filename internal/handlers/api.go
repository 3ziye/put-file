// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/3ziye/put-file/internal/auth"
	"github.com/3ziye/put-file/internal/logs"
	"github.com/3ziye/put-file/internal/models"
)

// HandleAuthVerify handles authentication verification requests
func HandleAuthVerify(validUsers map[string]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			SendAPIResponse(w, false, "Only GET requests are supported", nil, http.StatusMethodNotAllowed)
			return
		}

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			SendAPIResponse(w, false, "Authorization header is required", nil, http.StatusUnauthorized)
			return
		}

		// Parse Basic Auth
		const prefix = "Basic "
		if !strings.HasPrefix(authHeader, prefix) {
			SendAPIResponse(w, false, "Invalid authorization format", nil, http.StatusUnauthorized)
			return
		}

		// Validate user credentials
		username, password, ok := r.BasicAuth()
		if !ok {
			SendAPIResponse(w, false, "Invalid authorization credentials", nil, http.StatusUnauthorized)
			return
		}

		// Verify username and password
		if auth.ValidateCredentials(username, password, validUsers) {
			SendAPIResponse(w, true, "Authentication successful", map[string]string{"username": username}, http.StatusOK)
			return
		}

		SendAPIResponse(w, false, "Invalid username or password", nil, http.StatusUnauthorized)
	}
}

// HandleLogin Handle login requests
func HandleLogin(validUsers map[string]models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendAPIResponse(w, false, "Only POST requests are supported", nil, http.StatusMethodNotAllowed)
			return
		}

		// Parse JSON request body
		var credentials models.User
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			logs.Error("Failed to parse request body: %v", err)
			SendAPIResponse(w, false, "Failed to parse request body", nil, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Validate user credentials
		if auth.ValidateCredentials(credentials.Username, credentials.Password, validUsers) {
			SendAPIResponse(w, true, "Login successful", map[string]string{"username": credentials.Username}, http.StatusOK)
			return
		}

		SendAPIResponse(w, false, "Invalid username or password", nil, http.StatusUnauthorized)
	}
}

// HandleAPIUpload Handle API file uploads
func HandleAPIUpload(rootDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendAPIResponse(w, false, "Only POST requests are supported", nil, http.StatusMethodNotAllowed)
			return
		}

		// Parse form with max memory of 10MB
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			logs.Error("Failed to parse form: %v", err)
			SendAPIResponse(w, false, "Failed to parse form: "+err.Error(), nil, http.StatusBadRequest)
			return
		}

		// Get file
		file, handler, err := r.FormFile("file")
		if err != nil {
			logs.Error("Failed to get file: %v", err)
			SendAPIResponse(w, false, "Failed to get file: "+err.Error(), nil, http.StatusBadRequest)
			return
		}
		defer file.Close()

		// 检查文件类型
		if !isAllowedFileType(handler.Filename) {
			logs.Error("File type not allowed: %s", handler.Filename)
			SendAPIResponse(w, false, "File type not allowed", nil, http.StatusBadRequest)
			return
		}

		// 防止路径遍历攻击，只使用文件名部分
		safeFilename := filepath.Base(handler.Filename)
		// Construct file save path
		savePath := filepath.Join(rootDir, safeFilename)

		// Create target file
		dst, err := os.Create(savePath)
		if err != nil {
			logs.Error("Failed to create file: %v", err)
			SendAPIResponse(w, false, "Failed to create file: "+err.Error(), nil, http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy file content
		size, err := io.Copy(dst, file)
		if err != nil {
			logs.Error("Failed to save file: %v", err)
			SendAPIResponse(w, false, "Failed to save file: "+err.Error(), nil, http.StatusInternalServerError)
			return
		}

		// Get file info
		fileInfo, err := os.Stat(savePath)
		if err != nil {
			logs.Error("Failed to get file info: %v", err)
			SendAPIResponse(w, false, "Failed to get file info: "+err.Error(), nil, http.StatusInternalServerError)
			return
		}

		// Build file info response
		fileData := models.FileInfo{
			Name:      safeFilename,
			Size:      size,
			Path:      "/files/" + safeFilename,
			CreatedAt: fileInfo.ModTime(),
		}

		SendAPIResponse(w, true, "File uploaded successfully", fileData, http.StatusOK)
	}
}

// HandleAPIListFiles handles API requests to get file list
func HandleAPIListFiles(rootDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			SendAPIResponse(w, false, "Only GET requests are supported", nil, http.StatusMethodNotAllowed)
			return
		}

		// Read files in directory
	files, err := os.ReadDir(rootDir)
	if err != nil {
		logs.Error("Failed to read directory: %v", err)
		SendAPIResponse(w, false, "Failed to read directory: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

		// Build file list
		var fileList []models.FileInfo
		for _, file := range files {
			if !file.IsDir() {
				fileInfo, err := file.Info()
				if err != nil {
					continue
				}

				fileList = append(fileList, models.FileInfo{
					Name:      fileInfo.Name(),
					Size:      fileInfo.Size(),
					Path:      "/files/" + fileInfo.Name(),
					CreatedAt: fileInfo.ModTime(),
				})
			}
		}

		SendAPIResponse(w, true, "File list fetched successfully", fileList, http.StatusOK)
	}
}

// HandleAPIDeleteFile handles API requests to delete files
func HandleAPIDeleteFile(rootDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			SendAPIResponse(w, false, "Only DELETE requests are supported", nil, http.StatusMethodNotAllowed)
			return
		}

		// Get file name
		fileName := strings.TrimPrefix(r.URL.Path, "/api/files/")
		if fileName == "" {
			SendAPIResponse(w, false, "File name cannot be empty", nil, http.StatusBadRequest)
			return
		}

		// 防止路径遍历攻击，只使用文件名部分
		safeFilename := filepath.Base(fileName)
		// Build file path
		filePath := filepath.Join(rootDir, safeFilename)

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			logs.Error("File does not exist: %s", filePath)
			SendAPIResponse(w, false, "File does not exist", nil, http.StatusNotFound)
			return
		}

		// Delete file
		if err := os.Remove(filePath); err != nil {
			logs.Error("Failed to delete file: %v", err)
			SendAPIResponse(w, false, "Failed to delete file: "+err.Error(), nil, http.StatusInternalServerError)
			return
		}

		SendAPIResponse(w, true, "File deleted successfully", map[string]string{"filename": safeFilename}, http.StatusOK)
	}
}
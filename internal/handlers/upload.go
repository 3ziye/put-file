// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/3ziye/put-file/internal/logs"
)

// 检查文件扩展名是否允许
func isAllowedFileType(filename string) bool {
	allowedExtensions := []string{"pdf", "doc", "docx", "jpg", "jpeg", "png", "gif", "txt", "md", "html", "css", "js", "json", "xml", "zip", "rar", "tar", "gz"}
	extension := strings.ToLower(filepath.Ext(filename))
	if extension == "" {
		return false
	}
	// 移除扩展名前的点
	extension = extension[1:]
	for _, allowed := range allowedExtensions {
		if extension == allowed {
			return true
		}
	}
	return false
}

// HandleUpload handles file upload requests
func HandleUpload(rootDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
			return
		}

		// Parse form with max memory of 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		logs.Error("Failed to parse form: %v", err)
		http.Error(w, fmt.Sprintf("Failed to parse form: %v", err), http.StatusBadRequest)
		return
	}

		// Get file
	file, handler, err := r.FormFile("file")
	if err != nil {
		logs.Error("Failed to get file: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get file: %v", err), http.StatusBadRequest)
		return
	}
		defer file.Close()

		// 检查文件类型
		if !isAllowedFileType(handler.Filename) {
			logs.Error("File type not allowed: %s", handler.Filename)
			http.Error(w, "File type not allowed", http.StatusBadRequest)
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
		http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
		return
	}
		defer dst.Close()

		// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		logs.Error("Failed to save file: %v", err)
		http.Error(w, fmt.Sprintf("Failed to save file: %v", err), http.StatusInternalServerError)
		return
	}

		// 重定向回首页并带上成功状态参数
		w.Header().Set("Location", fmt.Sprintf("/?status=success&message=文件上传成功：%s", safeFilename))
		w.WriteHeader(http.StatusFound)
	}
}
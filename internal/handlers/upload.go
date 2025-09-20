// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/3ziye/GoStaticServe/internal/logs"
	"github.com/3ziye/GoStaticServe/internal/utils"
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

		// Display success message
		html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Upload Successful</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					max-width: 800px;
					margin: 0 auto;
					padding: 20px;
				}
				.container {
					background-color: #f9f9f9;
					padding: 20px;
					border-radius: 8px;
					box-shadow: 0 0 10px rgba(0,0,0,0.1);
				}
				.success {
					background-color: #d4edda;
					color: #155724;
					padding: 10px;
					border-radius: 4px;
				}
				.back {
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>File Upload Successful</h1>
				<div class="success">
					<p>File Name: %s</p>
					<p>File Size: %d bytes</p>
					<p>Access URL: <a href="/files/%s">/files/%s</a></p>
				</div>
				<div class="back">
					<a href="/">Back to Upload Page</a>
				</div>
			</div>
		</body>
		</html>
		`, safeFilename, handler.Size, safeFilename, safeFilename)

		fmt.Fprint(w, html)
	}
}
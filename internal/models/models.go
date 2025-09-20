// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package models

import "time"

// User struct for storing authentication information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// APIResponse struct for API responses
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// FileInfo struct for file information
type FileInfo struct {
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

// TemplateFileInfo struct for template rendering
type TemplateFileInfo struct {
	Name            string
	IsDir           bool
	Size            int64
	SizeFormatted   string
	ModTime         time.Time
	ModTimeFormatted string
}
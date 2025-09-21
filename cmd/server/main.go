// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/3ziye/put-file/internal/auth"
	"github.com/3ziye/put-file/internal/config"
	"github.com/3ziye/put-file/internal/handlers"
	"github.com/3ziye/put-file/internal/logs"
	"github.com/3ziye/put-file/internal/models"
	"github.com/3ziye/put-file/internal/utils"
)

var version = "dev"

func main() {
	// Define command line arguments
	configPath := flag.String("config", "config.json", "Configuration file path")
	rootDir := flag.String("root", "./uploads", "File server root directory")
	port := flag.String("port", "8080", "Server listening port")
	username := flag.String("username", "admin", "Default admin username")
	password := flag.String("password", "admin123", "Default admin password")
	logLevel := flag.String("log-level", "INFO", "Log level: DEBUG, INFO, WARN, ERROR, FATAL")
	logFile := flag.String("log-file", "", "Log file path, defaults to standard output")
	flag.Parse()

	// Initialize logging system
	logs.InitLogger(&logs.LoggerConfig{
		Level:  *logLevel,
		Output: *logFile,
		Prefix: "file-server",
	})

	// Default configuration
	serverConfig := config.GetDefaultConfig()

	// 1. Try to load configuration from file
	if _, err := os.Stat(*configPath); err == nil {
		loadedConfig, err := config.LoadConfig(*configPath)
		if err != nil {
			logs.Warn("Warning: Failed to load configuration from file, using default configuration: %v", err)
		} else {
			serverConfig = loadedConfig
		}
	} else {
		logs.Info("Configuration file %s does not exist, using default configuration", *configPath)
	}

	// 2. Apply environment variables (higher priority than configuration file)
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		serverConfig.ServerPort = envPort
	}
	if envRoot := os.Getenv("SERVER_ROOT"); envRoot != "" {
		serverConfig.RootDir = envRoot
	}
	if envUsername := os.Getenv("SERVER_USERNAME"); envUsername != "" {
		serverConfig.Username = envUsername
	}
	if envPassword := os.Getenv("SERVER_PASSWORD"); envPassword != "" {
		serverConfig.Password = envPassword
	}
	if envLogLevel := os.Getenv("SERVER_LOG_LEVEL"); envLogLevel != "" {
		serverConfig.LogLevel = envLogLevel
	}
	if envLogFile := os.Getenv("SERVER_LOG_FILE"); envLogFile != "" {
		serverConfig.LogFile = envLogFile
	}

	// 3. Apply command line arguments (higher priority than environment variables)
	if *port != "8080" {
		serverConfig.ServerPort = *port
	}
	if *rootDir != "./uploads" {
		serverConfig.RootDir = *rootDir
	}
	if *username != "admin" {
		serverConfig.Username = *username
	}
	if *password != "admin123" {
		serverConfig.Password = *password
	}
	if *logLevel != "INFO" {
		serverConfig.LogLevel = *logLevel
	}
	if *logFile != "" {
		serverConfig.LogFile = *logFile
	}

	// Initialize users
	validUsers := make(map[string]models.User)

	// Check if password is already hashed, hash it if not
	// Simple detection: bcrypt hash usually starts with $2a$, $2b$ or $2y$
	passwordValue := serverConfig.Password
	if !strings.HasPrefix(passwordValue, "$2a$") && !strings.HasPrefix(passwordValue, "$2b$") && !strings.HasPrefix(passwordValue, "$2y$") {
		// Password is not hashed, hash it
		hashedPassword, err := auth.HashPassword(passwordValue)
		if err != nil {
			logs.Error("Failed to hash password: %v", err)
		} else {
			passwordValue = hashedPassword
			logs.Info("Password has been automatically hashed")
		}
	}

	validUsers[serverConfig.Username] = models.User{
		Username: serverConfig.Username,
		Password: passwordValue,
	}

	// Ensure root directory exists
	if err := utils.EnsureDirectoryExists(serverConfig.RootDir); err != nil {
		logs.Fatal("Failed to create root directory: %v", err)
	}

	// Create authentication middleware
	authMiddleware := auth.NewAuthMiddleware(validUsers).MiddlewareFunc

	// Set up file server handler
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(serverConfig.RootDir))))

	// Set up upload handler
	http.HandleFunc("/upload", handlers.HandleUpload(serverConfig.RootDir))

	// API routes
	http.HandleFunc("/api/login", handlers.HandleLogin(validUsers))
	// Add authentication verification endpoint
	// This endpoint needs to be registered before authMiddleware as it validates credentials itself
	// And returns success message in response if credentials are valid
	// Client JavaScript code is calling this endpoint for verification
	http.HandleFunc("/api/auth/verify", handlers.HandleAuthVerify(validUsers))
	http.HandleFunc("/api/upload", authMiddleware(handlers.HandleAPIUpload(serverConfig.RootDir)))
	http.HandleFunc("/api/files", authMiddleware(handlers.HandleAPIListFiles(serverConfig.RootDir)))
	http.HandleFunc("/api/files/", authMiddleware(handlers.HandleAPIDeleteFile(serverConfig.RootDir)))

	// Provide static file service
	staticFS := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFS))

	// Provide upload page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if it's root path
		if r.URL.Path == "/" {
			// Provide static HTML page
			http.ServeFile(w, r, "web/static/index.html")
			return
		}

		// For other paths, return 404 error
		http.NotFound(w, r)
	})

	// Provide file list page
	http.HandleFunc("/filelist", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Read upload directory
		files, err := os.ReadDir(serverConfig.RootDir)
		if err != nil {
			logs.Error("Failed to read file directory: %v", err)
			http.Error(w, "Failed to read file list", http.StatusInternalServerError)
			return
		}

		// Prepare template data
		var templateFiles []models.TemplateFileInfo
		for _, file := range files {
			// Get file information
			fileInfo, err := file.Info()
			if err != nil {
				logs.Error("Failed to get file information: %v", err)
				continue
			}

			// Create template file information
			templateFile := models.TemplateFileInfo{
				Name:             fileInfo.Name(),
				IsDir:            file.IsDir(),
				Size:             fileInfo.Size(),
				SizeFormatted:    utils.FormatFileSize(fileInfo.Size()),
				ModTime:          fileInfo.ModTime(),
				ModTimeFormatted: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			}

			templateFiles = append(templateFiles, templateFile)
		}

		// Load template
		tmpl, err := template.ParseFiles("templates/filelist.html")
		if err != nil {
			logs.Error("Failed to load template file: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Render template
		data := struct {
			Files []models.TemplateFileInfo
		}{Files: templateFiles}

		if err := tmpl.Execute(w, data); err != nil {
			logs.Error("Failed to render template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}))

	// Start server
	addr := fmt.Sprintf(":%s", serverConfig.ServerPort)
	logs.Info("put-file version: %s", version)
	logs.Info("File server started successfully, listening on: %s", addr)
	logs.Info("File root directory: %s", serverConfig.RootDir)
	logs.Info("Visit http://localhost:%s to upload files", serverConfig.ServerPort)
	logs.Info("Uploaded files can be accessed via http://localhost:%s/files/", serverConfig.ServerPort)

	// Wrap default ServeMux with request logging middleware
	handler := logs.RequestLogger(http.DefaultServeMux)

	if err := http.ListenAndServe(addr, handler); err != nil {
		logs.Fatal("Failed to start server: %v", err)
	}
}

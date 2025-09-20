// Copyright 2025 (C) https://github.com/3ziye team@3ziye.com

package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/3ziye/GoStaticServe/internal/auth"
	"github.com/3ziye/GoStaticServe/internal/config"
	"github.com/3ziye/GoStaticServe/internal/handlers"
	"github.com/3ziye/GoStaticServe/internal/logs"
	"github.com/3ziye/GoStaticServe/internal/models"
	"github.com/3ziye/GoStaticServe/internal/utils"
)

var version = "dev"

func main() {
	// 定义命令行参数
	configPath := flag.String("config", "config.json", "配置文件路径")
	rootDir := flag.String("root", "./uploads", "文件服务器的根目录")
	port := flag.String("port", "8080", "服务器监听端口")
	username := flag.String("username", "admin", "默认管理员用户名")
	password := flag.String("password", "admin123", "默认管理员密码")
	logLevel := flag.String("log-level", "INFO", "日志级别: DEBUG, INFO, WARN, ERROR, FATAL")
	logFile := flag.String("log-file", "", "日志文件路径，默认为标准输出")
	flag.Parse()

	// 初始化日志系统
	logs.InitLogger(&logs.LoggerConfig{
		Level:  config.LogLevel,
		Output: config.LogFile,
		Prefix: "file-server",
	})

	// 默认配置
	config := config.GetDefaultConfig()

	// 1. 尝试从配置文件加载配置
	if _, err := os.Stat(*configPath); err == nil {
		loadedConfig, err := config.LoadConfig(*configPath)
		if err != nil {
			logs.Warn("警告: 从配置文件加载配置失败，使用默认配置: %v", err)
		} else {
			config = loadedConfig
		}
	} else {
		logs.Info("配置文件 %s 不存在，使用默认配置", *configPath)
	}

	// 2. 应用环境变量（优先级高于配置文件）
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		config.ServerPort = envPort
	}
	if envRoot := os.Getenv("SERVER_ROOT"); envRoot != "" {
		config.RootDir = envRoot
	}
	if envUsername := os.Getenv("SERVER_USERNAME"); envUsername != "" {
		config.Username = envUsername
	}
	if envPassword := os.Getenv("SERVER_PASSWORD"); envPassword != "" {
		config.Password = envPassword
	}
	if envLogLevel := os.Getenv("SERVER_LOG_LEVEL"); envLogLevel != "" {
		config.LogLevel = envLogLevel
	}
	if envLogFile := os.Getenv("SERVER_LOG_FILE"); envLogFile != "" {
		config.LogFile = envLogFile
	}

	// 3. 应用命令行参数（优先级高于环境变量）
	if *port != "8080" {
		config.ServerPort = *port
	}
	if *rootDir != "./uploads" {
		config.RootDir = *rootDir
	}
	if *username != "admin" {
		config.Username = *username
	}
	if *password != "admin123" {
		config.Password = *password
	}
	if *logLevel != "INFO" {
		config.LogLevel = *logLevel
	}
	if *logFile != "" {
		config.LogFile = *logFile
	}

	// 初始化用户
	validUsers := make(map[string]models.User)
	
	// 检查密码是否已经哈希，如果没有则进行哈希处理
	// 简单检测：bcrypt哈希值通常以 $2a$、$2b$ 或 $2y$ 开头
	password := config.Password
	if !strings.HasPrefix(password, "$2a$") && !strings.HasPrefix(password, "$2b$") && !strings.HasPrefix(password, "$2y$") {
		// 密码未哈希，进行哈希处理
		hashedPassword, err := auth.HashPassword(password)
		if err != nil {
			logs.Error("密码哈希处理失败: %v", err)
		} else {
			password = hashedPassword
			logs.Info("密码已自动进行哈希处理")
		}
	}
	
	validUsers[config.Username] = models.User{
		Username: config.Username,
		Password: password,
	}

	// 确保根目录存在
	if err := utils.EnsureDirectoryExists(config.RootDir); err != nil {
		logs.Fatal("无法创建根目录: %v", err)
	}

	// 创建认证中间件
	authMiddleware := auth.NewAuthMiddleware(validUsers).MiddlewareFunc

	// 设置文件服务器处理程序
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(config.RootDir))))

	// 设置上传处理程序
	http.HandleFunc("/upload", handlers.HandleUpload(config.RootDir))

	// API路由
	http.HandleFunc("/api/login", handlers.HandleLogin(validUsers))
	http.HandleFunc("/api/upload", authMiddleware(handlers.HandleAPIUpload(config.RootDir)))
	http.HandleFunc("/api/files", authMiddleware(handlers.HandleAPIListFiles(config.RootDir)))
	http.HandleFunc("/api/files/", authMiddleware(handlers.HandleAPIDeleteFile(config.RootDir)))

	// 提供静态文件服务
	staticFS := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFS))

	// 提供上传页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 检查是否为根路径
		if r.URL.Path == "/" {
			// 提供静态HTML页面
			http.ServeFile(w, r, "web/static/index.html")
			return
		}

		// 对于其他路径，返回404错误
		http.NotFound(w, r)
	})

	// 提供文件列表页面
	http.HandleFunc("/filelist", authMiddleware.MiddlewareFunc(func(w http.ResponseWriter, r *http.Request) {
		// 读取上传目录
		files, err := os.ReadDir(config.RootDir)
		if err != nil {
			logs.Error("读取文件目录失败: %v", err)
			http.Error(w, "读取文件列表失败", http.StatusInternalServerError)
			return
		}

		// 准备模板数据
		var templateFiles []models.TemplateFileInfo
		for _, file := range files {
			// 获取文件信息
			fileInfo, err := file.Info()
			if err != nil {
				logs.Error("获取文件信息失败: %v", err)
				continue
			}

			// 创建模板文件信息
			templateFile := models.TemplateFileInfo{
				Name:            fileInfo.Name(),
				IsDir:           file.IsDir(),
				Size:            fileInfo.Size(),
				SizeFormatted:   utils.FormatFileSize(fileInfo.Size()),
				ModTime:         fileInfo.ModTime(),
				ModTimeFormatted: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			}

			templateFiles = append(templateFiles, templateFile)
		}

		// 加载模板
		tmpl, err := template.ParseFiles("templates/filelist.html")
		if err != nil {
			logs.Error("加载模板文件失败: %v", err)
			http.Error(w, "内部服务器错误", http.StatusInternalServerError)
			return
		}

		// 渲染模板
		data := struct {
			Files []models.TemplateFileInfo
		}{Files: templateFiles}

		if err := tmpl.Execute(w, data); err != nil {
			logs.Error("渲染模板失败: %v", err)
			http.Error(w, "内部服务器错误", http.StatusInternalServerError)
			return
		}
	}))

	// 启动服务器
	addr := fmt.Sprintf(":%s", config.ServerPort)
	logs.Info("GoStaticServe 版本: %s", version)
	logs.Info("文件服务器启动成功，监听地址: %s", addr)
	logs.Info("文件根目录: %s", config.RootDir)
	logs.Info("访问 http://localhost:%s 上传文件", config.ServerPort)
	logs.Info("上传的文件可通过 http://localhost:%s/files/ 访问", config.ServerPort)

	// 使用请求日志中间件包装默认的ServeMux
	handler := logs.RequestLogger(http.DefaultServeMux)

	if err := http.ListenAndServe(addr, handler); err != nil {
		logs.Fatal("服务器启动失败: %v", err)
	}
}
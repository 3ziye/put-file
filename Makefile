# put-file Makefile

# 项目名称
PROJECT_NAME = put-file

# 主程序入口
MAIN_SRC = cmd/server/main.go

# 输出二进制文件
BINARY = $(PROJECT_NAME)

# Go编译参数
GOFLAGS = -ldflags "-s -w"

# 默认目标：显示帮助信息
default: help

# 目标：编译项目
all: build

# 编译项目
build: 
	@echo "Building $(PROJECT_NAME)..."
	go build $(GOFLAGS) -o $(BINARY) $(MAIN_SRC)
	@echo "Build completed successfully!"

# 清理编译产物
clean: 
	@echo "Cleaning up..."
	rm -f $(BINARY)
	@echo "Clean completed!"

# 运行项目
run: build
	@echo "Running $(PROJECT_NAME)..."
	./$(BINARY)

# 构建并安装到系统
install: build
	@echo "Installing $(PROJECT_NAME)..."
	install -d /usr/local/bin
	install -m 755 $(BINARY) /usr/local/bin/
	@echo "Installation completed!"

# 从系统中卸载
uninstall: 
	@echo "Uninstalling $(PROJECT_NAME)..."
	rm -f /usr/local/bin/$(BINARY)
	@echo "Uninstallation completed!"

# 交叉编译为Linux 64位
build-linux: 
	@echo "Building for Linux 64-bit..."
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $(BINARY)-linux-amd64 $(MAIN_SRC)
	@echo "Linux build completed!"

# 交叉编译为Windows 64位
build-windows: 
	@echo "Building for Windows 64-bit..."
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o $(BINARY)-windows-amd64.exe $(MAIN_SRC)
	@echo "Windows build completed!"

# 交叉编译为macOS 64位
build-macos: 
	@echo "Building for macOS 64-bit..."
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o $(BINARY)-darwin-amd64 $(MAIN_SRC)
	@echo "macOS build completed!"

# 运行测试
test: 
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed!"

# 显示帮助信息
default: help

help: 
	@echo "============================================================="
	@echo "                  $(PROJECT_NAME) Makefile                    "
	@echo "============================================================="
	@echo "Usage: make [target]"
	@echo "If no target is specified, 'help' will be executed by default"
	@echo "Targets:"
	@echo "  help      : Show this help message (default)"
	@echo "  all       : Build the project"
	@echo "  build     : Build the project"
	@echo "  clean     : Clean up build artifacts"
	@echo "  run       : Build and run the project"
	@echo "  install   : Build and install the project to system"
	@echo "  uninstall : Uninstall the project from system"
	@echo "  build-linux   : Cross-compile for Linux 64-bit"
	@echo "  build-windows : Cross-compile for Windows 64-bit"
	@echo "  build-macos   : Cross-compile for macOS 64-bit"
	@echo "  test      : Run tests"
	@echo "  help      : Show this help message"
	@echo "-------------------------------------------------------------"
	@echo "Example: make build       # Compile the project"
	@echo "         make run         # Build and run the project"
	@echo "         make clean       # Clean up build artifacts"
	@echo "============================================================="

.PHONY: all build clean run install uninstall build-linux build-windows build-macos test help default
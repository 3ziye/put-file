# 文件服务器开发指南

本文档提供了文件服务器项目的开发指南，帮助开发者快速设置开发环境、理解开发流程并进行代码贡献。

## 开发环境准备

### 必要工具

1. **Go开发环境**
   - 版本要求：Go 1.21或更高版本
   - 下载地址：https://golang.org/dl/

2. **Git**
   - 用于代码管理
   - 下载地址：https://git-scm.com/downloads

3. **文本编辑器/IDE**
   - 推荐使用VSCode + Go插件，或GoLand
   - VSCode下载：https://code.visualstudio.com/
   - GoLand下载：https://www.jetbrains.com/go/

4. **Docker（可选）**
   - 用于容器化部署和测试
   - 下载地址：https://www.docker.com/products/docker-desktop

### 环境配置

1. 安装Go并配置`GOPATH`环境变量
2. 确保`$GOPATH/bin`在系统的`PATH`环境变量中
3. 安装Git并配置用户信息

## 使用VSCode远程容器开发环境

项目支持使用VSCode的Remote - Containers扩展进行开发，这可以提供一个预配置的、隔离的开发环境，无需在本地安装Go等依赖。

### 前提条件

- 已安装VSCode
- 已安装VSCode的[Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)扩展
- 已安装Docker

### 使用步骤

1. 打开VSCode
2. 安装Remote - Containers扩展
3. 使用VSCode打开项目文件夹
4. 点击VSCode右下角的绿色远程指示器图标
5. 选择"Reopen in Container"选项
6. VSCode会自动构建容器并在容器内打开项目

### 开发环境特性

- **预配置环境**：包含所有必要的Go开发工具和依赖
- **热重载支持**：预装air工具支持代码热重载
- **调试支持**：预装delve调试器和相关配置
- **扩展推荐**：自动安装推荐的VSCode扩展
- **持久化缓存**：缓存Go依赖和构建文件，提高开发效率
- **文件同步**：本地文件与容器内文件实时同步

### 常用操作

- **运行应用**：在VSCode终端中执行`air -c .air.toml`启动带热重载的应用
- **调试应用**：使用VSCode的调试功能直接调试容器内的应用
- **安装依赖**：在VSCode终端中执行`go mod tidy`安装或更新依赖
- **代码格式化**：保存文件时自动格式化代码（已配置）

### 自定义配置

如需自定义开发环境配置：

1. 修改`.devcontainer/devcontainer.json`文件自定义容器配置、扩展和设置
2. 修改`.devcontainer/Dockerfile`文件自定义容器镜像
3. 修改`.devcontainer/docker-compose.yml`文件自定义多容器配置（如添加数据库等服务）

## 获取源代码

1. 克隆代码仓库：

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
```

2. 初始化Go模块：

```bash
go mod tidy
```

## 项目结构回顾

请参考[架构文档](architecture.md)了解项目的整体结构和各模块的功能。

## 本地开发流程

### 运行开发服务器

在开发过程中，可以使用以下命令启动开发服务器：

```bash
go run cmd/server/main.go
```

也可以指定参数：

```bash
go run cmd/server/main.go --port=8081 --root=./myfiles
```

### 构建项目

使用以下命令构建项目：

```bash
go build -o file-server cmd/server/main.go
```

构建完成后，可以直接运行生成的可执行文件：

```bash
./file-server
```

### 运行测试

使用以下命令运行测试：

```bash
go test ./tests/...
```

## 代码风格指南

为了保持代码的一致性和可读性，请遵循以下代码风格指南：

1. **Go代码风格**
   - 遵循Go语言的标准代码风格
   - 使用`gofmt`格式化代码
   - 使用`golint`检查代码质量

2. **命名规范**
   - 包名使用小写字母，不使用下划线
   - 函数名使用驼峰命名法，首字母大写表示公开函数
   - 变量名使用驼峰命名法，首字母小写表示私有变量

3. **注释规范**
   - 为公开函数、类型和变量添加注释
   - 注释应该清晰地描述代码的功能和用途
   - 使用GoDoc格式编写注释

## 配置管理

开发过程中，可以通过以下方式配置项目：

1. **配置文件**：编辑`config.json`文件
2. **环境变量**：设置相应的环境变量
3. **命令行参数**：启动时指定命令行参数

配置项的优先级从高到低为：命令行参数 > 环境变量 > 配置文件 > 默认配置。

主要配置项包括：
- ServerPort：服务器端口
- RootDir：文件根目录
- Username：管理员用户名
- Password：管理员密码
- LogLevel：日志级别（DEBUG/INFO/WARN/ERROR/FATAL）
- LogFile：日志文件路径（为空时输出到标准输出）

## 调试技巧

1. **使用日志**：在代码中添加日志输出，使用`logs`包的不同级别日志函数（Debug/Info/Warn/Error/Fatal）
   - 可以通过设置`--log-level=DEBUG`参数来查看更详细的日志信息
2. **使用Delve调试器**：Go的官方调试器
   ```bash
dlv debug cmd/server/main.go
```
3. **使用VSCode调试**：配置VSCode的调试功能，添加launch.json文件

## 测试指南

### 单元测试

单元测试应该放在被测包的同一目录下，文件名以`_test.go`结尾。例如，对于`internal/handlers`包，测试文件应该命名为`handlers_test.go`。

### 集成测试

集成测试应该放在`tests`目录下，测试多个组件的协同工作。

## Docker开发环境

对于使用Docker进行开发的开发者，可以使用以下命令构建和运行Docker镜像：

```bash
# 构建镜像
docker build -t file-server-dev .

# 运行容器
docker run -p 8080:8080 -v $(pwd):/app file-server-dev
```

## CI/CD流程

项目使用GitHub Actions进行持续集成和持续部署。CI/CD工作流配置在`.github/workflows/ci.yml`文件中，包含以下步骤：

1. 检出代码
2. 设置Go环境
3. 安装依赖
4. 运行测试
5. 构建项目

## 提交代码流程

1. 确保代码通过所有测试
2. 使用`gofmt`格式化代码
3. 提交代码，编写清晰的提交信息
4. 创建Pull Request

## 常见问题解决

### 1. 依赖问题

如果遇到依赖问题，可以尝试以下命令：

```bash
go clean -modcache
go mod tidy
```

### 2. 端口占用问题

如果启动服务器时遇到端口占用问题，可以使用以下命令查看占用端口的进程，并关闭它：

```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Linux/macOS
lsof -i :8080
kill -9 <PID>
```

或者，使用`--port`参数指定一个未被占用的端口。

### 3. 文件权限问题

如果上传或访问文件时遇到权限问题，确保上传目录具有正确的读写权限：

```bash
# Linux/macOS
chmod -R 755 ./uploads

# Windows
# 在文件资源管理器中设置权限
```

## 扩展开发

### 添加新功能

1. 确定功能所属的模块
2. 在相应的包中实现功能
3. 更新相关的接口和配置
4. 编写测试用例
5. 更新文档

### 修改现有功能

1. 了解现有功能的实现逻辑
2. 进行必要的修改
3. 确保所有测试通过
4. 更新相关文档

## 文档更新

添加或修改功能后，请及时更新相关文档：

1. README.md：项目概述、安装和使用方法
2. doc/api.md：API接口文档
3. doc/architecture.md：架构设计文档
4. doc/development.md：开发指南文档

## 联系方式

如有任何问题或建议，请联系项目维护者。
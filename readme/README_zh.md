# put-file

<h1 align="center" style="border-bottom: none"> 
     <a href="" target="_blank"> 
         <alt="put-file" src="" width="100" height="100"> 
     </a> 
     <br>put-file 
 </h1> 
 
 <div align="center" style="line-height: 2;"> 
   [<a href="/README.md">English</a>] | [<a href="/readme/README_ar.md">阿拉伯语</a>] | [<a href="/readme/README_da.md">丹麦语</a>] | [<a href="/readme/README_de.md">德语</a>] | [<a href="/readme/README_es.md">西班牙语</a>] | [<a href="/readme/README_fr.md">法语</a>] | [<a href="/readme/README_it.md">意大利语</a>] | [<a href="/readme/README_ja.md">日语</a>] | [<a href="/readme/README_ko.md">韩语</a>] | [<a href="/readme/README_nl.md">荷兰语</a>] | [<a href="/readme/README_no.md">挪威语</a>] | [<a href="/readme/README_pl.md">波兰语</a>] | [<a href="/readme/README_pt.md">葡萄牙语</a>] | [<a href="/readme/README_ru.md">俄语</a>] | [<a href="/readme/README_sv.md">瑞典语</a>] | [<a href="/readme/README_th.md">泰语</a>] | [<a href="/readme/README_vi.md">越南语</a>] | [<a href="/readme/README_zh.md">中文(简体)</a>] 
   <br> 
   
   | ** [问题反馈](https://github.com/3ziye/put-file/issues) ** | ** [发布版本](https://github.com/3ziye/put-file/releases) ** | ** [项目说明](https://github.com/3ziye/put-file/blob/main/README.md) ** | ** [架构文档](https://github.com/3ziye/put-file/blob/main/doc/architecture.md) ** | 
   <br> 
   
   [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) 
   &nbsp;&nbsp; 
   ![Go](https://img.shields.io/badge/language-Go-blue.svg) 
   &nbsp;&nbsp; 
   ![performance](https://img.shields.io/badge/performance-high-yellow.svg) 
   &nbsp;&nbsp; 
   ![status](https://img.shields.io/badge/status-Stable-green.svg) 
 </div> 
 
 <p align="center">put-file是一个高性能、轻量级的静态文件服务器，使用Go语言开发。它支持文件上传、下载、删除等基本操作，同时提供了权限控制、日志记录等功能，是简单文件服务需求的理想选择。</p>

## 功能特性

- 🚀 **高性能**：基于Go语言的高并发特性，能够高效处理大量文件请求
- 🔒 **权限控制**：支持基本的身份验证和文件访问权限控制
- 📝 **详细日志**：记录请求和操作日志，便于问题排查和监控
- 📁 **文件操作**：支持文件上传、下载、删除、重命名等操作
- 📋 **目录列表**：自动生成目录列表，方便浏览和访问文件
- ⚙️ **配置灵活**：支持通过配置文件、环境变量和命令行参数进行配置
- 🐳 **容器支持**：提供Docker镜像，便于部署和运行
- 🚀 **自动部署**：支持通过GitHub Actions自动部署到远程服务器

## 快速开始

### 安装

#### 使用Go安装

```bash
go install github.com/3ziye/put-file@latest
```

#### 从源码构建

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### 使用预编译的二进制文件

put-file提供了Linux、Windows和Mac系统的预编译二进制文件，可以直接下载使用。

1. 访问[GitHub Releases页面](https://github.com/3ziye/put-file/releases)，下载对应平台的二进制文件压缩包

2. 根据您的操作系统解压对应的文件:

   **Linux:**
   ```bash
   # 下载并解压Linux版本
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # 运行服务
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # 下载并解压Windows版本
   # 右键点击下载的zip文件，选择"解压到当前文件夹"
   
   # 运行服务
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # 下载并解压Mac版本
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # 运行服务
   ./put-file
   ```

3. 查看版本信息
   启动服务后，可以在控制台日志中看到当前版本号

#### 使用Docker

从Docker Hub拉取镜像:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

从GitHub Package拉取镜像:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker镜像

### 从GitHub Package拉取镜像

1. 确保你已安装Docker
2. 运行以下命令拉取镜像:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### 运行Docker镜像

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

此命令将运行容器，将容器的8080端口映射到主机的8080端口，并将主机的`./files`目录挂载到容器内的`/app/uploads`目录。

### 将镜像推送到GitHub Package (仅维护者)

1. 登录GitHub容器注册表:
   ```bash
docker login ghcr.io -u 您的GitHub用户名 -p 您的GitHub令牌
   ```

2. 构建镜像:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. 推送镜像:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## 自动部署到服务器

put-file支持通过GitHub Actions自动部署到远程服务器。有关详细的配置步骤，请参阅[部署文档](doc/DEPLOYMENT.md)。

自动部署的主要功能:
- 🔑 GitHub Secrets管理服务器凭证
- 📥 自动下载Linux二进制文件
- 📁 服务器备份和部署
- 🚀 systemd服务支持
- ⚡ 通过Release或手动操作触发部署

## 其他语言版本

- [English](README_en.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
- [Português](README_pt.md)
- [Italiano](README_it.md)
- [한국어](README_ko.md)
- [العربية](README_ar.md)
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
# put-file

<h1 align="center" style="border-bottom: none"> 
     <a href="" target="_blank"> 
         <alt="put-file" src="" width="100" height="100"> 
     </a> 
     <br>put-file 
 </h1> 
 
 <div align="center" style="line-height: 2;"> 
   [<a href="/README.md">English</a>] | [<a href="/readme/README_ar.md">العربية</a>] | [<a href="/readme/README_da.md">Dansk</a>] | [<a href="/readme/README_de.md">Deutsch</a>] | [<a href="/readme/README_es.md">Español</a>] | [<a href="/readme/README_fr.md">Français</a>] | [<a href="/readme/README_it.md">Italiano</a>] | [<a href="/readme/README_ja.md">日本語</a>] | [<a href="/readme/README_ko.md">한국어</a>] | [<a href="/readme/README_nl.md">Nederlands</a>] | [<a href="/readme/README_no.md">Norsk</a>] | [<a href="/readme/README_pl.md">Polski</a>] | [<a href="/readme/README_pt.md">Português</a>] | [<a href="/readme/README_ru.md">Русский</a>] | [<a href="/readme/README_sv.md">Svenska</a>] | [<a href="/readme/README_th.md">ไทย</a>] | [<a href="/readme/README_vi.md">Tiếng Việt</a>] | [<a href="/readme/README_zh.md">中文(简体)</a>] 
   <br> 
   
   | ** [Issues](https://github.com/3ziye/put-file/issues) ** | ** [Releases](https://github.com/3ziye/put-file/releases) ** | ** [README](https://github.com/3ziye/put-file/blob/main/README.md) ** | ** [Architecture](https://github.com/3ziye/put-file/blob/main/doc/architecture.md) ** | 
   <br> 
   
   [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) 
   &nbsp;&nbsp; 
   ![Go](https://img.shields.io/badge/language-Go-blue.svg) 
   &nbsp;&nbsp; 
   ![performance](https://img.shields.io/badge/performance-high-yellow.svg) 
   &nbsp;&nbsp; 
   ![status](https://img.shields.io/badge/status-Stable-green.svg) 
 </div> 
 
 <p align="center">put-file is a high-performance, lightweight static file server developed in Go language. It supports basic operations such as file upload, download, deletion, and also provides features like permission control and logging, making it an ideal choice for simple file serving needs.</p>

## Features

- 🚀 **High Performance**: Based on Go's high concurrency features, capable of efficiently handling a large number of file requests
- 🔒 **Permission Control**: Supports basic authentication and file access permission control
- 📝 **Detailed Logging**: Records requests and operation logs for troubleshooting and monitoring
- 📁 **File Operations**: Supports file upload, download, deletion, renaming, etc.
- 📋 **Directory Listing**: Automatically generates directory listings for easy browsing and file access
- ⚙️ **Flexible Configuration**: Supports configuration via config files, environment variables, and command line parameters
- 🐳 **Container Support**: Provides Docker images for easy deployment and operation
- 🚀 **Automatic Deployment**: Supports automatic deployment to remote servers via GitHub Actions

## Quick Start

### Installation

#### Install using Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Build from source

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Use precompiled binary files

put-file provides precompiled binary files for Linux, Windows, and Mac systems, which can be downloaded directly for use.

1. Visit the [GitHub Releases page](https://github.com/3ziye/put-file/releases) and download the compressed package of binary files corresponding to your platform

2. Extract the corresponding files according to your operating system:

   **Linux:**
   ```bash
   # Download and extract Linux version
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Run the service
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Download and extract Windows version
   # Right-click the downloaded zip file and select "Extract to current folder"
   
   # Run the service
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Download and extract Mac version
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Run the service
   ./put-file
   ```

3. Check version information
   After starting the service, you can see the current version number in the console logs

#### Use Docker

Pull the image from Docker Hub:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Pull the image from GitHub Package:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker Image

### Pull the image from GitHub Package

1. Make sure you have Docker installed
2. Run the following command to pull the image:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Run the Docker image

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

This command will run the container, map port 8080 of the container to port 8080 of the host, and mount the `./files` directory of the host to the `/app/uploads` directory inside the container.

### Push the image to GitHub Package (maintainers only)

1. Log in to GitHub Container Registry:
   ```bash
docker login ghcr.io -u YOUR_GITHUB_USERNAME -p YOUR_GITHUB_TOKEN
   ```

2. Build the image:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Push the image:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Automatic Deployment to Server

put-file supports automatic deployment to remote servers via GitHub Actions. For detailed configuration steps, please refer to the [deployment documentation](doc/DEPLOYMENT.md).

Key features of automatic deployment:
- 🔑 GitHub Secrets management for server credentials
- 📥 Automatic download of Linux binaries
- 📁 Server backup and deployment
- 🚀 systemd service support
- ⚡ Trigger deployment via Release or manual operation

## Other Language Versions

- [中文](README_zh.md)
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
# GoStaticServe

<h1 align="center" style="border-bottom: none"> 
     <a href="" target="_blank"> 
         <alt="GoStaticServe" src="" width="100" height="100"> 
     </a> 
     <br>GoStaticServe 
 </h1> 
 
 <div align="center" style="line-height: 2;"> 
   [<a href="/README.md">English</a>] | [<a href="/readme/README_ar.md">العربية</a>] | [<a href="/readme/README_da.md">Dansk</a>] | [<a href="/readme/README_de.md">Deutsch</a>] | [<a href="/readme/README_es.md">Español</a>] | [<a href="/readme/README_fr.md">Français</a>] | [<a href="/readme/README_it.md">Italiano</a>] | [<a href="/readme/README_ja.md">日本語</a>] | [<a href="/readme/README_ko.md">한국어</a>] | [<a href="/readme/README_nl.md">Nederlands</a>] | [<a href="/readme/README_no.md">Norsk</a>] | [<a href="/readme/README_pl.md">Polski</a>] | [<a href="/readme/README_pt.md">Português</a>] | [<a href="/readme/README_ru.md">Русский</a>] | [<a href="/readme/README_sv.md">Svenska</a>] | [<a href="/readme/README_th.md">ไทย</a>] | [<a href="/readme/README_vi.md">Tiếng Việt</a>] | [<a href="/readme/README_zh.md">中文(简体)</a>] 
   <br> 
   
   | ** [Issues](https://github.com/3ziye/GoStaticServe/issues) ** | ** [Releases](https://github.com/3ziye/GoStaticServe/releases) ** | ** [README](https://github.com/3ziye/GoStaticServe/blob/main/README.md) ** | ** [Architecture](https://github.com/3ziye/GoStaticServe/blob/main/doc/architecture.md) ** | 
   <br> 
   
   [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) 
   &nbsp;&nbsp; 
   ![Go](https://img.shields.io/badge/language-Go-blue.svg) 
   &nbsp;&nbsp; 
   ![performance](https://img.shields.io/badge/performance-high-yellow.svg) 
   &nbsp;&nbsp; 
   ![status](https://img.shields.io/badge/status-Stable-green.svg) 
 </div> 
 
 <p align="center">GoStaticServe is a high-performance, lightweight static file server developed in Go language. It supports basic operations such as file upload, download, deletion, and also provides features like permission control and logging, making it an ideal choice for simple file serving needs.</p>

## About 3ziye Open Source Organization

<p align="center">Welcome to the 3ziye GitHub open source organization! We are rooted in technology and innovative in thinking, deeply cultivating the field of industrial digital intelligence. We are committed to building a technical exchange platform for developers and enterprises worldwide through open source collaboration, jointly injecting continuous momentum into industrial digital intelligence, and working together to create a technology ecosystem that empowers the future.</p>

<p align="center">🌟 <strong>Our Vision</strong><br>
To become a leading provider of digital solutions in the industry, driving enterprise innovation and development with technology, and creating greater value for customers.</p>

<p align="center">In the open source field, we hope to extend this vision — through opening core technologies, sharing solutions, breaking down technical barriers, helping more developers grow rapidly, assisting more enterprises in reducing the cost of digital transformation, and ultimately achieving the goal of "technology for all, ecological win-win"</p>

<p align="center">For more information, please visit our <a href="https://github.com/3ziye" target="_blank">GitHub Organization Page</a> or contact us at <a href="mailto:team@3ziye.com">team@3ziye.com</a>.</p>

## Features

- 🚀 **High Performance**: Based on Go's high concurrency features, capable of efficiently handling a large number of file requests
- 🔒 **Access Control**: Supports basic authentication and file access permission control
- 📝 **Detailed Logging**: Records requests and operation logs for troubleshooting and monitoring
- 📁 **File Operations**: Supports file upload, download, deletion, renaming, etc.
- 📋 **Directory Listing**: Automatically generates directory listings for easy browsing and file access
- ⚙️ **Flexible Configuration**: Supports configuration through configuration files, environment variables, and command-line parameters
- 🐳 **Container Support**: Provides Docker images for easy deployment and operation
- 🚀 **Automatic Deployment**: Supports automatic deployment to remote servers via GitHub Actions

## Quick Start

### Installation

#### Install Using Go

If you have Go installed, you can install GoStaticServe directly using the `go install` command:

```bash
# Install the latest version
go install github.com/3ziye/GoStaticServe@latest
```

#### Build from Source

1. Clone the repository
   ```bash
   git clone https://github.com/3ziye/GoStaticServe.git
   cd GoStaticServe
   ```

2. Build the project
   ```bash
   go build -o GoStaticServe cmd/main.go
   ```

3. Run the service
   ```bash
   ./GoStaticServe
   ```

#### Using Precompiled Binary Files

GoStaticServe provides precompiled binary files for Linux, Windows, and Mac systems, which can be downloaded and used directly. Each release package includes all necessary files such as the executable binary, configuration templates, web interface files, and installation scripts.

1. Visit the [GitHub Releases page](https://github.com/3ziye/GoStaticServe/releases) and download the compressed package of binary files corresponding to your platform

2. Extract the corresponding files according to your operating system:

#### Using Installation Scripts

Each release package includes dedicated installation scripts for each platform to simplify the installation process:

For Linux:

```bash
# Download and extract the Linux version
wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
cd GoStaticServe_vX.Y.Z_linux_amd64

# Run the installation script
chmod +x scripts/install_linux.sh
./scripts/install_linux.sh

# Or run the service directly without installation
./GoStaticServe
```

For Windows:

```bash
# Download and extract the Windows version
# Right-click on the downloaded zip file and select "Extract All"
cd GoStaticServe_vX.Y.Z_windows_amd64

# Run the installation script (double-click or run from command prompt)
scripts\install_windows.bat

# Or run the service directly without installation
GoStaticServe.exe
```

For Mac:

```bash
# Download and extract the Mac version
curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
cd GoStaticServe_vX.Y.Z_darwin_amd64

# Run the installation script
chmod +x scripts/install_mac.sh
./scripts/install_mac.sh

# Or run the service directly without installation
chmod +x GoStaticServe
./GoStaticServe
```

### What's Included in the Release Package

Each release package contains the following files and directories to ensure a complete experience:

- Executable binary file for the target platform
- `config.json.example`: Configuration template file
- `web/static/`: Web interface files
- `doc/`: Documentation files
- `uploads/`: Default directory for storing uploaded files
- `scripts/`: Installation scripts for each platform
- `README.md`: Project documentation

3. View version information
After starting the service, you can see the current version number in the console logs

#### Using Docker

Pull the image from Docker Hub:

```bash
docker pull 3ziye/gostaticserve:latest
```

Run the Docker container:

```bash
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Pull the image from GitHub Package:

```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
```

Run the GitHub Package container:

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## GitHub Package Docker Image

### Pulling the Image from GitHub Package

1. Ensure you have Docker installed
2. Run the following command to pull the image:

```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
```

### Running the GitHub Package Image

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

This command will run the container, map port 8080 of the container to port 8080 of the host, and mount the host's `./files` directory to the `/app/uploads` directory inside the container.

### Pushing the Image to GitHub Package (Maintainers Only)

1. Log in to GitHub Container Registry:

```bash
docker login ghcr.io -u YOUR_GITHUB_USERNAME -p YOUR_GITHUB_TOKEN
```

2. Build the image:

```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
```

3. Push the image:

```bash
docker push ghcr.io/3ziye/gostaticserve:latest
```

## Automatic Deployment to Server

GoStaticServe supports automatic deployment to remote servers via GitHub Actions. For detailed configuration steps, please refer to the [deployment documentation](doc/DEPLOYMENT.md).

Key features include:
- Secure management of server credentials via GitHub Secrets
- Automatic download of the latest Linux binary version
- Automatic backup and deployment on the server side
- Support for systemd service management
- Can be triggered manually or through Release publishing
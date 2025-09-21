# File Server Development Guide

This document provides development guidelines for the file server project, helping developers quickly set up a development environment, understand the development process, and contribute code.

## Development Environment Setup

### Required Tools

1. **Go Development Environment**
   - Version requirement: Go 1.21 or higher
   - Download link: https://golang.org/dl/

2. **Git**
   - For code management
   - Download link: https://git-scm.com/downloads

3. **Text Editor/IDE**
   - Recommended: VSCode + Go plugin, or GoLand
   - VSCode download: https://code.visualstudio.com/
   - GoLand download: https://www.jetbrains.com/go/

4. **Docker (Optional)**
   - For containerized deployment and testing
   - Download link: https://www.docker.com/products/docker-desktop

### Environment Configuration

1. Install Go and configure the `GOPATH` environment variable
2. Ensure `$GOPATH/bin` is in the system's `PATH` environment variable
3. Install Git and configure user information

## Using VSCode Remote Container Development Environment

The project supports development using VSCode's Remote - Containers extension, which provides a pre-configured, isolated development environment without the need to install Go and other dependencies locally.

### Prerequisites

- VSCode installed
- VSCode [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension installed
- Docker installed

### Usage Steps

1. Open VSCode
2. Install the Remote - Containers extension
3. Open the project folder with VSCode
4. Click the green remote indicator icon in the lower right corner of VSCode
5. Select the "Reopen in Container" option
6. VSCode will automatically build the container and open the project inside it

### Development Environment Features

- **Pre-configured Environment**: Contains all necessary Go development tools and dependencies
- **Hot Reload Support**: Pre-installed air tool for code hot reloading
- **Debug Support**: Pre-installed delve debugger and related configurations
- **Extension Recommendations**: Automatically installs recommended VSCode extensions
- **Persistent Caching**: Caches Go dependencies and build files to improve development efficiency
- **File Synchronization**: Real-time synchronization between local files and files inside the container

### Common Operations

- **Run Application**: Execute `air -c .air.toml` in the VSCode terminal to start the application with hot reloading
- **Debug Application**: Use VSCode's debugging features to directly debug the application inside the container
- **Install Dependencies**: Execute `go mod tidy` in the VSCode terminal to install or update dependencies
- **Code Formatting**: Code is automatically formatted when saving files (configured)

### Custom Configuration

To customize the development environment configuration:

1. Modify the `.devcontainer/devcontainer.json` file to customize container configuration, extensions, and settings
2. Modify the `.devcontainer/Dockerfile` file to customize the container image
3. Modify the `.devcontainer/docker-compose.yml` file to customize multi-container configurations (e.g., adding database services)

## Getting Source Code

1. Clone the repository:

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
```

2. Initialize Go modules:

```bash
go mod tidy
```

## Project Structure Review

Please refer to the [Architecture Document](architecture.md) to understand the overall structure of the project and the functionality of each module.

## Local Development Process

### Running the Development Server

During development, you can start the development server using the following command:

```bash
go run cmd/server/main.go
```

You can also specify parameters:

```bash
go run cmd/server/main.go --port=8081 --root=./myfiles
```

### Building the Project

Build the project using the following command:

```bash
go build -o file-server cmd/server/main.go
```

After building, you can run the generated executable directly:

```bash
./file-server
```

### Running Tests

Run tests using the following command:

```bash
go test ./tests/...
```

## Code Style Guide

To maintain code consistency and readability, please follow these code style guidelines:

1. **Go Code Style**
   - Follow Go's standard code style
   - Use `gofmt` to format code
   - Use `golint` to check code quality

2. **Naming Conventions**
   - Package names use lowercase letters, no underscores
   - Function names use camelCase, with uppercase first letter for public functions
   - Variable names use camelCase, with lowercase first letter for private variables

3. **Commenting Conventions**
   - Add comments for public functions, types, and variables
   - Comments should clearly describe the functionality and purpose of the code
   - Use GoDoc format for comments

## Configuration Management

During development, you can configure the project through the following methods:

1. **Configuration File**: Edit the `config.json` file
2. **Environment Variables**: Set corresponding environment variables
3. **Command-line Arguments**: Specify command-line arguments when starting

The priority of configuration items from highest to lowest is: command-line arguments > environment variables > configuration file > default configuration.

Main configuration items include:
- ServerPort: Server port
- RootDir: File root directory
- Username: Administrator username
- Password: Administrator password
- LogLevel: Log level (DEBUG/INFO/WARN/ERROR/FATAL)
- LogFile: Log file path (outputs to standard output when empty)

## Debugging Techniques

1. **Using Logs**: Add log output in code, use different level log functions from the `logs` package (Debug/Info/Warn/Error/Fatal)
   - You can view more detailed log information by setting the `--log-level=DEBUG` parameter
2. **Using Delve Debugger**: Go's official debugger
   ```bash
dlv debug cmd/server/main.go
```
3. **Using VSCode Debugging**: Configure VSCode's debugging features, add a launch.json file

## Testing Guide

### Unit Testing

Unit tests should be placed in the same directory as the tested package, with filenames ending in `_test.go`. For example, for the `internal/handlers` package, the test file should be named `handlers_test.go`.

### Integration Testing

Integration tests should be placed in the `tests` directory to test the collaborative work of multiple components.

## Docker Development Environment

For developers using Docker for development, you can build and run Docker images using the following commands:

```bash
# Build image
docker build -t file-server-dev .

# Run container
docker run -p 8080:8080 -v $(pwd):/app file-server-dev
```

## CI/CD Process

The project uses GitHub Actions for continuous integration and continuous deployment. The CI/CD workflow is configured in the `.github/workflows/ci.yml` file and includes the following steps:

1. Check out code
2. Set up Go environment
3. Install dependencies
4. Run tests
5. Build project

## Code Submission Process

1. Ensure code passes all tests
2. Format code using `gofmt`
3. Commit code with clear commit messages
4. Create a Pull Request

## Common Problem Solutions

### 1. Dependency Issues

If you encounter dependency issues, you can try the following commands:

```bash
go clean -modcache
go mod tidy
```

### 2. Port Occupancy Issues

If you encounter port occupancy issues when starting the server, you can use the following commands to check the process occupying the port and close it:

```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Linux/macOS
lsof -i :8080
kill -9 <PID>
```

Alternatively, specify an unoccupied port using the `--port` parameter.

### 3. File Permission Issues

If you encounter permission issues when uploading or accessing files, ensure the upload directory has the correct read and write permissions:

```bash
# Linux/macOS
chmod -R 755 ./uploads

# Windows
# Set permissions in File Explorer
```

## Extension Development

### Adding New Features

1. Determine which module the feature belongs to
2. Implement the feature in the corresponding package
3. Update related interfaces and configurations
4. Write test cases
5. Update documentation

### Modifying Existing Features

1. Understand the implementation logic of the existing feature
2. Make necessary modifications
3. Ensure all tests pass
4. Update related documentation

## Documentation Updates

After adding or modifying features, please update relevant documentation in a timely manner:

1. README.md: Project overview, installation, and usage methods
2. doc/api.md: API interface documentation
3. doc/architecture.md: Architecture design documentation
4. doc/development.md: Development guide documentation

## Contact Information

If you have any questions or suggestions, please contact the project maintainers.
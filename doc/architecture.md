# File Server Architecture Document

This document describes the overall architecture design, module division, and relationships between components of the file server project, helping developers understand the project structure and working principles.

## Project Architecture Overview

The file server adopts a typical Go language Web application architecture, following the principle of separation of concerns, dividing functions into multiple independent modules. The overall architecture is as follows:

```
┌─────────────────────────────────────┐
│             HTTP Layer              │
│  (cmd/server/main.go - Routing)     │
└───────────────┬─────────────────────┘
                │
┌───────────────┼─────────────────────┐
│               │                     │
▼               ▼                     ▼
┌───────────┐ ┌───────────┐ ┌────────────────┐
│ Handler   │ │ Auth      │ │    Config      │
│ handlers  │ │   auth    │ │    config      │
└─────┬─────┘ └───────────┘ └────────────────┘
      │
      │
      ▼
┌───────────┐ ┌───────────┐ ┌───────────┐
│  Model    │ │  Utility  │ │  Logging  │
│  models   │ │   utils   │ │   logs    │
└───────────┘ └───────────┘ └───────────┘
```

## Directory Structure Details

The project adopts a standard Go project directory structure, mainly including the following parts:

### cmd/server/

The entry point of the application, containing the `main.go` file, responsible for:
- Parsing command-line arguments and environment variables
- Loading configuration files
- Initializing various components
- Setting up HTTP routes and middleware
- Starting the web server

### internal/

Internal package directory containing the core functionality implementation of the project. According to Go language conventions, these packages will not be imported by external projects.

#### internal/auth/

Implements authentication-related functionality, mainly including:
- Basic authentication middleware
- User credential verification functions

#### internal/logs/

Implements a unified logging system, supporting different log levels and file output, mainly including:
- Logger interface definition
- Log level control
- Log output formatting

#### internal/models/

Defines data models and structures, mainly including:
- User: User model
- APIResponse: Unified API response format
- FileInfo: File information model

#### internal/config/

Responsible for configuration loading and management, supporting reading configuration from multiple sources:
- Configuration files (JSON format)
- Environment variables
- Command-line parameters

#### internal/handlers/

Implements HTTP request handlers, processing various HTTP requests:
- Web interface file upload
- API interface processing (login, file upload, file list, file deletion)
- Unified API response processing

#### internal/utils/

Provides common utility functions, such as directory creation and other basic functions.

### tests/

Contains project test files, used to verify the functionality of various components.

### doc/

Contains project documentation, such as API documentation, architecture documentation, etc.

## Core Processes

### Core Processes

### 1. Server Startup Process

1. Parse command-line arguments
2. Load configuration files
3. Apply environment variable overrides
4. Apply command-line argument overrides
5. Initialize user information
6. Ensure upload directory exists
7. Create authentication middleware
8. Set up HTTP routes
9. Start the web server

### 2. File Upload Process

#### Web Interface Upload
1. User accesses the upload page
2. Selects a file and clicks the upload button
3. Server receives the file and saves it to the specified directory
4. Returns an upload success page

#### API Upload
1. Client sends a POST request with Basic authentication
2. Authentication middleware verifies user credentials
3. Server receives the file and saves it to the specified directory
4. Returns a JSON format response

### 3. File Access Process

1. User accesses the file URL through a browser (/files/filename)
2. Server finds and returns the corresponding file
3. If the file does not exist, returns a 404 error

## Configuration Management

The file server supports three levels of configuration priority:
1. Command-line arguments (highest priority)
2. Environment variables
3. Configuration files
4. Default configuration (lowest priority)

This design allows users to flexibly configure the server according to their needs, either through persistent configuration via configuration files or temporary configuration through environment variables or command-line arguments.

## Security Considerations

1. **Authentication**: API interfaces are protected using Basic authentication to prevent unauthorized access
2. **Access Control**: Currently implements simple username and password-based access control
3. **File Size Limit**: Upload file size is limited to 10MB, which can be adjusted in the code
4. **Path Security**: Uses `filepath.Join` and standard library functions to handle file paths, preventing path traversal attacks

## Scalability Considerations

The project's modular design gives it good scalability, allowing for easy addition of new features:

1. **Adding new API interfaces**: Simply add new handler functions in the `internal/handlers/` directory and register routes in `main.go`
2. **Modifying authentication methods**: Can extend the `internal/auth/` package to implement more complex authentication methods
3. **Adding new configuration items**: Simply add new fields to the Config struct in `internal/config/config.go` and use them in relevant code

## Performance Optimization

1. **Static File Service**: Uses Go standard library's `http.FileServer` to handle static files with good performance
2. **Memory Limits**: Sets memory limits when parsing multipart/form-data to prevent memory overflow
3. **Concurrent Processing**: Go's HTTP server supports concurrent request processing by default, no additional configuration needed

## Monitoring and Maintenance

Currently, the project provides basic log output where you can see server listening address, root directory, and other information at startup. For production environments, consider adding a more complete logging system and monitoring mechanism.
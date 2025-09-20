# File Server Testing Guide

This document explains how to run test cases for the file server and provides an overview of the testing content.

## Test File Structure

The project contains the following test files located in the `tests` directory:

1. **main_test.go** - Basic function unit tests
2. **api_test.go** - API handler function unit tests
3. **e2e_test.go** - End-to-end tests

## Running Tests

### Run All Tests

To run all tests, including unit tests and end-to-end tests, execute the following command in the project root directory:

```bash
# Run all tests (including end-to-end tests)
go test ./tests/... -v
```

### Run Specific Test Files

You can also run specific test files separately:

```bash
# Run basic function tests
go test ./tests/main_test.go -v

# Run API tests
go test ./tests/api_test.go -v

# Run end-to-end tests
go test ./tests/e2e_test.go -v

# Run all tests
go test ./tests/... -v
```

### Skip End-to-End Tests

End-to-end tests require starting an actual server and may be slower. If you only want to run unit tests, you can use the `-short` flag:

```bash
# Skip end-to-end tests and run only unit tests
go test ./tests/... -v -short
```

## Test Content Overview

### 1. Basic Function Unit Tests (main_test.go)

- **ensureDirectoryExists** - Tests directory creation functionality, including creating non-existent directories and handling existing directories
- **sendAPIResponse** - Tests API response sending functionality, including success responses and failure responses

### 2. API Handler Function Unit Tests (api_test.go)

- **handleLogin** - Tests login functionality, including successful login, non-existent username, incorrect password, and unsupported HTTP methods
- **handleAPIUpload** - Tests file upload functionality, verifying that files are correctly saved and appropriate responses are returned
- **handleAPIListFiles** - Tests file list retrieval functionality, verifying that the list of files in a directory is correctly returned
- **handleAPIDeleteFile** - Tests file deletion functionality, including deleting existing files and non-existent files

### 3. End-to-End Tests (e2e_test.go)

End-to-end tests simulate a complete user interaction flow:

1. **Start Server** - Start the file server using a temporary configuration file and upload directory
2. **API Login Authentication** - Test login API for authentication
3. **API File Upload** - Test uploading files via the API
4. **API Get File List** - Test retrieving file list and verifying that the uploaded file is in the list
5. **File Download** - Test downloading files via HTTP and verifying file content
6. **API Delete File** - Test deleting files via the API
7. **Verify File Deletion** - Confirm that the file has been successfully deleted and can no longer be downloaded

## Test Environment Notes

- Tests create temporary directories and files, which are automatically cleaned up after testing
- End-to-end tests start an actual server instance listening on port 8081
- Running tests requires a Go development environment
- Tests do not affect existing configuration files and uploaded files

## Continuous Integration

These tests can be integrated into CI/CD workflows to ensure that code changes do not break existing functionality.

In a CI environment, you can simply run:

```bash
# Run all tests (including end-to-end tests)
go test ./tests/... -v

# Or run only unit tests
go test ./tests/... -v -short
```
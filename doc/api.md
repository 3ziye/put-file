# File Server API Documentation

This document provides detailed information about all API interfaces provided by the file server, including request formats, response formats, and usage examples.

## Authentication

All API interfaces, except for the login interface, require Basic authentication. Authentication information can be obtained through the login interface or using the username and password from the configuration.

## API Interface List

### 1. Login Interface

**URL**: `/api/login`

**Method**: `POST`

**Description**: Used to verify user credentials and obtain login status

**Request Body**: 
```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Response**: 
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "username": "admin"
  },
  "error": null
}
```

### 2. Upload File Interface

**URL**: `/api/upload`

**Method**: `POST`

**Description**: Upload files to the server

**Authentication**: Requires Basic authentication

**Request Body**: `multipart/form-data` containing a file field named `file`

**Response**: 
```json
{
  "success": true,
  "message": "File uploaded successfully",
  "data": {
    "name": "example.jpg",
    "size": 102400,
    "path": "/files/example.jpg",
    "createdAt": "2023-07-01T12:00:00Z"
  },
  "error": null
}
```

### 3. Get File List Interface

**URL**: `/api/files`

**Method**: `GET`

**Description**: Get list of all files on the server

**Authentication**: Requires Basic authentication

**Response**: 
```json
{
  "success": true,
  "message": "File list fetched successfully",
  "data": [
    {
      "name": "example1.jpg",
      "size": 102400,
      "path": "/files/example1.jpg",
      "createdAt": "2023-07-01T12:00:00Z"
    },
    {
      "name": "example2.txt",
      "size": 5120,
      "path": "/files/example2.txt",
      "createdAt": "2023-07-02T10:30:00Z"
    }
  ],
  "error": null
}
```

### 4. Delete File Interface

**URL**: `/api/files/{filename}`

**Method**: `DELETE`

**Description**: Delete specified file from the server

**Authentication**: Requires Basic authentication

**Path Parameters**: 
- `filename`: The name of the file to delete

**Response**: 
```json
{
  "success": true,
  "message": "File deleted successfully",
  "data": {
    "filename": "example.jpg"
  },
  "error": null
}
```

## API Response Format

All API interfaces follow a unified response format:

```json
{
  "success": boolean,    // Whether the operation was successful
  "message": string,     // Description of operation result
  "data": any,           // Data returned when operation is successful
  "error": string|null   // Error message returned when operation fails
}
```

## HTTP Status Codes

The API interfaces use the following HTTP status codes to indicate different response statuses:

- `200 OK`: Request processed successfully
- `400 Bad Request`: Request parameter error or format does not meet requirements
- `401 Unauthorized`: Authentication information not provided or authentication failed
- `404 Not Found`: Requested resource does not exist
- `500 Internal Server Error`: Internal server error

## Example Code

### Upload File using curl

```bash
curl -u admin:admin123 -F "file=@/path/to/file.txt" http://localhost:8080/api/upload
```

### Get File List using curl

```bash
curl -u admin:admin123 http://localhost:8080/api/files
```

### Delete File using curl

```bash
curl -X DELETE -u admin:admin123 http://localhost:8080/api/files/example.txt
```

### Upload File using JavaScript (Fetch API)

```javascript
const fileInput = document.querySelector('input[type="file"]');
const file = fileInput.files[0];
const formData = new FormData();
formData.append('file', file);

fetch('http://localhost:8080/api/upload', {
  method: 'POST',
  headers: {
    'Authorization': 'Basic ' + btoa('admin:admin123')
  },
  body: formData
})
.then(response => response.json())
.then(data => {
  console.log('Upload successful:', data);
})
.catch(error => {
  console.error('Upload failed:', error);
});
```
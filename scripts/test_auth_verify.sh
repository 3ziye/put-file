#!/bin/bash
# Script to test /api/auth/verify endpoint

# Set server information
SERVER_URL="http://localhost:8080/api/auth/verify"
USERNAME="admin"
PASSWORD="admin123"

# Check if curl is installed
if ! command -v curl &> /dev/null
then
    echo "Error: curl is not installed, please install curl first."
    exit 1
fi

# Execute curl request to test authentication endpoint
# -v Show detailed output
# -u Set basic authentication credentials
# -H Set content type
echo "Testing /api/auth/verify endpoint..."
curl -v -u "$USERNAME:$PASSWORD" -H "Content-Type: application/json" "$SERVER_URL"

# Show some debug information
echo "\nIf you receive a 404 error, please check:"
echo "1. Is the server running?"
echo "2. Is the server running on port 8080?"
echo "3. Are the username and password correct?"
echo "\nUse the following command to check server status (Linux/Mac):"
echo "ps aux | grep put-file"
echo "\nOr use Windows command:"
echo "tasklist | findstr server.exe"
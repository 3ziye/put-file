@echo off
REM Test script for /api/auth/verify endpoint

REM Set server information
set SERVER_URL=http://localhost:8080/api/auth/verify
set USERNAME=admin
set PASSWORD=admin123

REM Execute curl request to test authentication endpoint
REM -v Show detailed output
REM -u Set basic authentication credentials
REM -H Set content type
echo Testing /api/auth/verify endpoint...
curl -v -u %USERNAME%:%PASSWORD% -H "Content-Type: application/json" %SERVER_URL%

REM Display some debug information
echo.
echo If you receive a 404 error, please check:
echo 1. Is the server running?
echo 2. Is the server running on port 8080?
echo 3. Are the username and password correct?
echo.
echo Use the following command to check server status:
echo tasklist ^| findstr server.exe
echo.
echo Use the following command to view all running Go programs:
echo tasklist ^| findstr go
echo.
pause
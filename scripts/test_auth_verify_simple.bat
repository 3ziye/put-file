@echo off
REM Simplified test script
echo Testing /api/auth/verify endpoint...
curl -v -u admin:admin123 http://localhost:8080/api/auth/verify
pause
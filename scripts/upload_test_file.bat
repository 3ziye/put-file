@echo off
curl -u admin:admin123 -F "file=@%~dp0test.png" http://localhost:8080/upload
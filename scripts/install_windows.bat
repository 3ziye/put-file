@echo off

REM put-file Windows Installation Script

REM Set installation directory
set INSTALL_DIR=%USERPROFILE%\put-file

echo ======================================
echo put-file Installation Script

echo This script will install put-file in %INSTALL_DIR%

echo Administrator privileges may be required during installation.
echo ======================================

REM Create installation directory
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

REM Copy files to installation directory
echo Copying files to installation directory...
copy put-file.exe "%INSTALL_DIR%"
copy config.json.example "%INSTALL_DIR%\config.json"
copy README.md "%INSTALL_DIR%"
xcopy /s web "%INSTALL_DIR%\web"
xcopy /s doc "%INSTALL_DIR%\doc"

REM Create uploads directory
if not exist "%INSTALL_DIR%\uploads" mkdir "%INSTALL_DIR%\uploads"

echo Installation completed!
echo 

echo Configuration notes:
echo - Configuration file is located at %INSTALL_DIR%\config.json

echo Usage:
echo 1. Open Command Prompt

echo 2. Run the following commands to start the server:
echo    cd %INSTALL_DIR%
echo    put-file.exe

echo 3. Open browser and visit http://localhost:8080

echo To modify default port or other configurations, please edit the config.json file.
echo ======================================

pause
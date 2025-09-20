@echo off

REM GoStaticServe Windows 安装脚本

REM 设置安装目录
sets INSTALL_DIR=%USERPROFILE%\GoStaticServe

echo ======================================
echo GoStaticServe 安装脚本

echo 此脚本将在 %INSTALL_DIR% 安装 GoStaticServe

echo 安装过程中可能需要管理员权限。
echo ======================================

REM 创建安装目录
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

REM 复制文件到安装目录
echo 复制文件到安装目录...
copy GoStaticServe.exe "%INSTALL_DIR%"
copy config.json.example "%INSTALL_DIR%\config.json"
copy README.md "%INSTALL_DIR%"
xcopy /s web "%INSTALL_DIR%\web"
xcopy /s doc "%INSTALL_DIR%\doc"

REM 创建uploads目录
if not exist "%INSTALL_DIR%\uploads" mkdir "%INSTALL_DIR%\uploads"

echo 安装完成！
echo 

echo 配置说明：
echo - 配置文件位于 %INSTALL_DIR%\config.json

echo 使用方法：
echo 1. 打开命令提示符

echo 2. 运行以下命令启动服务器：
echo    cd %INSTALL_DIR%
echo    GoStaticServe.exe

echo 3. 打开浏览器访问 http://localhost:8080

echo 如需修改默认端口或其他配置，请编辑 config.json 文件。
echo ======================================

pause
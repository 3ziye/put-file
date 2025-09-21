#!/bin/bash

# put-file Linux Installation Script

# Set installation directory
default_install_dir="$HOME/put-file"
echo "======================================="
echo "put-file Installation Script"
echo "======================================="

read -p "Please enter installation directory [default: $default_install_dir]: " install_dir
install_dir=${install_dir:-$default_install_dir}

# Create installation directory
mkdir -p "$install_dir"

# Copy files to installation directory
echo "Copying files to installation directory..."
cp put-file "$install_dir"
cp config.json.example "$install_dir/config.json"
cp README.md "$install_dir"
cp -r web "$install_dir"
cp -r doc "$install_dir"

# Create uploads directory
mkdir -p "$install_dir/uploads"

# Set executable permissions
chmod +x "$install_dir/put-file"

# Create start script
cat > "$install_dir/start_server.sh" << EOF
#!/bin/bash
cd "$install_dir"
./put-file
EOF

chmod +x "$install_dir/start_server.sh"

echo "Installation completed!"
echo

echo "Configuration notes:"
echo "- Configuration file is located at $install_dir/config.json"

echo "Usage:"
echo "1. Open terminal"
echo "2. Run the following commands to start the server:"
echo "   cd $install_dir"
echo "   ./start_server.sh or ./put-file"
echo "3. 打开浏览器访问 http://localhost:8080"

echo "如需修改默认端口或其他配置，请编辑 config.json 文件。"
echo "======================================="

# 询问是否要立即启动服务器
read -p "是否立即启动服务器？(y/n) " start_now
if [[ $start_now == [Yy]* ]]; then
    cd "$install_dir"
    ./put-file
fi
#!/bin/bash

# GoStaticServe Linux 安装脚本

# 设置安装目录
default_install_dir="$HOME/GoStaticServe"
echo "======================================="
echo "GoStaticServe 安装脚本"
echo "======================================="

read -p "请输入安装目录 [默认: $default_install_dir]: " install_dir
install_dir=${install_dir:-$default_install_dir}

# 创建安装目录
mkdir -p "$install_dir"

# 复制文件到安装目录
echo "复制文件到安装目录..."
cp GoStaticServe "$install_dir"
cp config.json.example "$install_dir/config.json"
cp README.md "$install_dir"
cp -r web "$install_dir"
cp -r doc "$install_dir"

# 创建uploads目录
mkdir -p "$install_dir/uploads"

# 设置可执行权限
chmod +x "$install_dir/GoStaticServe"

# 创建启动脚本
cat > "$install_dir/start_server.sh" << EOF
#!/bin/bash
cd "$install_dir"
./GoStaticServe
EOF

chmod +x "$install_dir/start_server.sh"

echo "安装完成！"
echo

echo "配置说明："
echo "- 配置文件位于 $install_dir/config.json"

echo "使用方法："
echo "1. 打开终端"
echo "2. 运行以下命令启动服务器："
echo "   cd $install_dir"
echo "   ./start_server.sh 或 ./GoStaticServe"
echo "3. 打开浏览器访问 http://localhost:8080"

echo "如需修改默认端口或其他配置，请编辑 config.json 文件。"
echo "======================================="

# 询问是否要立即启动服务器
read -p "是否立即启动服务器？(y/n) " start_now
if [[ $start_now == [Yy]* ]]; then
    cd "$install_dir"
    ./GoStaticServe
fi
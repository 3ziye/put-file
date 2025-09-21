#!/bin/bash

# put-file Mac Installation Script

# Set installation directory
default_install_dir="$HOME/put-file"
echo "======================================="
echo "put-file Installation Script (Mac)"
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

# Create .command file for double-click startup convenience
cat > "$install_dir/Start Server.command" << EOF
#!/bin/bash
cd "$(dirname "$0")"
./start_server.sh
EOF

chmod +x "$install_dir/Start Server.command"

echo "Installation completed!"
echo

echo "Configuration notes:"
echo "- Configuration file is located at $install_dir/config.json"

echo "Usage:"
echo "1. Method 1: Double-click the $install_dir/Start Server.command file"
echo "2. Method 2: Open terminal and run the following commands:"
echo "   cd $install_dir"
echo "   ./start_server.sh or ./put-file"
echo "3. Open browser and visit http://localhost:8080"

echo "To modify the default port or other configurations, please edit the config.json file."
echo "======================================="

# Ask whether to start the server immediately
read -p "Would you like to start the server immediately? (y/n) " start_now
if [[ $start_now == [Yy]* ]]; then
    cd "$install_dir"
    ./put-file
fi
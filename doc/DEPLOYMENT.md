# put-file Deployment Guide

This document details how to automatically deploy put-file to a remote server using GitHub Actions.

## I. Automatic Deployment (CD) Configuration

put-file provides a complete CI/CD process, implementing automatic build and deployment through GitHub Actions.

### 1. GitHub Secrets Configuration

First, you need to add the following Secrets to the "Settings > Secrets and variables > Actions" page of your GitHub repository:

| Secret Name | Description | Example Value |
|------------|-------------|--------------|
| `DEPLOY_SERVER_HOST` | Server IP address or domain name | `192.168.1.100` or `example.com` |
| `DEPLOY_SERVER_USERNAME` | Server login username | `ubuntu` or `root` |
| `DEPLOY_SERVER_SSH_KEY` | SSH private key (with server login permissions) | Complete private key content |
| `DEPLOY_SERVER_PORT` | SSH port (optional, default 22) | `22` |

### 2. Server-Side Preparation

Before deployment, you need to complete the following preparation work on the target server:

1. **Create Deployment Directory**
   ```bash
   sudo mkdir -p /opt/put-file/uploads
   sudo chown -R your_username:your_group /opt/put-file
   ```

2. **Create Running User** (optional, recommended to run the service with a non-root user)
   ```bash
   sudo adduser --system --no-create-home --group put-file
   sudo chown -R put-file:put-file /opt/put-file
   ```

3. **Install and Configure systemd Service**
   
   Use the systemd service file template provided in the project (`doc/put-file.service`):
   
   ```bash
   # Copy the service file to systemd directory
   sudo cp put-file.service /etc/systemd/system/
   
   # Modify configuration according to actual environment
   sudo nano /etc/systemd/system/put-file.service
   
   # Enable and start the service
   sudo systemctl daemon-reload
   sudo systemctl enable put-file
   sudo systemctl start put-file
   
   # Check service status
   sudo systemctl status put-file
   ```

4. **Configure Firewall**
   ```bash
   # If using ufw firewall
   sudo ufw allow 8080/tcp
   
   # If using firewalld
   sudo firewall-cmd --add-port=8080/tcp --permanent
   sudo firewall-cmd --reload
   ```

## II. Automatic Deployment Process

### 1. Trigger Methods

Automatic deployment can be triggered in two ways:

1. **Release New Version**: After creating and publishing a GitHub Release, the deployment process is automatically triggered
2. **Manual Trigger**: Manually trigger in the GitHub Actions page, need to specify the version to deploy

### 2. Deployment Process Description

When the deployment process is triggered, the following steps will be executed:

1. Download the latest Linux amd64 version binary file from GitHub Release
2. Extract and prepare deployment files
3. Connect to the remote server via SSH
4. Backup current version files (if they exist)
5. Upload the new version binary file
6. Restart the put-file service

### 3. Manually Trigger Deployment

1. Go to the "Actions" page of the GitHub repository
2. Select the "Deploy to Server" workflow
3. Click the "Run workflow" button
4. Enter the version number to deploy (e.g.: v1.0.0) in the pop-up form
5. Click "Run workflow" to confirm

## III. Manual Deployment (Optional)

If you need to deploy manually or perform manual intervention when automatic deployment fails, you can follow these steps:

### 1. Download Binary File

Download the Linux binary file of the corresponding version from the GitHub Releases page:

```bash
wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
```

### 2. Extract and Deploy

```bash
# Extract files
mkdir -p /opt/put-file/backups/$(date +%Y%m%d%H%M%S)
cp -a /opt/put-file/put-file /opt/put-file/backups/$(date +%Y%m%d%H%M%S)/

tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz

# Deploy new version
sudo cp put-file /opt/put-file/
sudo chmod +x /opt/put-file/put-file
sudo chown put-file:put-file /opt/put-file/put-file

# Restart service
sudo systemctl restart put-file
```

## IV. Configuration File Deployment

put-file supports more flexible configuration through configuration files. If you need to use a configuration file:

1. **Create Configuration File**
   ```bash
   nano /opt/put-file/config.json
   ```

2. **Configuration File Content Example**
   ```json
   {
     "ServerPort": "8080",
     "RootDir": "/opt/put-file/uploads",
     "Username": "admin",
     "Password": "your_secure_password",
     "LogLevel": "INFO",
     "LogFile": "/var/log/put-file.log"
   }
   ```

3. **Update systemd Service Configuration**
   Modify the `ExecStart` line to start with the configuration file:
   ```
   ExecStart=/opt/put-file/put-file --config=/opt/put-file/config.json
   ```

4. **Restart Service**
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl restart put-file
   ```

## V. Post-Deployment Verification

After deployment is complete, you can verify if the service is running normally through the following methods:

1. **Check Service Status**
   ```bash
   sudo systemctl status put-file
   ```

2. **View Logs**
   ```bash
   # View systemd logs
   journalctl -u put-file
   
   # View application logs
   tail -f /var/log/put-file.log
   ```

3. **Access the Service**
   Access `http://server-ip:8080` in a browser to verify if the service is running normally

## VI. Troubleshooting Common Issues

1. **Service Startup Failure**
   - Check if the configuration file is correct
   - Check file and directory permissions
   - View logs for detailed error information

2. **SSH Connection Failure**
   - Check if the server information in GitHub Secrets is correct
   - Ensure the server's SSH port is open
   - Ensure the SSH key has sufficient permissions

3. **Automatic Deployment Not Triggered**
   - Check GitHub Actions workflow configuration
   - Confirm whether Release has been created and published correctly

4. **Deployment Successful but Service Unavailable**
   - Check firewall settings
   - Confirm server port is correctly mapped
   - View application logs for detailed information

Through the above steps, you can successfully configure and use the automatic deployment function of put-file to achieve continuous integration and continuous deployment of code.
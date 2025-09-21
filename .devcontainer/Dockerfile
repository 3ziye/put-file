# Use official Go image as base
FROM golang:1.24

# Install necessary system dependencies
RUN apt update && apt install -y \
    git \
    bash \
    gcc \
    curl \
    make \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /workspace

# Add Go binary directory to PATH
ENV PATH="/go/bin:/usr/local/go/bin:$PATH"

# Set Go module and proxy
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

# Create uploads directory and set permissions
RUN mkdir -p /workspace/uploads && chmod -R 755 /workspace/uploads

# Create air configuration directory
RUN mkdir -p /workspace/.air

# Copy project files (devcontainer will auto-mount workspace, this is just an example)
# COPY go.mod go.sum ./
# RUN go mod download

# Expose application and debug ports
EXPOSE 8080 2345

# Set default command
CMD ["bash"]
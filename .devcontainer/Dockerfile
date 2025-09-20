# 使用Go官方镜像作为基础
FROM golang:1.21-alpine

# 安装必要的系统依赖
RUN apk update && apk add --no-cache \
    git \
    bash \
    gcc \
    musl-dev \
    curl \
    wget \
    openssh-client \
    vim \
    nano \
    make \
    && rm -rf /var/cache/apk/*

# 设置工作目录
WORKDIR /workspace

# 安装Go开发工具
RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install gotest.tools/gotestsum@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    && go install github.com/cweill/gotests/...@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest

# 添加Go二进制目录到PATH
ENV PATH="/go/bin:/usr/local/go/bin:$PATH"

# 设置Go模块和代理
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

# 创建uploads目录并设置权限
RUN mkdir -p /workspace/uploads && chmod -R 755 /workspace/uploads

# 创建air配置文件目录
RUN mkdir -p /workspace/.air

# 复制项目文件（devcontainer会自动挂载工作区，这里仅作为示例）
# COPY go.mod go.sum ./
# RUN go mod download

# 暴露应用端口和调试端口
EXPOSE 8080 2345

# 设置默认命令
CMD ["bash"]
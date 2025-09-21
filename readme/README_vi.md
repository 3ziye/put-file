<h1 align="center" style="border-bottom: none"> 
     <a href="" target="_blank"> 
         <alt="put-file" src="" width="100" height="100"> 
     </a> 
     <br>put-file 
 </h1> 
 
 <div align="center" style="line-height: 2;"> 
   [<a href="/README.md">English</a>] | [<a href="/readme/README_ar.md">العربية</a>] | [<a href="/readme/README_da.md">Dansk</a>] | [<a href="/readme/README_de.md">Deutsch</a>] | [<a href="/readme/README_es.md">Español</a>] | [<a href="/readme/README_fr.md">Français</a>] | [<a href="/readme/README_it.md">Italiano</a>] | [<a href="/readme/README_ja.md">日本語</a>] | [<a href="/readme/README_ko.md">한국어</a>] | [<a href="/readme/README_nl.md">Nederlands</a>] | [<a href="/readme/README_no.md">Norsk</a>] | [<a href="/readme/README_pl.md">Polski</a>] | [<a href="/readme/README_pt.md">Português</a>] | [<a href="/readme/README_ru.md">Русский</a>] | [<a href="/readme/README_sv.md">Svenska</a>] | [<a href="/readme/README_th.md">ไทย</a>] | [<a href="/readme/README_vi.md">Tiếng Việt</a>] | [<a href="/readme/README_zh.md">中文(简体)</a>] 
   <br> 
   
   | ** [Issues](https://github.com/3ziye/put-file/issues) ** | ** [Releases](https://github.com/3ziye/put-file/releases) ** | ** [README](https://github.com/3ziye/put-file/blob/main/README.md) ** | ** [Architecture](https://github.com/3ziye/put-file/blob/main/doc/architecture.md) ** | 
   <br> 
   
   [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) 
   &nbsp;&nbsp; 
   ![Go](https://img.shields.io/badge/language-Go-blue.svg) 
   &nbsp;&nbsp; 
   ![performance](https://img.shields.io/badge/performance-high-yellow.svg) 
   &nbsp;&nbsp; 
   ![status](https://img.shields.io/badge/status-Stable-green.svg) 
 </div> 
 
 <p align="center">put-file là một máy chủ tệp tĩnh hiệu suất cao và nhẹ được phát triển bằng ngôn ngữ Go. Nó hỗ trợ các hoạt động cơ bản như tải lên, tải xuống và xóa tệp, cũng như cung cấp các tính năng như kiểm soát quyền hạn và ghi nhật ký.</p>

## Tính năng

- 🚀 **Hiệu suất cao**: Dựa trên đặc điểm đồng thời cao của ngôn ngữ Go, có khả năng xử lý hiệu quả một số lượng lớn yêu cầu tệp
- 🔒 **Kiểm soát quyền hạn**: Hỗ trợ xác thực cơ bản và kiểm soát quyền truy cập vào tệp
- 📝 **Nhật ký chi tiết**: Ghi lại các yêu cầu và nhật ký hoạt động cho mục đích khắc phục sự cố và giám sát
- 📁 **Hoạt động với tệp**: Hỗ trợ tải lên, tải xuống, xóa, đổi tên tệp, v.v.
- 📋 **Danh sách thư mục**: Tự động tạo danh sách thư mục để dễ dàng duyệt và truy cập tệp
- ⚙️ **Cấu hình linh hoạt**: Hỗ trợ cấu hình thông qua tệp cấu hình, biến môi trường và tham số dòng lệnh
- 🐳 **Hỗ trợ container**: Cung cấp hình ảnh Docker để dễ dàng triển khai và chạy
- 🚀 **Triển khai tự động**: Hỗ trợ triển khai tự động trên máy chủ từ xa thông qua GitHub Actions

## Khởi động nhanh

### Cài đặt

#### Cài đặt bằng Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Xây dựng từ mã nguồn

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Sử dụng tệp nhị phân đã biên dịch sẵn

put-file cung cấp các tệp nhị phân đã biên dịch sẵn cho các hệ thống Linux, Windows và Mac, có thể tải xuống trực tiếp để sử dụng.

1. Truy cập [trang phát hành GitHub](https://github.com/3ziye/put-file/releases) và tải xuống gói tệp nhị phân tương ứng với nền tảng của bạn

2. Giải nén các tệp tương ứng tùy thuộc vào hệ điều hành đang sử dụng:

   **Linux:**
   ```bash
   # Tải xuống và giải nén phiên bản cho Linux
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Chạy dịch vụ
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Tải xuống và giải nén phiên bản cho Windows
   # Nhấp chuột phải vào tệp zip đã tải xuống và chọn "Giải nén tại đây"
   
   # Chạy dịch vụ
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Tải xuống và giải nén phiên bản cho Mac
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Chạy dịch vụ
   ./put-file
   ```

3. Kiểm tra thông tin phiên bản
   Sau khi khởi động dịch vụ, bạn có thể xem số phiên bản hiện tại trong nhật ký console

#### Sử dụng Docker

Tải xuống hình ảnh từ Docker Hub:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Tải xuống hình ảnh từ GitHub Package:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## Hình ảnh Docker từ GitHub Package

### Tải xuống hình ảnh từ GitHub Package

1. Đảm bảo rằng Docker đã được cài đặt
2. Chạy lệnh sau để tải xuống hình ảnh:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Chạy hình ảnh Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Lệnh này sẽ chạy container, ánh xạ cổng 8080 của container đến cổng 8080 của máy chủ và gắn thư mục `./files` của máy chủ vào thư mục `/app/uploads` bên trong container.

### Tải hình ảnh lên GitHub Package (chỉ dành cho người bảo trì)

1. Đăng nhập vào GitHub Container Registry:
   ```bash
docker login ghcr.io -u TÊN_TÀI_KHOẢN_GITHUB -p TOKEN_GITHUB
   ```

2. Xây dựng hình ảnh:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Tải lên hình ảnh:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Triển khai tự động trên máy chủ

put-file hỗ trợ triển khai tự động trên máy chủ từ xa thông qua GitHub Actions. Để biết các bước cấu hình chi tiết, vui lòng tham khảo [tài liệu triển khai](doc/DEPLOYMENT.md).

Các tính năng chính của triển khai tự động:
- 🔑 Quản lý thông tin xác thực máy chủ thông qua GitHub Secrets
- 📥 Tải xuống tự động các tệp nhị phân Linux
- 📁 Sao lưu và triển khai trên máy chủ
- 🚀 Hỗ trợ dịch vụ systemd
- ⚡ Kích hoạt triển khai thông qua phiên bản (Release) hoặc thao tác thủ công

## Các phiên bản khác bằng các ngôn ngữ khác nhau

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
- [Português](README_pt.md)
- [Italiano](README_it.md)
- [한국어](README_ko.md)
- [العربية](README_ar.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
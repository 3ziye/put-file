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
 
 <p align="center">put-file는 Go 언어로 개발된 고성능 경량 정적 파일 서버입니다. 파일 업로드, 다운로드, 삭제와 같은 기본 작업을 지원하며, 권한 제어 및 로그 기록과 같은 기능도 제공합니다.</p>

## 기능

- 🚀 **고성능**: Go 언어의 높은 동시성 특성을 바탕으로 많은 파일 요청을 효율적으로 처리할 수 있습니다
- 🔒 **권한 제어**: 기본 인증 및 파일 액세스 권한 제어를 지원합니다
- 📝 **상세한 로깅**: 문제 해결 및 모니터링을 위해 요청과 작업 로그를 기록합니다
- 📁 **파일 작업**: 파일 업로드, 다운로드, 삭제, 이름 변경 등을 지원합니다
- 📋 **디렉토리 목록**: 파일 탐색과 액세스를 용이하게 하기 위해 자동으로 디렉토리 목록을 생성합니다
- ⚙️ **유연한 구성**: 구성 파일, 환경 변수 및 명령줄 매개변수를 통해 구성을 지원합니다
- 🐳 **컨테이너 지원**: 배포와 실행을 용이하게 하기 위해 Docker 이미지를 제공합니다
- 🚀 **자동 배포**: GitHub Actions를 통해 원격 서버에 자동으로 배포를 지원합니다

## 빠른 시작

### 설치

#### Go를 사용하여 설치

```bash
go install github.com/3ziye/put-file@latest
```

#### 소스 코드에서 빌드

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### 사전 컴파일된 바이너리 사용

put-file는 Linux, Windows 및 Mac 시스템을 위한 사전 컴파일된 바이너리 파일을 제공하며, 직접 다운로드하여 사용할 수 있습니다.

1. [GitHub 릴리스 페이지](https://github.com/3ziye/put-file/releases)에 접속하여 해당 플랫폼에 해당하는 바이너리 파일 패키지를 다운로드합니다

2. 운영 체제에 따라 해당 파일을 추출합니다:

   **Linux:**
   ```bash
   # Linux 버전 다운로드 및 추출
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # 서비스 실행
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Windows 버전 다운로드 및 추출
   # 다운로드한 zip 파일을 마우스 오른쪽 버튼으로 클릭하고 "여기에서 추출"을 선택합니다
   
   # 서비스 실행
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Mac 버전 다운로드 및 추출
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # 서비스 실행
   ./put-file
   ```

3. 버전 정보 확인
   서비스를 시작한 후 콘솔 로그에서 현재 버전 번호를 확인할 수 있습니다

#### Docker 사용

Docker Hub에서 이미지를 다운로드합니다:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

GitHub Package에서 이미지를 다운로드합니다:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker 이미지

### GitHub Package에서 이미지 다운로드

1. Docker가 설치되어 있는지 확인합니다
2. 다음 명령어를 실행하여 이미지를 다운로드합니다:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Docker 이미지 실행

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

이 명령어는 컨테이너를 실행하고, 컨테이너의 8080 포트를 호스트의 8080 포트에 매핑하고, 호스트의 `./files` 디렉토리를 컨테이너 내의 `/app/uploads` 디렉토리에 마운트합니다.

### GitHub Package에 이미지 업로드 (유지 관리자 전용)

1. GitHub Container Registry에 로그인합니다:
   ```bash
docker login ghcr.io -u GITHUB_USERNAME -p GITHUB_TOKEN
   ```

2. 이미지를 빌드합니다:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. 이미지를 업로드합니다:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## 서버 자동 배포

put-file는 GitHub Actions를 통해 원격 서버에 자동으로 배포할 수 있습니다. 구성에 대한 자세한 단계는 [배포 문서](doc/DEPLOYMENT.md)를 참조하세요.

자동 배포의 주요 기능:
- 🔑 GitHub Secrets를 통한 서버 자격 증명 관리
- 📥 Linux 바이너리 자동 다운로드
- 📁 서버 백업 및 배포
- 🚀 systemd 서비스 지원
- ⚡ 릴리스(Release) 또는 수동 작업을 통한 배포 트리거

## 다양한 언어의 다른 버전

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
- [Português](README_pt.md)
- [Italiano](README_it.md)
- [العربية](README_ar.md)
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
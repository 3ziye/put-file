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
 
 <p align="center">put-fileは、Go言語で開発された高性能で軽量な静的ファイルサーバーです。ファイルのアップロード、ダウンロード、削除などの基本的な操作をサポートし、アクセス権限管理やログ記録などの機能も提供します。</p>

## 主な機能

- 🚀 **高性能**：Go言語の高並行処理機能に基づいており、大量のファイルリクエストを効率的に処理できます
- 🔒 **アクセス権限管理**：基本的な認証とファイルアクセス権限管理をサポートします
- 📝 **詳細なログ記録**：トラブルシューティングやモニタリングのためにリクエストと操作ログを記録します
- 📁 **ファイル操作**：ファイルのアップロード、ダウンロード、削除、名前変更などをサポートします
- 📋 **ディレクトリ一覧**：簡単に閲覧してファイルにアクセスできるように、ディレクトリ一覧を自動的に生成します
- ⚙️ **柔軟な設定**：設定ファイル、環境変数、コマンドラインパラメータを通じて設定できます
- 🐳 **コンテナサポート**：デプロイと実行を容易にするためにDockerイメージを提供します
- 🚀 **自動デプロイ**：GitHub Actionsを介したリモートサーバーへの自動デプロイをサポートします

## クイックスタート

### インストール

#### Goを使用してインストール

```bash
go install github.com/3ziye/put-file@latest
```

#### ソースコードからビルド

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### プリコンパイル済みバイナリファイルを使用

put-fileは、Linux、Windows、Macシステム用のプリコンパイル済みバイナリファイルを提供しており、直接ダウンロードして使用できます。

1. [GitHub Releasesページ](https://github.com/3ziye/put-file/releases)にアクセスし、お使いのプラットフォームに対応するバイナリファイルの圧縮パッケージをダウンロードします

2. お使いのオペレーティングシステムに応じて、対応するファイルを解凍します：

   **Linux:**
   ```bash
   # Linuxバージョンをダウンロードして解凍
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # サービスを実行
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Windowsバージョンをダウンロードして解凍
   # ダウンロードしたzipファイルを右クリックし、「現在のフォルダーに解凍」を選択します
   
   # サービスを実行
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Macバージョンをダウンロードして解凍
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # サービスを実行
   ./put-file
   ```

3. バージョン情報を確認する
   サービスを起動すると、コンソールログに現在のバージョン番号が表示されます

#### Dockerを使用する

Docker Hubからイメージをプルします：
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

GitHub Packageからイメージをプルします：
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Dockerイメージ

### GitHub Packageからイメージをプルする

1. Dockerがインストールされていることを確認します
2. 以下のコマンドを実行してイメージをプルします：
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Dockerイメージを実行する

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

このコマンドにより、コンテナが実行され、コンテナのポート8080がホストのポート8080にマップされ、ホストの`./files`ディレクトリがコンテナ内の`/app/uploads`ディレクトリにマウントされます。

### GitHub Packageにイメージをプッシュする（メンテナーのみ）

1. GitHub Container Registryにログインします：
   ```bash
docker login ghcr.io -u あなたのGitHubユーザー名 -p あなたのGitHubトークン
   ```

2. イメージを構築します：
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. イメージをプッシュします：
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## サーバーへの自動デプロイ

put-fileは、GitHub Actionsを介してリモートサーバーに自動的にデプロイすることをサポートしています。詳細な設定手順については、[デプロイドキュメント](doc/DEPLOYMENT.md)を参照してください。

自動デプロイの主な機能：
- 🔑 サーバー資格情報のGitHub Secrets管理
- 📥 Linuxバイナリの自動ダウンロード
- 📁 サーバーのバックアップとデプロイ
- 🚀 systemdサービスのサポート
- ⚡ Releaseまたは手動操作によるデプロイのトリガー

## 他の言語バージョン

- [English](README_en.md)
- [中文](README_zh.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
- [Português](README_pt.md)
- [Italiano](README_it.md)
- [한국어](README_ko.md)
- [العربية](README_ar.md)
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
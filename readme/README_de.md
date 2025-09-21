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
 
 <p align="center">put-file ist ein leistungsstarker, leichter statischer Dateiserver, der in der Programmiersprache Go entwickelt wurde. Er unterstützt grundlegende Operationen wie Dateiupload, -download und -löschung und bietet darüber hinaus Funktionen wie Berechtigungskontrolle und Protokollierung.</p>

## Funktionen

- 🚀 **Hohe Leistung**: Basierend auf den Hochkonkurrenzmerkmale der Programmiersprache Go, in der Lage, eine große Anzahl von Dateianfragen effizient zu verarbeiten
- 🔒 **Berechtigungskontrolle**: Unterstützt grundlegende Authentifizierung und Dateizugriffsberechtigungen
- 📝 **Detaillierte Protokollierung**: Protokolliert Anfragen und Betriebsprotokolle für die Fehlerbehebung und Überwachung
- 📁 **Dateioperationen**: Unterstützt Dateiupload, -download, -löschung, -umbenennung usw.
- 📋 **Verzeichnisauflistung**: Generiert automatisch Verzeichnislisten für einfaches Browsen und Zugriff auf Dateien
- ⚙️ **Flexible Konfiguration**: Unterstützt Konfiguration über Konfigurationsdateien, Umgebungsvariablen und Befehlszeilenparameter
- 🐳 **Container-Unterstützung**: Bietet Docker-Images für einfache Bereitstellung und Ausführung
- 🚀 **Automatische Bereitstellung**: Unterstützt die automatische Bereitstellung auf Remote-Servern über GitHub Actions

## Schnellstart

### Installation

#### Mit Go installieren

```bash
go install github.com/3ziye/put-file@latest
```

#### Aus der Quelle kompilieren

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Vorcompilierte Binärdateien verwenden

put-file bietet vorcompilierte Binärdateien für Linux-, Windows- und Mac-Systeme, die direkt heruntergeladen und verwendet werden können.

1. Besuchen Sie die [GitHub-Releases-Seite](https://github.com/3ziye/put-file/releases) und laden Sie das komprimierte Paket der Binärdateien für Ihre Plattform herunter

2. Extrahieren Sie die entsprechenden Dateien entsprechend Ihrem Betriebssystem:

   **Linux:**
   ```bash
   # Linux-Version herunterladen und extrahieren
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Den Dienst starten
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Windows-Version herunterladen und extrahieren
   # Klicken Sie mit der rechten Maustaste auf die heruntergeladene Zip-Datei und wählen Sie "Hier extrahieren"
   
   # Den Dienst starten
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Mac-Version herunterladen und extrahieren
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Den Dienst starten
   ./put-file
   ```

3. Versionsinformationen prüfen
   Nach dem Start des Dienstes können Sie die aktuelle Versionsnummer in den Konsolenprotokollen sehen

#### Docker verwenden

Image von Docker Hub herunterladen:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Image von GitHub Package herunterladen:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker-Image

### Image von GitHub Package herunterladen

1. Stellen Sie sicher, dass Docker installiert ist
2. Führen Sie den folgenden Befehl aus, um das Image herunterzuladen:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Docker-Image ausführen

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Dieser Befehl startet den Container, ordnet Port 8080 des Containers Port 8080 des Hosts zu und mountet das Verzeichnis `./files` des Hosts in das Verzeichnis `/app/uploads` im Container.

### Image nach GitHub Package pushen (nur für Maintainer)

1. Melden Sie sich bei GitHub Container Registry an:
   ```bash
docker login ghcr.io -u IHR_GITHUB_BENUTZERNAME -p IHR_GITHUB_TOKEN
   ```

2. Erstellen Sie das Image:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Pushen Sie das Image:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Automatische Bereitstellung auf einem Server

put-file unterstützt die automatische Bereitstellung auf Remote-Servern über GitHub Actions. Für detaillierte Konfigurationsschritte lesen Sie bitte die [Bereitstellungsdokumentation](doc/DEPLOYMENT.md).

Hauptmerkmale der automatischen Bereitstellung:
- 🔑 Verwaltung von Serveranmeldeinformationen über GitHub Secrets
- 📥 Automatischer Download von Linux-Binärdateien
- 📁 Server-Sicherung und -bereitstellung
- 🚀 Unterstützung für systemd-Dienste
- ⚡ Auslösen der Bereitstellung über Releases oder manuelle Operationen

## Andere Sprachversionen

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
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
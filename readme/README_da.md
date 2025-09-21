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
 
 <p align="center">put-file er en højtydende, letvægts statisk filserver udviklet i programmeringssproget Go. Den understøtter grundlæggende operationer såsom upload, download og sletning af filer, og tilbyder også funktioner som adgangskontrol og logning.</p>

## Funktioner

- 🚀 **Høj ydelse**: Baseret på Gos høje koncurrenceegenskaber, i stand til effektivt at håndtere et stort antal filforespørgsler
- 🔒 **Adgangskontrol**: Understøtter grundlæggende autentificering og adgangskontrol for filer
- 📝 **Detaljeret logning**: Registrerer forespørgsler og operationslogger til fejlfinding og overvågning
- 📁 **Fileoperationer**: Understøtter upload, download, sletning, omdøbning af filer, osv.
- 📋 **Kataloglister**: Automatisk generering af kataloglister til nem browsing og filadgang
- ⚙️ **Fleksibel konfiguration**: Understøtter konfiguration via konfigurationsfiler, miljøvariabler og kommandolinjeparametre
- 🐳 **Containerunderstøttelse**: Leverer Docker-billeder til let deployment og drift
- 🚀 **Automatisk deployment**: Understøtter automatisk deployment til eksterne servere via GitHub Actions

## Hurtigstart

### Installation

#### Installer med Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Byg fra kildekode

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Brug forudkompilerede binærfiler

put-file leverer forudkompilerede binærfiler til Linux, Windows og Mac-systemer, som kan downloades direkte til brug.

1. Besøg [GitHub-releasesiden](https://github.com/3ziye/put-file/releases) og download den komprimerede pakke med binærfiler, der svarer til din platform

2. Udpak de tilsvarende filer i henhold til dit operativsystem:

   **Linux:**
   ```bash
   # Download og udpak Linux-versionen
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Kør tjenesten
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Download og udpak Windows-versionen
   # Højreklik på den downloadede zip-fil og vælg "Udpak her"
   
   # Kør tjenesten
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Download og udpak Mac-versionen
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Kør tjenesten
   ./put-file
   ```

3. Tjek versionsoplysninger
   Efter at have startet tjenesten kan du se det aktuelle versionsnummer i konsolloggerne

#### Brug Docker

Download billedet fra Docker Hub:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Download billedet fra GitHub Package:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker-billede

### Download billedet fra GitHub Package

1. Sørg for, at Docker er installeret
2. Kør følgende kommando for at downloade billedet:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Kør Docker-billedet

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Denne kommando kører containeren, mapper port 8080 på containeren til port 8080 på værten og monterer mappen `./files` på værten til mappen `/app/uploads` inde i containeren.

### Upload billedet til GitHub Package (kun til maintainere)

1. Log ind på GitHub Container Registry:
   ```bash
docker login ghcr.io -u GITHUB_BRUGERNAVN -p GITHUB_TOKEN
   ```

2. Byg billedet:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Upload billedet:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Automatisk deployment til servere

put-file understøtter automatisk deployment til eksterne servere via GitHub Actions. For detaljerede konfigurationstrin, se [deployeringsdokumentationen](doc/DEPLOYMENT.md).

Hovedfunktioner ved automatisk deployment:
- 🔑 Håndtering af serverautentiseringsoplysninger via GitHub Secrets
- 📥 Automatisk download af Linux-binærer
- 📁 Serverbackup og deployment
- 🚀 Understøttelse af systemd-tjenester
- ⚡ Triggering af deployment via udgivelser (Release) eller manuelle operationer

## Andre versioner på forskellige sprog

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
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
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
 
 <p align="center">put-file er en høyytelses, lettvekt statisk filserver utviklet i programmeringsspråket Go. Den støtter grunnleggende operasjoner som opplasting, nedlasting og sletting av filer, og tilbyr også funksjoner som tilgangskontroll og logging.</p>

## Funksjoner

- 🚀 **Høy ytelse**: Basert på Gos høye konkurrensegenskaper, i stand til å håndtere et stort antall filforespørsler effektivt
- 🔒 **Tilgangskontroll**: Støtter grunnleggende autentisering og tilgangskontroll for filer
- 📝 **Detaljert logging**: Registrerer forespørsler og operasjonslogger for feilsøking og overvåking
- 📁 **Fileoperasjoner**: Støtter opplasting, nedlasting, sletting, omdøping av filer, osv.
- 📋 **Kataloglister**: Automatisk generering av kataloglister for enkel browsing og filtilgang
- ⚙️ **Fleksibel konfigurasjon**: Støtter konfigurasjon via konfigurasjonsfiler, miljøvariabler og kommandolinjeparametere
- 🐳 **Containerstøtte**: Leverer Docker-avbildninger for enkel deployering og operasjon
- 🚀 **Automatisk deployering**: Støtter automatisk deployering til eksterne servere via GitHub Actions

## Hurtigstart

### Installasjon

#### Installer med Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Bygg fra kildekode

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Bruk forhåndskompilerte binærfiler

put-file leverer forhåndskompilerte binærfiler for Linux, Windows og Mac-systemer, som kan lastes ned direkte for bruk.

1. Besøk [GitHub-releasenside](https://github.com/3ziye/put-file/releases) og last ned komprimerte pakken med binærfiler som tilsvarer din plattform

2. Pakk ut de tilsvarende filene i henhold til ditt operativsystem:

   **Linux:**
   ```bash
   # Last ned og pakk ut Linux-versjonen
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Kjør tjenesten
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Last ned og pakk ut Windows-versjonen
   # Høyreklikk på den lastede zip-filen og velg "Pakk ut her"
   
   # Kjør tjenesten
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Last ned og pakk ut Mac-versjonen
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Kjør tjenesten
   ./put-file
   ```

3. Sjekk versjonsinformasjon
   Etter å ha startet tjenesten kan du se det gjeldende versjonsnummeret i konsoll-loggene

#### Bruk Docker

Last ned avbildningen fra Docker Hub:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Last ned avbildningen fra GitHub Package:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker-avbildning

### Last ned avbildningen fra GitHub Package

1. Sørg for at Docker er installert
2. Kjør følgende kommando for å laste ned avbildningen:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Kjør Docker-avbildningen

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Denne kommandoen kjører containeren, kobler port 8080 på containeren til port 8080 på verten og monterer katalogen `./files` på verten til katalogen `/app/uploads` inne i containeren.

### Last opp avbildningen til GitHub Package (kun for vedlikeholdere)

1. Logg inn på GitHub Container Registry:
   ```bash
docker login ghcr.io -u GITHUB_BRUKERNAVN -p GITHUB_TOKEN
   ```

2. Bygg avbildningen:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Last opp avbildningen:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Automatisk deployering til servere

put-file støtter automatisk deployering til eksterne servere via GitHub Actions. For detaljerte konfigurasjonssteg, se [deployeringsdokumentasjonen](doc/DEPLOYMENT.md).

Hovedfunksjoner til automatisk deployering:
- 🔑 Håndtering av serverautentiseringsuppgifter via GitHub Secrets
- 📥 Automatisk nedlasting av Linux-binærer
- 📁 Serverbackup og deployering
- 🚀 Støtte for systemd-tjenester
- ⚡ Utløsing av deployering via utgivelser (Release) eller manuelle operasjoner

## Andre versjoner på forskjellige språk

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
- [Dansk](README_da.md)
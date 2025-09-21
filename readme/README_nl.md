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
 
 <p align="center">put-file is een hoogwaardige en lichtgewicht statische bestands server ontwikkeld in de programmeertaal Go. Het ondersteunt basisbewerkingen zoals uploaden, downloaden en verwijderen van bestanden, en biedt ook functies zoals toegangsbeheer en logregistratie.</p>

## Kenmerken

- 🚀 **Hoge prestaties**: Gebaseerd op de hoge concurrency-kenmerken van de programmeertaal Go, in staat om efficiënt met een groot aantal bestandaanvragen om te gaan
- 🔒 **Toegangsbeheer**: Ondersteunt basisauthenticatie en toegangscontrole voor bestanden
- 📝 **Gedetailleerde logboekregistratie**: Registreert aanvragen en bewerkingslogboeken voor probleemoplossing en monitoring
- 📁 **Bestandsbewerkingen**: Ondersteunt uploaden, downloaden, verwijderen, hernoemen van bestanden, enz.
- 📋 **Maplijsten**: Genereert automatisch maplijsten om het navigeren en openen van bestanden te vergemakkelijken
- ⚙️ **Flexibele configuratie**: Ondersteunt configuratie via configuratiebestanden, omgevingsvariabelen en opdrachtregelparameters
- 🐳 **Containerondersteuning**: Biedt Docker-afbeeldingen voor eenvoudige implementatie en uitvoering
- 🚀 **Automatische implementatie**: Ondersteunt automatische implementatie op externe servers via GitHub Actions

## Snelstart

### Installatie

#### Installeren met Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Bouwen vanuit broncode

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Geprecompileerde binaire bestanden gebruiken

put-file biedt geprecompileerde binaire bestanden voor Linux, Windows en Mac-systemen, die rechtstreeks kunnen worden gedownload voor gebruik.

1. Bezoek de [GitHub-releasepagina](https://github.com/3ziye/put-file/releases) en download het bijpassende pakket met binaire bestanden voor uw platform

2. Pak de bijpassende bestanden uit afhankelijk van het besturingssysteem dat u gebruikt:

   **Linux:**
   ```bash
   # Download en pak de Linux-versie uit
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Start de service
   ./put-file
   ```
   
   **Windows:**
   ```powershell
   # Download en pak de Windows-versie uit
   # Klik met de rechtermuisknop op het gedownloade zip-bestand en selecteer "Uitpakken hier"
   
   # Start de service
   .\put-file.exe
   ```
   
   **Mac:**
   ```bash
   # Download en pak de Mac-versie uit
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Start de service
   ./put-file
   ```

3. Controleer versie-informatie
   Na het starten van de service kunt u het huidige versienummer zien in de consolelogboeken

#### Docker gebruiken

Download de afbeelding van Docker Hub:
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Download de afbeelding van GitHub Package:
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## GitHub Package Docker-afbeelding

### Download de afbeelding van GitHub Package

1. Zorg ervoor dat Docker is geïnstalleerd
2. Voer de volgende opdracht uit om de afbeelding te downloaden:
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Voer de Docker-afbeelding uit

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Met deze opdracht wordt de container uitgevoerd, wordt poort 8080 van de container toegewezen aan poort 8080 van de host en wordt de map `./files` van de host gekoppeld aan de map `/app/uploads` in de container.

### Upload de afbeelding naar GitHub Package (alleen voor beheerders)

1. Meld u aan bij GitHub Container Registry:
   ```bash
docker login ghcr.io -u GITHUB_GEBRUIKERSNAAM -p GITHUB_TOKEN
   ```

2. Bouw de afbeelding:
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Upload de afbeelding:
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Automatische implementatie op servers

put-file ondersteunt automatische implementatie op externe servers via GitHub Actions. Raadpleeg de [implementatiedocumentatie](doc/DEPLOYMENT.md) voor gedetailleerde configuratiestappen.

Belangrijke functies van automatische implementatie:
- 🔑 Beheer van serverreferenties via GitHub Secrets
- 📥 Automatisch downloaden van Linux-binaire bestanden
- 📁 Serverback-up en implementatie
- 🚀 Ondersteuning voor systemd-services
- ⚡ Triggeren van implementaties via releases of handmatige bewerkingen

## Andere versies in verschillende talen

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
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
<h1 align="center" style="border-bottom: none"> 
     <a href="" target="_blank"> 
         <alt="GoStaticServe" src="" width="100" height="100"> 
     </a> 
     <br>GoStaticServe 
 </h1> 
 
 <div align="center" style="line-height: 2;"> 
   [<a href="/README.md">English</a>] | [<a href="/readme/README_ar.md">العربية</a>] | [<a href="/readme/README_da.md">Dansk</a>] | [<a href="/readme/README_de.md">Deutsch</a>] | [<a href="/readme/README_es.md">Español</a>] | [<a href="/readme/README_fr.md">Français</a>] | [<a href="/readme/README_it.md">Italiano</a>] | [<a href="/readme/README_ja.md">日本語</a>] | [<a href="/readme/README_ko.md">한국어</a>] | [<a href="/readme/README_nl.md">Nederlands</a>] | [<a href="/readme/README_no.md">Norsk</a>] | [<a href="/readme/README_pl.md">Polski</a>] | [<a href="/readme/README_pt.md">Português</a>] | [<a href="/readme/README_ru.md">Русский</a>] | [<a href="/readme/README_sv.md">Svenska</a>] | [<a href="/readme/README_th.md">ไทย</a>] | [<a href="/readme/README_vi.md">Tiếng Việt</a>] | [<a href="/readme/README_zh.md">中文(简体)</a>] 
   <br> 
   
   | ** [Issues](https://github.com/3ziye/GoStaticServe/issues) ** | ** [Releases](https://github.com/3ziye/GoStaticServe/releases) ** | ** [README](https://github.com/3ziye/GoStaticServe/blob/main/README.md) ** | ** [Architecture](https://github.com/3ziye/GoStaticServe/blob/main/doc/architecture.md) ** | 
   <br> 
   
   [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) 
   &nbsp;&nbsp; 
   ![Go](https://img.shields.io/badge/language-Go-blue.svg) 
   &nbsp;&nbsp; 
   ![performance](https://img.shields.io/badge/performance-high-yellow.svg) 
   &nbsp;&nbsp; 
   ![status](https://img.shields.io/badge/status-Stable-green.svg) 
 </div> 
 
 <p align="center">GoStaticServe är en högpresterande och lättvikts statisk filserver utvecklad i programmeringsspråket Go. Den stödjer grundläggande operationer som uppladdning, nedladdning och borttagning av filer, och erbjuder också funktioner som behörighetskontroll och loggning.</p>

## Funktioner

- 🚀 **Hög prestanda**: Baserad på Go-språkets höga koncurrency-funktioner, kan hantera ett stort antal filbegäranden effektivt
- 🔒 **Behörighetskontroll**: Stödjer grundläggande autentisering och åtkomstkontroll för filer
- 📝 **Detaljerad loggning**: Registrerar förfrågningar och operationsloggar för felsökning och övervakning
- 📁 **Filoperationer**: Stödjer uppladdning, nedladdning, borttagning, omdöpning av filer, etc.
- 📋 **Kataloglistor**: Genererar automatiskt kataloglistor för att underlätta navigering och åtkomst till filer
- ⚙️ **Flexibel konfiguration**: Stödjer konfiguration via konfigurationsfiler, miljövariabler och kommandoradsparametrar
- 🐳 **Containerstöd**: Tillhandahåller Docker-avbildningar för enkel distribution och exekvering
- 🚀 **Automatisk distribution**: Stödjer automatisk distribution till fjärrserverar via GitHub Actions

## Snabbstart

### Installation

#### Installera med Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Bygg från källkod

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Använd förkompilerade binärfiler

GoStaticServe tillhandahåller förkompilerade binärfiler för Linux, Windows och Mac-system, som kan laddas ner direkt för användning.

1. Besök [GitHub-releasidan](https://github.com/3ziye/GoStaticServe/releases) och ladda ner det motsvarande paketet med binärfiler för din plattform

2. Packa upp de relevanta filerna beroende på vilket operativsystem du använder:

   **Linux:**
   ```bash
   # Ladda ner och packa upp Linux-versionen
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Starta tjänsten
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Ladda ner och packa upp Windows-versionen
   # Högerklicka på den nedladdade zip-filen och välj "Packa upp här"
   
   # Starta tjänsten
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Ladda ner och packa upp Mac-versionen
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Starta tjänsten
   ./GoStaticServe
   ```

3. Kontrollera versionsinformation
   Efter att ha startat tjänsten kan du se det aktuella versionsnumret i konsolloggarna

#### Använd Docker

Ladda ner avbildningen från Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Ladda ner avbildningen från GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## GitHub Package Docker-avbildning

### Ladda ner avbildningen från GitHub Package

1. Kontrollera att Docker är installerat
2. Kör följande kommando för att ladda ner avbildningen:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Kör Docker-avbildningen

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Detta kommando kör containern, mappar port 8080 på containern till port 8080 på värden och monterar katalogen `./files` på värden till katalogen `/app/uploads` inuti containern.

### Ladda upp avbildningen till GitHub Package (endast för underhållare)

1. Logga in på GitHub Container Registry:
   ```bash
docker login ghcr.io -u GITHUB_ANVÄNDARNAMN -p GITHUB_TOKEN
   ```

2. Bygg avbildningen:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Ladda upp avbildningen:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Automatisk distribution till servrar

GoStaticServe stödjer automatisk distribution till fjärrserverar via GitHub Actions. För detaljerade konfigurationssteg, se [distributionsdokumentationen](doc/DEPLOYMENT.md).

Viktiga funktioner för automatisk distribution:
- 🔑 Hantering av serverautentiseringsuppgifter via GitHub Secrets
- 📥 Automatisk nedladdning av Linux-binärer
- 📁 Serverbackup och distribution
- 🚀 Stöd för systemd-tjänster
- ⚡ Utlös distribution via släpp (Release) eller manuella åtgärder

## Andra versioner på olika språk

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
- [Norsk](README_no.md)
- [Dansk](README_da.md)
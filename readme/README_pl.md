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
 
 <p align="center">GoStaticServe to wydajny i lekki serwer plików statycznych rozwijany w języku Go. Obsługuje podstawowe operacje takie jak wysyłanie, pobieranie i usuwanie plików, a także zapewnia funkcje takie jak kontrola dostępu i rejestrowanie dzienników.</p>

## Cechy

- 🚀 **Wysoka wydajność**: Opiera się na wysokiej współbieżności języka Go, umożliwia efektywne zarządzanie dużą liczbą żądań plików
- 🔒 **Kontrola dostępu**: Obsługuje podstawową uwierzytelnianie i kontrolę dostępu do plików
- 📝 **Szczegółowe logowanie**: Rejestruje żądania i dzienniki operacji w celu rozwiązywania problemów i monitorowania
- 📁 **Operacje na plikach**: Obsługuje przesyłanie, pobieranie, usuwanie, zmianę nazwy plików itp.
- 📋 **Listy katalogów**: Automatycznie generuje listy katalogów ułatwiając przeglądanie i dostęp do plików
- ⚙️ **Elastyczna konfiguracja**: Obsługuje konfigurację poprzez pliki konfiguracyjne, zmienne środowiskowe i parametry wiersza poleceń
- 🐳 **Obsługa kontenerów**: Dostarcza obrazy Dockera ułatwiające wdrażanie i uruchamianie
- 🚀 **Automatyczne wdrażanie**: Obsługuje automatyczne wdrażanie na zdalnych serwerach poprzez GitHub Actions

## Szybki start

### Instalacja

#### Instalacja przy użyciu Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Budowanie z kodu źródłowego

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Używanie wstępnie skompilowanych plików binarnych

GoStaticServe udostępnia wstępnie skompilowane pliki binarne dla systemów Linux, Windows i Mac, które można bezpośrednio pobrać do użycia.

1. Przejdź do [strony wydań GitHub](https://github.com/3ziye/GoStaticServe/releases) i pobierz pakiet plików binarnych odpowiedni dla Twojej platformy

2. Rozpakuj odpowiednie pliki w zależności od używanego systemu operacyjnego:

   **Linux:**
   ```bash
   # Pobierz i rozpakuj wersję dla Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Uruchom usługę
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Pobierz i rozpakuj wersję dla Windows
   # Kliknij prawym przyciskiem myszy pobrany plik zip i wybierz "Rozpakuj tutaj"
   
   # Uruchom usługę
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Pobierz i rozpakuj wersję dla Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Uruchom usługę
   ./GoStaticServe
   ```

3. Sprawdź informacje o wersji
   Po uruchomieniu usługi można zobaczyć aktualny numer wersji w dziennikach konsoli

#### Używanie Dockera

Pobierz obraz z Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Pobierz obraz z GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## Obraz Docker z GitHub Package

### Pobierz obraz z GitHub Package

1. Upewnij się, że Docker jest zainstalowany
2. Wykonaj następującą komendę, aby pobrać obraz:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Uruchom obraz Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Ta komenda uruchomi kontener, zamapuje port 8080 kontenera na port 8080 hosta i zamontuje katalog `./files` hosta do katalogu `/app/uploads` wewnątrz kontenera.

### Wrzuć obraz do GitHub Package (tylko dla maintainerów)

1. Zaloguj się do GitHub Container Registry:
   ```bash
docker login ghcr.io -u NAZWA_UŻYTKOWNIKA_GITHUB -p TOKEN_GITHUB
   ```

2. Zbuduj obraz:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Wrzuć obraz:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Automatyczne wdrażanie na serwerach

GoStaticServe obsługuje automatyczne wdrażanie na zdalnych serwerach poprzez GitHub Actions. Aby uzyskać szczegółowe kroki konfiguracji, zapoznaj się z [dokumentacją wdrażania](doc/DEPLOYMENT.md).

Główne funkcje automatycznego wdrażania:
- 🔑 Zarządzanie poświadczeniami serwera poprzez GitHub Secrets
- 📥 Automatyczne pobieranie binariów Linux
- 📁 Kopia zapasowa i wdrażanie na serwerze
- 🚀 Wsparcie dla usług systemd
- ⚡ Wyzwalanie wdrażania poprzez wydania (Release) lub operacje ręczne

## Inne wersje w różnych językach

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
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
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
 
 <p align="center">GoStaticServe è un server di file statici performante e leggero sviluppato nel linguaggio Go. Supporta operazioni di base come upload, download ed eliminazione di file, e fornisce anche funzionalità come controllo degli accessi e registrazione dei log.</p>

## Caratteristiche

- 🚀 **Alta performance**: Basato sulle caratteristiche di elevata concorrenza del linguaggio Go, in grado di gestire efficientemente un gran numero di richieste di file
- 🔒 **Controllo degli accessi**: Supporta l'autenticazione di base e il controllo degli accessi ai file
- 📝 **Registrazione dettagliata**: Registra le richieste e i log delle operazioni per la risoluzione dei problemi e il monitoraggio
- 📁 **Operazioni sui file**: Supporta upload, download, eliminazione, rinominazione dei file, ecc.
- 📋 **Elenco delle directory**: Genera automaticamente elenchi di directory per facilitare la navigazione e l'accesso ai file
- ⚙️ **Configurazione flessibile**: Supporta la configurazione tramite file di configurazione, variabili d'ambiente e parametri della riga di comando
- 🐳 **Supporto per i container**: Fornisce immagini Docker per semplificare la distribuzione e l'esecuzione
- 🚀 **Distribuzione automatica**: Supporta la distribuzione automatica su server remoti tramite GitHub Actions

## Avvio rapido

### Installazione

#### Installare usando Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Compilare dal codice sorgente

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Usare i file binari precompilati

GoStaticServe fornisce file binari precompilati per i sistemi Linux, Windows e Mac, che possono essere scaricati direttamente per l'uso.

1. Visitare la [pagina delle release di GitHub](https://github.com/3ziye/GoStaticServe/releases) e scaricare il pacchetto compresso dei file binari corrispondente alla propria piattaforma

2. Estrarre i file corrispondenti in base al sistema operativo in uso:

   **Linux:**
   ```bash
   # Scaricare ed estrarre la versione per Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Eseguire il servizio
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Scaricare ed estrarre la versione per Windows
   # Fare clic con il tasto destro sul file zip scaricato e selezionare "Estrai qui"
   
   # Eseguire il servizio
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Scaricare ed estrarre la versione per Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Eseguire il servizio
   ./GoStaticServe
   ```

3. Verificare le informazioni sulla versione
   Dopo aver avviato il servizio, è possibile vedere il numero di versione attuale nei log della console

#### Usare Docker

Scaricare l'immagine da Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Scaricare l'immagine da GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## Immagine Docker di GitHub Package

### Scaricare l'immagine da GitHub Package

1. Assicurarsi che Docker sia installato
2. Eseguire il seguente comando per scaricare l'immagine:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Eseguire l'immagine Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Questo comando eseguirà il container, mappando la porta 8080 del container alla porta 8080 dell'host e montando la directory `./files` dell'host nella directory `/app/uploads` all'interno del container.

### Caricare l'immagine su GitHub Package (solo per i maintainer)

1. Accedere a GitHub Container Registry:
   ```bash
docker login ghcr.io -u TUO_NOME_UTENTE_GITHUB -p TOKEN_GITHUB
   ```

2. Costruire l'immagine:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Caricare l'immagine:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Distribuzione automatica su server

GoStaticServe supporta la distribuzione automatica su server remoti tramite GitHub Actions. Per i passaggi dettagliati di configurazione, consultare la [documentazione sulla distribuzione](doc/DEPLOYMENT.md).

Caratteristiche principali della distribuzione automatica:
- 🔑 Gestione delle credenziali del server tramite GitHub Secrets
- 📥 Download automatico dei binari Linux
- 📁 Backup e distribuzione sul server
- 🚀 Supporto per i servizi systemd
- ⚡ Attivazione della distribuzione tramite release o operazioni manuali

## Altre versioni in diverse lingue

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
- [Português](README_pt.md)
- [한국어](README_ko.md)
- [العربية](README_ar.md)
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
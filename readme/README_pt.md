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
 
 <p align="center">GoStaticServe é um servidor de arquivos estáticos de alto desempenho e leve, desenvolvido na linguagem Go. Ele suporta operações básicas como upload, download e exclusão de arquivos, e também fornece recursos como controle de permissões e registro de logs.</p>

## Recursos

- 🚀 **Alto desempenho**: Baseado nas características de alta concorrência da linguagem Go, capaz de lidar eficientemente com um grande número de solicitações de arquivos
- 🔒 **Controle de permissões**: Suporta autenticação básica e controle de permissões de acesso a arquivos
- 📝 **Registro detalhado**: Registra solicitações e logs de operações para solução de problemas e monitoramento
- 📁 **Operações com arquivos**: Suporta upload, download, exclusão, renomeação de arquivos, etc.
- 📋 **Listagem de diretórios**: Gera automaticamente listagens de diretórios para facilitar a navegação e acesso aos arquivos
- ⚙️ **Configuração flexível**: Suporta configuração por meio de arquivos de configuração, variáveis de ambiente e parâmetros de linha de comando
- 🐳 **Suporte a contêineres**: Fornece imagens Docker para facilitar a implantação e execução
- 🚀 **Implantação automática**: Suporta implantação automática em servidores remotos por meio do GitHub Actions

## Início rápido

### Instalação

#### Instalar usando o Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Construir a partir do código fonte

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Usar arquivos binários pré-compilados

O GoStaticServe fornece arquivos binários pré-compilados para sistemas Linux, Windows e Mac, que podem ser baixados diretamente para uso.

1. Acesse a [página de lançamentos do GitHub](https://github.com/3ziye/GoStaticServe/releases) e baixe o pacote compactado de arquivos binários correspondente à sua plataforma

2. Extraia os arquivos correspondentes de acordo com seu sistema operacional:

   **Linux:**
   ```bash
   # Baixar e extrair a versão para Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Executar o serviço
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Baixar e extrair a versão para Windows
   # Clique com o botão direito no arquivo zip baixado e selecione "Extrair aqui"
   
   # Executar o serviço
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Baixar e extrair a versão para Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Executar o serviço
   ./GoStaticServe
   ```

3. Verificar informações de versão
   Após iniciar o serviço, você pode ver o número da versão atual nos logs do console

#### Usar o Docker

Baixar a imagem do Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Baixar a imagem do GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## Imagem Docker do GitHub Package

### Baixar a imagem do GitHub Package

1. Certifique-se de ter o Docker instalado
2. Execute o seguinte comando para baixar a imagem:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Executar a imagem Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Este comando executará o contêiner, mapeará a porta 8080 do contêiner para a porta 8080 do host e montará o diretório `./files` do host no diretório `/app/uploads` dentro do contêiner.

### Enviar a imagem para o GitHub Package (apenas para mantenedores)

1. Faça login no GitHub Container Registry:
   ```bash
docker login ghcr.io -u SEU_NOME_DE_USUARIO_GITHUB -p SEU_TOKEN_GITHUB
   ```

2. Construa a imagem:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Envie a imagem:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Implantação automática em servidores

O GoStaticServe suporta implantação automática em servidores remotos por meio do GitHub Actions. Para obter etapas detalhadas de configuração, consulte a [documentação de implantação](doc/DEPLOYMENT.md).

Recursos principais da implantação automática:
- 🔑 Gerenciamento de credenciais do servidor por meio do GitHub Secrets
- 📥 Download automático de binários do Linux
- 📁 Backup e implantação no servidor
- 🚀 Suporte a serviços systemd
- ⚡ Disparo da implantação por meio de lançamentos (Release) ou operações manuais

## Outras versões em diferentes idiomas

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
- [Русский](README_ru.md)
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
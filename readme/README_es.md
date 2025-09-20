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
 
 <p align="center">GoStaticServe es un servidor de archivos estáticos de alto rendimiento y ligero desarrollado en el lenguaje Go. Admite operaciones básicas como subida, descarga y eliminación de archivos, y también proporciona funciones como control de permisos y registro de logs.</p>

## Características

- 🚀 **Alto rendimiento**: Basado en las características de alta concurrencia del lenguaje Go, capaz de manejar eficientemente una gran cantidad de solicitudes de archivos
- 🔒 **Control de permisos**: Admite autenticación básica y control de permisos de acceso a archivos
- 📝 **Registro detallado**: Registra solicitudes y logs de operaciones para la resolución de problemas y la monitorización
- 📁 **Operaciones con archivos**: Admite subida, descarga, eliminación, cambio de nombre de archivos, etc.
- 📋 **Listado de directorios**: Genera automáticamente listados de directorios para facilitar la navegación y el acceso a los archivos
- ⚙️ **Configuración flexible**: Admite configuración a través de archivos de configuración, variables de entorno y parámetros de línea de comandos
- 🐳 **Soporte para contenedores**: Proporciona imágenes Docker para facilitar el despliegue y la ejecución
- 🚀 **Despliegue automático**: Admite el despliegue automático en servidores remotos a través de GitHub Actions

## Inicio rápido

### Instalación

#### Instalar usando Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Construir desde el código fuente

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Usar archivos binarios precompilados

GoStaticServe proporciona archivos binarios precompilados para sistemas Linux, Windows y Mac, que se pueden descargar directamente para su uso.

1. Visite la [página de lanzamientos de GitHub](https://github.com/3ziye/GoStaticServe/releases) y descargue el paquete comprimido de archivos binarios correspondiente a su plataforma

2. Extraiga los archivos correspondientes según su sistema operativo:

   **Linux:**
   ```bash
   # Descargar y extraer la versión para Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Ejecutar el servicio
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Descargar y extraer la versión para Windows
   # Haga clic derecho en el archivo zip descargado y seleccione "Extraer aquí"
   
   # Ejecutar el servicio
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Descargar y extraer la versión para Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Ejecutar el servicio
   ./GoStaticServe
   ```

3. Consultar información de versión
   Después de iniciar el servicio, puede ver el número de versión actual en los registros de la consola

#### Usar Docker

Descargar la imagen desde Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Descargar la imagen desde GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## Imagen Docker de GitHub Package

### Descargar la imagen desde GitHub Package

1. Asegúrese de tener Docker instalado
2. Ejecute el siguiente comando para descargar la imagen:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Ejecutar la imagen Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Este comando ejecutará el contenedor, mapeará el puerto 8080 del contenedor al puerto 8080 del host y montará el directorio `./files` del host en el directorio `/app/uploads` dentro del contenedor.

### Subir la imagen a GitHub Package (solo para mantenedores)

1. Inicie sesión en GitHub Container Registry:
   ```bash
docker login ghcr.io -u SU_NOMBRE_DE_USUARIO_DE_GITHUB -p SU_TOKEN_DE_GITHUB
   ```

2. Construya la imagen:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Suba la imagen:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Despliegue automático en servidores

GoStaticServe admite el despliegue automático en servidores remotos a través de GitHub Actions. Para conocer los pasos de configuración detallados, consulte la [documentación de despliegue](doc/DEPLOYMENT.md).

Características principales del despliegue automático:
- 🔑 Gestión de credenciales de servidor a través de GitHub Secrets
- 📥 Descarga automática de binarios de Linux
- 📁 Copia de seguridad y despliegue en el servidor
- 🚀 Soporte para servicios systemd
- ⚡ Desencadenamiento del despliegue a través de lanzamientos (Release) o operaciones manuales

## Otras versiones en diferentes idiomas

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
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
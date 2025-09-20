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
 
 <p align="center">GoStaticServe — это высокопроизводительный, легковесный сервер статических файлов, разработанный на языке Go. Он поддерживает базовые операции, такие как загрузка, скачивание и удаление файлов, а также предоставляет такие функции, как контроль прав доступа и ведение журналов.</p>

## Функции

- 🚀 **Высокая производительность**: Основан на особенностях высокой конкурентности языка Go, способен эффективно обрабатывать большое количество запросов на файлы
- 🔒 **Контроль прав доступа**: Поддерживает базовую аутентификацию и контроль прав доступа к файлам
- 📝 **Детальное журналирование**: Записывает запросы и журналы операций для устранения неполадок и мониторинга
- 📁 **Операции с файлами**: Поддерживает загрузку, скачивание, удаление, переименование файлов и т.д.
- 📋 **Перечень каталогов**: Автоматически генерирует списки каталогов для удобного просмотра и доступа к файлам
- ⚙️ **Гибкая настройка**: Поддерживает настройку через файлы конфигурации, переменные окружения и параметры командной строки
- 🐳 **Поддержка контейнеров**: Предоставляет образы Docker для удобного развертывания и запуска
- 🚀 **Автоматическое развертывание**: Поддерживает автоматическое развертывание на удаленных серверах через GitHub Actions

## Быстрый старт

### Установка

#### Установка с помощью Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### Сборка из исходного кода

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### Использование предварительно скомпилированных двоичных файлов

GoStaticServe предоставляет предварительно скомпилированные двоичные файлы для систем Linux, Windows и Mac, которые можно напрямую загрузить и использовать.

1. Посетите страницу [GitHub Releases](https://github.com/3ziye/GoStaticServe/releases) и скачайте сжатый пакет двоичных файлов, соответствующих вашей платформе

2. Распакуйте соответствующие файлы в зависимости от вашей операционной системы:

   **Linux:**
   ```bash
   # Скачайте и распакуйте версию для Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # Запустите службу
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # Скачайте и распакуйте версию для Windows
   # Щелкните правой кнопкой мыши по загруженному zip-файлу и выберите "Извлечь сюда"
   
   # Запустите службу
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # Скачайте и распакуйте версию для Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # Запустите службу
   ./GoStaticServe
   ```

3. Проверка информации о версии
   После запуска службы вы можете увидеть текущий номер версии в журналах консоли

#### Использование Docker

Загрузка образа из Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

Загрузка образа из GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## Образ Docker из GitHub Package

### Загрузка образа из GitHub Package

1. Убедитесь, что Docker установлен
2. Выполните следующую команду для загрузки образа:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### Запуск образа Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

Эта команда запустит контейнер, сопоставит порт 8080 контейнера с портом 8080 хоста и смонтирует каталог `./files` хоста в каталог `/app/uploads` внутри контейнера.

### Отправка образа в GitHub Package (только для сопровождающих)

1. Войдите в GitHub Container Registry:
   ```bash
docker login ghcr.io -u ВАШЕ_ИМЯ_ПОЛЬЗОВАТЕЛЯ_GITHUB -p ВАШ_ТОКЕН_GITHUB
   ```

2. Соберите образ:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. Отправьте образ:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## Автоматическое развертывание на сервере

GoStaticServe поддерживает автоматическое развертывание на удаленных серверах через GitHub Actions. Для получения подробных инструкций по настройке обратитесь к [документации по развертыванию](doc/DEPLOYMENT.md).

Основные функции автоматического развертывания:
- 🔑 Управление учетными данными сервера через GitHub Secrets
- 📥 Автоматическая загрузка двоичных файлов Linux
- 📁 Резервное копирование и развертывание на сервере
- 🚀 Поддержка служб systemd
- ⚡ Запуск развертывания через выпуски (Release) или вручную

## Другие языковые версии

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
- [Français](README_fr.md)
- [Deutsch](README_de.md)
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
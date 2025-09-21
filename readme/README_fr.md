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
 
 <p align="center">put-file est un serveur de fichiers statiques performant et léger développé en langage Go. Il prend en charge des opérations de base telles que le téléchargement, le téléchargement et la suppression de fichiers, et fournit également des fonctionnalités telles que le contrôle d'accès et la journalisation.</p>

## Fonctionnalités

- 🚀 **Haute performance** : Basé sur les caractéristiques de haute concurrence du langage Go, capable de gérer efficacement un grand nombre de requêtes de fichiers
- 🔒 **Contrôle d'accès** : Prend en charge l'authentification de base et le contrôle d'accès aux fichiers
- 📝 **Journalisation détaillée** : Enregistre les requêtes et les journaux d'opérations pour le dépannage et la surveillance
- 📁 **Opérations sur les fichiers** : Prend en charge le téléchargement, le téléchargement, la suppression, le renommage des fichiers, etc.
- 📋 **Liste de répertoires** : Génère automatiquement des listes de répertoires pour faciliter la navigation et l'accès aux fichiers
- ⚙️ **Configuration flexible** : Prend en charge la configuration via des fichiers de configuration, des variables d'environnement et des paramètres de ligne de commande
- 🐳 **Support pour conteneurs** : Fournit des images Docker pour faciliter le déploiement et l'exécution
- 🚀 **Déploiement automatique** : Prend en charge le déploiement automatique sur des serveurs distants via GitHub Actions

## Démarrage rapide

### Installation

#### Installer avec Go

```bash
go install github.com/3ziye/put-file@latest
```

#### Construire à partir du code source

```bash
git clone https://github.com/3ziye/put-file.git
cd put-file
go mod init github.com/3ziye/put-file
go build -o put-file cmd/server/main.go
```

#### Utiliser des fichiers binaires précompilés

put-file fournit des fichiers binaires précompilés pour les systèmes Linux, Windows et Mac, qui peuvent être téléchargés directement pour une utilisation.

1. Visitez la [page des versions GitHub](https://github.com/3ziye/put-file/releases) et téléchargez le package compressé de fichiers binaires correspondant à votre plateforme

2. Extraire les fichiers correspondants en fonction de votre système d'exploitation :

   **Linux :**
   ```bash
   # Télécharger et extraire la version Linux
   wget https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_linux_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_linux_amd64.tar.gz
chmod +x put-file
   
   # Exécuter le service
   ./put-file
   ```
   
   **Windows :**
   ```powershell
   # Télécharger et extraire la version Windows
   # Cliquez avec le bouton droit sur le fichier zip téléchargé et sélectionnez "Extraire ici"
   
   # Exécuter le service
   .\put-file.exe
   ```
   
   **Mac :**
   ```bash
   # Télécharger et extraire la version Mac
   curl -OL https://github.com/3ziye/put-file/releases/download/vX.Y.Z/put-file_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf put-file_vX.Y.Z_darwin_amd64.tar.gz
chmod +x put-file
   
   # Exécuter le service
   ./put-file
   ```

3. Vérifier les informations de version
   Après avoir démarré le service, vous pouvez voir le numéro de version actuel dans les journaux de la console

#### Utiliser Docker

Télécharger l'image depuis Docker Hub :
```bash
docker pull 3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/put-file:latest
```

Télécharger l'image depuis GitHub Package :
```bash
docker pull ghcr.io/3ziye/put-file:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

## Image Docker de GitHub Package

### Télécharger l'image depuis GitHub Package

1. Assurez-vous d'avoir installé Docker
2. Exécutez la commande suivante pour télécharger l'image :
   ```bash
docker pull ghcr.io/3ziye/put-file:latest
   ```

### Exécuter l'image Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/put-file:latest
```

Cette commande exécutera le conteneur, mappera le port 8080 du conteneur au port 8080 de l'hôte et montera le répertoire `./files` de l'hôte dans le répertoire `/app/uploads` du conteneur.

### Pousser l'image vers GitHub Package (pour les mainteneurs uniquement)

1. Connectez-vous à GitHub Container Registry :
   ```bash
docker login ghcr.io -u VOTRE_NOM_D'UTILISATEUR_GITHUB -p VOTRE_TOKEN_GITHUB
   ```

2. Construisez l'image :
   ```bash
docker build -t ghcr.io/3ziye/put-file:latest .
   ```

3. Poussez l'image :
   ```bash
docker push ghcr.io/3ziye/put-file:latest
   ```

## Déploiement automatique sur un serveur

put-file prend en charge le déploiement automatique sur des serveurs distants via GitHub Actions. Pour connaître les étapes de configuration détaillées, consultez la [documentation de déploiement](doc/DEPLOYMENT.md).

Fonctionnalités principales du déploiement automatique :
- 🔑 Gestion des informations d'identification du serveur via GitHub Secrets
- 📥 Téléchargement automatique des binaires Linux
- 📁 Sauvegarde et déploiement sur le serveur
- 🚀 Support pour les services systemd
- ⚡ Déclenchement du déploiement via des versions (Release) ou des opérations manuelles

## Autres versions dans différentes langues

- [English](README_en.md)
- [中文](README_zh.md)
- [日本語](README_ja.md)
- [Español](README_es.md)
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
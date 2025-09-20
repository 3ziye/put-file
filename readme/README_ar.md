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
 
 <p align="center">GoStaticServe هو خادم ملفات ثابتة عالي الأداء وخفيف الوزن تم تطويره بلغة Go. يدعم عمليات أساسية مثل تحميل الملفات، تنزيلها وحذفها، ويقدم أيضًا ميزات مثل التحكم في الأذونات وتسجيل السجلات.</p>

## الميزات

- 🚀 **أداء عالي**: يعتمد على خصائص Go للمتزامنات العالية، قادر على معالجة عدد كبير من طلبات الملفات بكفاءة
- 🔒 **التحكم في الأذونات**: يدعم المصادقة الأساسية ومراقبة وصول الملفات
- 📝 **تسجيل مفصل**: يسجل الطلبات والسجلات الخاصة بالعمليات لأغراض استكشاف الأخطاء والمراقبة
- 📁 **عمليات على الملفات**: يدعم تحميل الملفات، تنزيلها، حذفها، إعادة تسميتها، إلخ
- 📋 **قائمة الدلائل**: يولد تلقائيًا قوائم الدلائل لتسهيل التنقل والوصول إلى الملفات
- ⚙️ **تكوين مرن**: يدعم التكوين من خلال ملفات التكوين، متغيرات البيئة ووسائط سطر الأوامر
- 🐳 **دعم الحاويات**: يوفر صور Docker لتسهيل النشر والتشغيل
- 🚀 **النشر التلقائي**: يدعم النشر التلقائي على خوادم بعيدة عبر GitHub Actions

## بداية سريعة

### التثبيت

#### التثبيت باستخدام Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### البناء من الكود المصدري

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### استخدام ملفات ثنائية معدة مسبقًا

يوفر GoStaticServe ملفات ثنائية معدة مسبقًا لنظم Linux وWindows وMac، يمكن تنزيلها مباشرة للاستخدام.

1. زيارة [صفحة الإصدارات على GitHub](https://github.com/3ziye/GoStaticServe/releases) وتنزيل حزمة ملفات الثنائية المناسبة لنظام التشغيل الخاص بك

2. استخراج الملفات المناسبة بناءً لنظام التشغيل المستخدم:

   **Linux:**
   ```bash
   # تنزيل واستخراج الإصدار الخاص بـ Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # تشغيل الخدمة
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # تنزيل واستخراج الإصدار الخاص بـ Windows
   # انقر بزر الماوس الأيمن على ملف zip المحمل واختر "استخراج هنا"
   
   # تشغيل الخدمة
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # تنزيل واستخراج الإصدار الخاص بـ Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # تشغيل الخدمة
   ./GoStaticServe
   ```

3. التحقق من معلومات الإصدار
   بعد بدء الخدمة، يمكنك رؤية رقم الإصدار الحالي في سجلات وحدة التحكم

#### استخدام Docker

تنزيل الصورة من Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

تنزيل الصورة من GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## صورة Docker الخاصة بـ GitHub Package

### تنزيل الصورة من GitHub Package

1. تأكد من تثبيت Docker
2. قم بتشغيل الأمر التالي لتنزيل الصورة:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### تشغيل صورة Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

يقوم هذا الأمر بتشغيل الحاوية، وتعيين منفذ 8080 للحاوية على منفذ 8080 للخادم، وتصنيف الدليل `./files` للخادم في الدليل `/app/uploads` داخل الحاوية.

### تحميل الصورة إلى GitHub Package (للمشرفين فقط)

1. تسجيل الدخول إلى GitHub Container Registry:
   ```bash
docker login ghcr.io -u اسم_المستخدم_الخاص بـ GitHub -p رمز_الوصول_الخاص بـ GitHub
   ```

2. بناء الصورة:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. تحميل الصورة:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## النشر التلقائي على الخوادم

يدعم GoStaticServe النشر التلقائي على خوادم بعيدة عبر GitHub Actions. للحصول على خطوات التكوين التفصيلية، يُرجى الاطلاع على [وثائق النشر](doc/DEPLOYMENT.md).

الميزات الرئيسية للنشر التلقائي:
- 🔑 إدارة بيانات اعتماد الخادم عبر GitHub Secrets
- 📥 تنزيل تلقائي للملفات الثنائية الخاصة بـ Linux
- 📁 النسخ الاحتياطي والنشر على الخادم
- 🚀 دعم خدمات systemd
- ⚡ تشغيل النشر عبر الإصدارات (Release) أو العمليات اليدوية

## إصدارات أخرى بلغات مختلفة

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
- [Tiếng Việt](README_vi.md)
- [ไทย](README_th.md)
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
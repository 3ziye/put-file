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
 
 <p align="center">GoStaticServe เป็นเซิร์ฟเวอร์ไฟล์แบบคงที่ที่มีประสิทธิภาพสูงและเบากว่าที่พัฒนาด้วยภาษา Go มันสนับสนุนการทำงานพื้นฐาน เช่น อัปโหลด ดาวน์โหลด และลบไฟล์ และยังให้คุณสมบัติต่างๆ เช่น การควบคุมสิทธิ์และบันทึก日志</p>

## คุณสมบัติ

- 🚀 **ประสิทธิภาพสูง**: ใช้คุณสมบัติความพร้อมรับงานสูงของภาษา Go ทำให้สามารถประมวลผลคำขอไฟล์จำนวนมากได้อย่างมีประสิทธิภาพ
- 🔒 **การควบคุมสิทธิ์**: รองรับการยืนยันตัวตนพื้นฐานและการควบคุมสิทธิ์เข้าถึงไฟล์
- 📝 **บันทึก日志ละเอียด**: บันทึกคำขอและ日志การทำงานเพื่อแก้ไขปัญหาและการตรวจสอบ
- 📁 **การทำงานกับไฟล์**: รองรับอัปโหลด ดาวน์โหลด ลบ เปลี่ยนชื่อไฟล์ ฯลฯ
- 📋 **รายชื่อโฟลเดอร์**: สร้างรายชื่อโฟลเดอร์อัตโนมัติเพื่อความสะดวกในการนำทางและเข้าถึงไฟล์
- ⚙️ **การตั้งค่าแบบยืดหยุ่น**: รองรับการตั้งค่าผ่านไฟล์การตั้งค่าตัวแปรสภาพแวดล้อมและพารามิเตอร์บรรทัดคำสั่ง
- 🐳 **การสนับสนุนคอนเทนเนอร์**: ให้รูปภาพ Docker เพื่อความสะดวกในการปรับใช้และการทำงาน
- 🚀 **การปรับใช้อัตโนมัติ**: รองรับการปรับใช้อัตโนมัติในเซิร์ฟเวอร์ระยะไกลผ่าน GitHub Actions

## เริ่มต้นอย่างรวดเร็ว

### การติดตั้ง

#### ติดตั้งโดยใช้ Go

```bash
go install github.com/3ziye/GoStaticServe@latest
```

#### สร้างจากโค้ดต้นฉบับ

```bash
git clone https://github.com/3ziye/GoStaticServe.git
cd GoStaticServe
go mod init github.com/3ziye/GoStaticServe
go build -o GoStaticServe cmd/server/main.go
```

#### 使用แฟ้มไบนารีที่คอมไพล์ไว้ก่อนหน้า

GoStaticServe ให้แฟ้มไบนารีที่คอมไพล์ไว้ก่อนหน้าสำหรับระบบ Linux Windows และ Mac ซึ่งสามารถดาวน์โหลดโดยตรงเพื่อใช้งานได้

1. เยี่ยมชม [หน้าเผยแพร่ของ GitHub](https://github.com/3ziye/GoStaticServe/releases) และดาวน์โหลดแพ็คเกจไฟล์ไบนารีที่ตรงกับแพลตฟอร์มของคุณ

2. แตกไฟล์ที่เกี่ยวข้องตามระบบปฏิบัติการที่ใช้:

   **Linux:**
   ```bash
   # ดาวน์โหลดและแตกไฟล์เวอร์ชันสำหรับ Linux
   wget https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_linux_amd64.tar.gz
chmod +x GoStaticServe
   
   # รันบริการ
   ./GoStaticServe
   ```
   
   **Windows:**
   ```powershell
   # ดาวน์โหลดและแตกไฟล์เวอร์ชันสำหรับ Windows
   # คลิกขวาที่ไฟล์ zip ที่ดาวน์โหลดแล้วเลือก "Extract here"
   
   # รันบริการ
   .\GoStaticServe.exe
   ```
   
   **Mac:**
   ```bash
   # ดาวน์โหลดและแตกไฟล์เวอร์ชันสำหรับ Mac
   curl -OL https://github.com/3ziye/GoStaticServe/releases/download/vX.Y.Z/GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
tar -xzf GoStaticServe_vX.Y.Z_darwin_amd64.tar.gz
chmod +x GoStaticServe
   
   # รันบริการ
   ./GoStaticServe
   ```

3. ตรวจสอบข้อมูลเวอร์ชัน
   หลังจากเริ่มบริการ คุณสามารถดูหมายเลขเวอร์ชันปัจจุบันใน日志คอนโซลได้

#### 使用 Docker

ดาวน์โหลดภาพจาก Docker Hub:
```bash
docker pull 3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads 3ziye/gostaticserve:latest
```

ดาวน์โหลดภาพจาก GitHub Package:
```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

## ภาพ Docker จาก GitHub Package

### ดาวน์โหลดภาพจาก GitHub Package

1. ตรวจสอบให้แน่ใจว่า Docker ถูกติดตั้ง
2. รันคำสั่งต่อไปนี้เพื่อดาวน์โหลดภาพ:
   ```bash
docker pull ghcr.io/3ziye/gostaticserve:latest
   ```

### รันภาพ Docker

```bash
docker run -p 8080:8080 -v ./files:/app/uploads ghcr.io/3ziye/gostaticserve:latest
```

คำสั่งนี้จะรันคอนเทนเนอร์ แมปพอร์ต 8080 ของคอนเทนเนอร์ไปยังพอร์ต 8080 ของโฮสต์ และเมานท์ไดเร็กทอรี `./files` ของโฮสต์ไปยังไดเร็กทอรี `/app/uploads` ภายในคอนเทนเนอร์

### อัปโหลดภาพไปยัง GitHub Package (สำหรับผู้ดูแลเท่านั้น)

1. ลงชื่อเข้าใช้ GitHub Container Registry:
   ```bash
docker login ghcr.io -u ชื่อผู้ใช้GitHub -p โทเคนGitHub
   ```

2. สร้างภาพ:
   ```bash
docker build -t ghcr.io/3ziye/gostaticserve:latest .
   ```

3. อัปโหลดภาพ:
   ```bash
docker push ghcr.io/3ziye/gostaticserve:latest
   ```

## การปรับใช้อัตโนมัติในเซิร์ฟเวอร์

GoStaticServe รองรับการปรับใช้อัตโนมัติในเซิร์ฟเวอร์ระยะไกลผ่าน GitHub Actions สำหรับขั้นตอนการตั้งค่าโดยละเอียด โปรดดู [เอกสารการปรับใช้](doc/DEPLOYMENT.md)

คุณสมบัติหลักของการปรับใช้อัตโนมัติ:
- 🔑 จัดการข้อมูลประจำตัวเซิร์ฟเวอร์ผ่าน GitHub Secrets
- 📥 ดาวน์โหลดไบนารี Linux อัตโนมัติ
- 📁 คัดลอกสำรองและปรับใช้ในเซิร์ฟเวอร์
- 🚀 รองรับบริการ systemd
- ⚡ เตือนการปรับใช้ผ่านการเผยแพร่ (Release) หรือการดำเนินการด้วยตนเอง

## เวอร์ชันอื่นในภาษาต่างๆ

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
- [Nederlands](README_nl.md)
- [Polski](README_pl.md)
- [Svenska](README_sv.md)
- [Norsk](README_no.md)
- [Dansk](README_da.md)
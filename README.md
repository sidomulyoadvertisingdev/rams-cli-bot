# Boot Rams — KlikBCA CLI Scraper Bot (v1.0.0)

Aplikasi CLI interaktif berbasis bahasa pemrograman Go untuk melakukan penarikan data mutasi rekening KlikBCA secara manual maupun terjadwal secara otomatis. Dilengkapi dengan maskot Robot AI berkedip yang interaktif dan penanganan sesi logout yang aman untuk menghindari deteksi bot oleh sistem bank.

## 🚀 Fitur Utama
- **Dashboard CLI Interaktif**: Dilengkapi maskot Robot AI yang berkedip di menu utama terminal.
- **Dua Mode Operasi**:
  1. **Jalankan Bot `/run`**: Mengambil mutasi rekening hari ini secara cepat dan langsung mencetak tabel transaksi di terminal.
  2. **Jadwalkan Bot `/scheduler`**: Menjalankan bot secara periodik (misal: setiap 30 atau 60 menit) di latar belakang (**Headless Mode** senyap) sehingga tidak mengganggu aktivitas Anda di layar.
- **Konfigurasi Kredensial Mudah `/setting`**: Pengaturan User ID dan Password/PIN KlikBCA langsung dari CLI yang akan disimpan dengan aman di file `.env`.
- **Auto-Logout & Auto-Close**: Secara otomatis me-logout akun bank dan menutup browser Chrome setelah data mutasi berhasil ditarik untuk mencegah pembekuan akun.
- **Aman Dari Blokir Dialog**: Dilengkapi sistem pendeteksi dialog peringatan otomatis untuk langsung menyetujui/menutup pop-up limitasi akses KlikBCA.
- **Perintah Pintasan Global `rams`**: Dapat dijalankan langsung dari direktori mana pun di terminal macOS Anda cukup dengan mengetik `rams`.

---

## 🛠️ Prasyarat
Sebelum menginstal, pastikan sistem Anda sudah terpasang:
1. **Go (Golang)**: Versi 1.20 atau yang lebih baru.
   - Periksa dengan perintah: `go version`
2. **Google Chrome / Chromium**: Terpasang di lokasi standar macOS.
3. **GitHub CLI (Opsional)**: Untuk integrasi repository.

---

## ⚙️ Cara Instalasi & Penggunaan

### 1. Kloning Repositori
```bash
git clone https://github.com/sidomulyoadvertisingdev/bca-cli-bot.git
cd bca-cli-bot
```

### 2. Konfigurasi File Lingkungan (.env)
Salin berkas contoh konfigurasi ke berkas `.env` aktif:
```bash
cp .env.example .env
```
Buka file `.env` yang baru dibuat dan isi kredensial akun KlikBCA Anda:
```env
BCA_USER=USER_ID_ANDA
BCA_PASS=PIN_PASSWORD_ANDA
```
*(Catatan: File `.env` sudah masuk dalam `.gitignore` sehingga aman dan tidak akan pernah terunggah ke GitHub).*

### 3. Menjalankan Aplikasi
Untuk menjalankan bot langsung via Go:
```bash
go run .
```
Atau menggunakan berkas shell pembantu:
```bash
./run.sh
```

---

## 🌍 Pemasangan Perintah Global `rams`
Agar aplikasi bisa dibuka dari direktori mana saja di terminal Anda (layaknya aplikasi CLI profesional):

1. **Bangun Binary & Pasang Symlink**:
   ```bash
   go build -o rams .
   mkdir -p ~/.local/bin
   ln -sf $(pwd)/rams ~/.local/bin/rams
   ```
2. **Pastikan `~/.local/bin` terdaftar di PATH Anda**:
   Buka berkas konfigurasi terminal Anda (seperti `~/.zshrc` atau `~/.bash_profile`), lalu tambahkan baris berikut jika belum ada:
   ```bash
   export PATH="$HOME/.local/bin:$PATH"
   ```
   Terapkan perubahan dengan menjalankan `source ~/.zshrc` atau buka jendela terminal baru.

3. **Jalankan Aplikasi dari Direktori Mana Saja**:
   ```bash
   rams
   ```

---

## 📂 Struktur Proyek
```
bca-cli-bot/
├── bot/
│   ├── browser.go   # Otomasi Go-rod KlikBCA & penanganan dialog
│   └── model.go     # Model data mutasi & pencetakan tabel
├── .env             # File kredensial login (diabaikan oleh git)
├── .env.example     # Contoh template konfigurasi kredensial
├── .gitignore       # Konfigurasi pengecualian file Git
├── go.mod           # Konfigurasi dependensi modul Go
├── go.sum           # Checksum dependensi Go
├── main.go          # Menu CLI utama, scheduler, & animasi Robot AI
├── README.md        # Panduan instalasi dan penggunaan proyek
└── run.sh           # Berkas shell script pembantu
```

---

## ⚠️ Disklaimer
*Aplikasi ini dibuat murni untuk keperluan otomasi administrasi personal. Penulis tidak bertanggung jawab atas penyalahgunaan akun atau pemblokiran akses akibat penggunaan interval scheduler yang terlalu sering. Disarankan untuk menggunakan interval scheduler minimal 30 - 60 menit demi keamanan akun Anda.*

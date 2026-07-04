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

## ⚙️ 1. Cara Instalasi (Installation Guide)

Ikuti langkah-langkah berikut secara berurutan untuk memasang bot di komputer Anda:

### Langkah 1.1: Kloning Repositori
Unduh kode sumber proyek dari repositori GitHub Anda:
```bash
git clone https://github.com/sidomulyoadvertisingdev/bca-cli-bot.git
cd bca-cli-bot
```

### Langkah 1.2: Konfigurasi Kredensial (.env)
Salin berkas template contoh konfigurasi ke berkas `.env` aktif:
```bash
cp .env.example .env
```
Buka berkas `.env` yang baru dibuat menggunakan editor teks pilihan Anda, lalu isi kredensial akun KlikBCA Anda:
```env
BCA_USER=USER_ID_KLIKBCA_ANDA
BCA_PASS=PIN_PASSWORD_ANDA
```
*(Catatan: Berkas `.env` telah didaftarkan dalam `.gitignore` sehingga tidak akan pernah terunggah ke repositori publik/pribadi di GitHub demi menjaga keamanan akun Anda).*

### Langkah 1.3: Pemasangan Perintah Global `rams` (Opsional - Direkomendasikan)
Agar aplikasi dapat dibuka dari folder mana saja di terminal Anda (seperti aplikasi CLI sistem):
1. **Kompilasi Biner & Pasang Tautan (Symlink)**:
   ```bash
   go build -o rams .
   mkdir -p ~/.local/bin
   ln -sf $(pwd)/rams ~/.local/bin/rams
   ```
2. **Daftarkan ke PATH Terminal**:
   Buka berkas profil terminal Anda (seperti `~/.zshrc` untuk pengguna Zsh default macOS, atau `~/.bash_profile` untuk Bash), lalu tambahkan baris berikut di baris paling bawah:
   ```bash
   export PATH="$HOME/.local/bin:$PATH"
   ```
   Terapkan perubahan dengan mengetik perintah `source ~/.zshrc` atau buka jendela terminal baru.

---

## 🚀 2. Cara Menjalankan Aplikasi (Execution Guide)

Setelah proses instalasi dan konfigurasi di atas selesai, Anda dapat menjalankan aplikasi dengan salah satu cara di bawah ini:

### Metode 2.1: Menjalankan Secara Global (Sangat Direkomendasikan)
Jika Anda telah menyelesaikan **Langkah 1.3**, Anda cukup membuka terminal baru di direktori mana saja dan langsung mengetik:
```bash
rams
```
Aplikasi dashboard CLI **Boot Rams** akan langsung terbuka dan siap digunakan.

### Metode 2.2: Menjalankan Secara Lokal (Folder Proyek)
Jika Anda tidak memasang perintah global, Anda harus masuk ke dalam direktori proyek terlebih dahulu, kemudian jalankan dengan salah satu opsi berikut:
- **Menggunakan Go run**:
  ```bash
  go run .
  ```
- **Menggunakan berkas shell script pembantu**:
  ```bash
  ./run.sh
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

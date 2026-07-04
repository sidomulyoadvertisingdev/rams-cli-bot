package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"bca-cli-bot/bot"
)

func getEnvFilePath() string {
	execPath, err := os.Executable()
	if err == nil {
		realPath, err := filepath.EvalSymlinks(execPath)
		if err == nil {
			return filepath.Join(filepath.Dir(realPath), ".env")
		}
	}
	return ".env"
}

func showStaticLayout() {
	logo := []string{
		"╭─USER@SYSTEM:~─────────────────────────────────────────────────────────────────────────────╮",
		"│             _                                                                             │",
		"│            (○)                                                                            │",
		"│             │                                                                             │",
		"│         .───┴───.        ___    _   __  __  ___   ___   ___  _____                        │",
		"│       .-'       '-.     | _ \\  /_\\ |  \\/  |/ __| | _ ) / _ \\|_   _|                       │",
		"│      /   .-----.   \\    |   / / _ \\| |\\/| |\\__ \\ | _ \\| (_) | | |                         │",
		"│     /   /  o o  \\   \\   |_|_\\/_/ \\_\\_|  |_||___/ |___/ \\___/  |_|                         │",
		"│    |   |   \\ = /   |                                                                      │",
		"│    |---|    '-'    |---| ───────────────────────────────────────────────────────────────  │",
		"│    |   |           |   | > AI ASSISTANT • SMART • FAST • RELIABLE                         │",
		"│     \\   \\         /   /  .──────────────────────────────────────────────────────────.     │",
		"│      '-._'-----'_.-'     │ > [★] INTELLIGENT RESPONSE │ SYSTEM STATUS: ONLINE       │     │",
		"│        /         \\       │ > [⚡] HIGH PERFORMANCE    │ VERSION      : 1.0.0        │     │",
		"│       /  .-----.  \\      │ > [🔒] SECURE & PRIVATE    │ UPTIME       : 24/7         │     │",
		"│      /  /   AI  \\  \\     │ > [✦] ALWAYS ONLINE        │ MODE         : ACTIVE       │     │",
		"│     |  |   [=]   |  |    '──────────────────────────────────────────────────────────'     │",
		"│     |  |         |  |                                                                     │",
		"│                          > RAMS BOT READY TO ASSIST YOU... █                              │",
		"╰───────────────────────────────────────────────────────────────────────────────────────────╯",
	}
	greenColor := "\033[92m"
	resetColor := "\033[0m"
	for _, line := range logo {
		fmt.Println(greenColor + line + resetColor)
	}
	fmt.Println("\n📌 MENU UTAMA:")
	fmt.Println("👉 [1] /run       - Jalankan Bot Mutasi Rekening")
	fmt.Println("👉 [2] /setting    - Atur kredensial login (.env)")
	fmt.Println("👉 [3] /scheduler  - Jalankan bot terjadwal otomatis")
	fmt.Println("👉 [4] /exit       - Keluar dari aplikasi")
}

func playIntroBlinkAnimation() {
	logoOpen := []string{
		"╭─USER@SYSTEM:~─────────────────────────────────────────────────────────────────────────────╮",
		"│             _                                                                             │",
		"│            (○)                                                                            │",
		"│             │                                                                             │",
		"│         .───┴───.        ___    _   __  __  ___   ___   ___  _____                        │",
		"│       .-'       '-.     | _ \\  /_\\ |  \\/  |/ __| | _ ) / _ \\|_   _|                       │",
		"│      /   .-----.   \\    |   / / _ \\| |\\/| |\\__ \\ | _ \\| (_) | | |                         │",
		"│     /   /  o o  \\   \\   |_|_\\/_/ \\_\\_|  |_||___/ |___/ \\___/  |_|                         │",
		"│    |   |   \\ = /   |                                                                      │",
		"│    |---|    '-'    |---| ───────────────────────────────────────────────────────────────  │",
		"│    |   |           |   | > AI ASSISTANT • SMART • FAST • RELIABLE                         │",
		"│     \\   \\         /   /  .──────────────────────────────────────────────────────────.     │",
		"│      '-._'-----'_.-'     │ > [★] INTELLIGENT RESPONSE │ SYSTEM STATUS: ONLINE       │     │",
		"│        /         \\       │ > [⚡] HIGH PERFORMANCE    │ VERSION      : 1.0.0        │     │",
		"│       /  .-----.  \\      │ > [🔒] SECURE & PRIVATE    │ UPTIME       : 24/7         │     │",
		"│      /  /   AI  \\  \\     │ > [✦] ALWAYS ONLINE        │ MODE         : ACTIVE       │     │",
		"│     |  |   [=]   |  |    '──────────────────────────────────────────────────────────'     │",
		"│     |  |         |  |                                                                     │",
		"│                          > RAMS BOT READY TO ASSIST YOU... █                              │",
		"╰───────────────────────────────────────────────────────────────────────────────────────────╯",
	}

	eyeOpenLine := "│     /   /  o o  \\   \\   |_|_\\/_/ \\_\\_|  |_||___/ |___/ \\___/  |_|                         │"
	eyeClosedLine := "│     /   /  - -  \\   \\   |_|_\\/_/ \\_\\_|  |_||___/ |___/ \\___/  |_|                         │"

	greenColor := "\033[92m"
	resetColor := "\033[0m"

	// 1. Tampilkan logo awal (mata terbuka) dan menu utama
	clearScreen()
	for _, line := range logoOpen {
		fmt.Println(greenColor + line + resetColor)
	}
	fmt.Println("\n📌 MENU UTAMA:")
	fmt.Println("👉 [1] /run       - Jalankan Bot Mutasi Rekening")
	fmt.Println("👉 [2] /setting    - Atur kredensial login (.env)")
	fmt.Println("👉 [3] /scheduler  - Jalankan bot terjadwal otomatis")
	fmt.Println("👉 [4] /exit       - Keluar dari aplikasi")
	
	// Kedip pertama (tutup mata sebentar lalu buka)
	time.Sleep(400 * time.Millisecond)
	fmt.Print("\033[18A\r") // Naik ke baris mata
	fmt.Print("\033[K")
	fmt.Println(greenColor + eyeClosedLine + resetColor)
	fmt.Print("\033[17B\r") // Turun kembali
	
	time.Sleep(150 * time.Millisecond)
	fmt.Print("\033[18A\r")
	fmt.Print("\033[K")
	fmt.Println(greenColor + eyeOpenLine + resetColor)
	fmt.Print("\033[17B\r")

	// Kedip kedua cepat
	time.Sleep(200 * time.Millisecond)
	fmt.Print("\033[18A\r")
	fmt.Print("\033[K")
	fmt.Println(greenColor + eyeClosedLine + resetColor)
	fmt.Print("\033[17B\r")

	time.Sleep(150 * time.Millisecond)
	fmt.Print("\033[18A\r")
	fmt.Print("\033[K")
	fmt.Println(greenColor + eyeOpenLine + resetColor)
	fmt.Print("\033[17B\r")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// Pemuatan .env di awal aplikasi
	loadEnvFile()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Jalankan kedipan saat masuk menu utama, setelah selesai baru minta input
		playIntroBlinkAnimation()

		fmt.Print("\nPilih opsi atau ketik perintah (contoh: /run): ")

		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		switch strings.ToLower(input) {
		case "1", "/run":
			runBotFlow(scanner)
		case "2", "/setting":
			runSettingFlow(scanner)
		case "3", "/scheduler":
			runSchedulerFlow(scanner)
		case "4", "/exit", "exit":
			fmt.Println("\n👋 Terima kasih telah menggunakan Boot Rams. Sampai jumpa!")
			return
		default:
			fmt.Println("\n❌ Pilihan tidak dikenal. Gunakan /run, /setting, /scheduler, atau /exit.")
			fmt.Println("Tekan [Enter] untuk kembali ke menu utama...")
			scanner.Scan()
		}
	}
}

func loadEnvFile() {
	envPath := getEnvFilePath()
	if _, err := os.Stat(envPath); err == nil {
		file, err := os.Open(envPath)
		if err == nil {
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					val := strings.TrimSpace(parts[1])
					val = strings.Trim(val, `"'`)
					os.Setenv(key, val)
				}
			}
		}
	}
}

func loadEnvMap() map[string]string {
	config := map[string]string{
		"BCA_USER": "",
		"BCA_PASS": "",
	}
	envPath := getEnvFilePath()
	file, err := os.Open(envPath)
	if err != nil {
		return config
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			if key == "BCA_USER" || key == "BCA_PASS" {
				config[key] = val
			}
		}
	}
	return config
}

func saveEnv(user, pass string) error {
	content := fmt.Sprintf("BCA_USER=%s\nBCA_PASS=%s\n", user, pass)
	return os.WriteFile(getEnvFilePath(), []byte(content), 0644)
}

func runSettingFlow(scanner *bufio.Scanner) {
	fmt.Println("\n🛠️  PENGATURAN KREDENSIAL MBANKING:")
	
	currentConfig := loadEnvMap()
	
	fmt.Printf("1. Masukkan User ID KlikBCA [%s]: ", currentConfig["BCA_USER"])
	scanner.Scan()
	user := strings.TrimSpace(scanner.Text())
	if user == "" {
		user = currentConfig["BCA_USER"]
	}
	
	fmt.Printf("2. Masukkan Password/PIN KlikBCA [%s]: ", currentConfig["BCA_PASS"])
	scanner.Scan()
	pass := strings.TrimSpace(scanner.Text())
	if pass == "" {
		pass = currentConfig["BCA_PASS"]
	}
	
	err := saveEnv(user, pass)
	if err != nil {
		fmt.Printf("\n❌ Gagal menyimpan pengaturan: %v\n", err)
	} else {
		fmt.Println("\n✅ Pengaturan kredensial berhasil disimpan ke file .env!")
		os.Setenv("BCA_USER", user)
		os.Setenv("BCA_PASS", pass)
	}
}

func runBotFlow(scanner *bufio.Scanner) {
	fmt.Println("\n🚀 Memulai Bot Mutasi Rekening KlikBCA...")
	
	user := os.Getenv("BCA_USER")
	pass := os.Getenv("BCA_PASS")
	
	if user == "" || pass == "" {
		config := loadEnvMap()
		user = config["BCA_USER"]
		pass = config["BCA_PASS"]
		if user != "" {
			os.Setenv("BCA_USER", user)
		}
		if pass != "" {
			os.Setenv("BCA_PASS", pass)
		}
	}

	if user == "" || pass == "" {
		fmt.Println("\n❌ User ID atau Password belum diatur! Gunakan perintah /setting terlebih dahulu.")
		return
	}

	fmt.Println("\n👤 Menjalankan bot untuk User ID:", user)
	fmt.Println("⏳ Mohon tunggu, bot sedang membuka browser...")

	defaultStart, defaultEnd := bot.DefaultDateRange()
	
	_, err := bot.RunBCA(user, pass, defaultStart, defaultEnd)
	if err != nil {
		fmt.Printf("\n❌ Terjadi kesalahan saat menjalankan bot: %v\n", err)
	} else {
		fmt.Println("\n✅ Penarikan data mutasi selesai!")
	}
}

func runSchedulerFlow(scanner *bufio.Scanner) {
	fmt.Println("\n⏰ PENGATURAN PENJADWALAN OTOMATIS:")
	fmt.Print("Masukkan interval waktu bot berjalan otomatis (dalam menit, contoh: 30): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	
	var interval int
	_, err := fmt.Sscanf(input, "%d", &interval)
	if err != nil || interval <= 0 {
		fmt.Println("❌ Interval tidak valid. Harus berupa angka positif lebih besar dari 0.")
		return
	}

	fmt.Printf("\n🚀 Mode Penjadwalan Aktif! Bot akan berjalan setiap %d menit.\n", interval)
	fmt.Println("📌 Ketik 'stop' lalu tekan Enter untuk menghentikan dan kembali ke menu utama.")
	fmt.Println("⚠️  Perhatian: Untuk menghindari deteksi bot oleh BCA, jangan gunakan interval yang terlalu sering (disarankan minimal 30-60 menit).")
	
	// Set headless ke true agar tidak mengganggu pengguna dengan membuka jendela Chrome berulang kali
	os.Setenv("BCA_HEADLESS", "true")
	defer os.Setenv("BCA_HEADLESS", "false")

	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	defer ticker.Stop()

	// Jalankan pertama kali saat diaktifkan
	fmt.Printf("\n[ %s ] Menjalankan bot pertama kali...\n", time.Now().Format("15:04:05"))
	triggerBotExecution()

	// Channel untuk menangkap input stop
	stopChan := make(chan bool, 1)
	go func() {
		for {
			if scanner.Scan() {
				text := strings.TrimSpace(scanner.Text())
				if strings.ToLower(text) == "stop" {
					stopChan <- true
					return
				}
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\n[ %s ] Menjalankan bot terjadwal otomatis...\n", time.Now().Format("15:04:05"))
			triggerBotExecution()
		case <-stopChan:
			fmt.Println("\n🛑 Mode penjadwalan otomatis dinonaktifkan.")
			return
		}
	}
}

func triggerBotExecution() {
	user := os.Getenv("BCA_USER")
	pass := os.Getenv("BCA_PASS")
	if user == "" || pass == "" {
		config := loadEnvMap()
		user = config["BCA_USER"]
		pass = config["BCA_PASS"]
	}

	if user == "" || pass == "" {
		fmt.Println("❌ Gagal: Kredensial login belum diatur.")
		return
	}

	defaultStart, defaultEnd := bot.DefaultDateRange()
	_, err := bot.RunBCA(user, pass, defaultStart, defaultEnd)
	if err != nil {
		fmt.Printf("❌ Gagal menjalankan bot terjadwal: %v\n", err)
	} else {
		fmt.Println("✅ Penjadwalan bot berhasil menyelesaikan tugas.")
	}
	fmt.Printf("⏳ Menunggu jadwal berikutnya...\n")
}

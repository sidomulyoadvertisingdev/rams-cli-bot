package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		`╭─USER@SYSTEM:~─────────────────────────────────────────────────────────────────────────────╮`,
		`│             _                                                                             │`,
		`│            (○)                                                                            │`,
		`│             │                                                                             │`,
		`│         .───┴───.        ___    _   __  __  ___   ___   ___  _____                        │`,
		`│       .-'       '-.     | _ \  /_\ |  \/  |/ __| | _ ) / _ \|_   _|                       │`,
		`│      /   .-----.   \    |   / / _ \| |\/| |\__ \ | _ \| (_) | | |                         │`,
		`│     /   /  o o  \   \   |_|_\/_/ \_\_|  |_||___/ |___/ \___/  |_|                         │`,
		`│    |   |   \ = /   |                                                                      │`,
		`│    |---|    '-'    |---| ───────────────────────────────────────────────────────────────  │`,
		`│    |   |           |   | > AI ASSISTANT • SMART • FAST • RELIABLE                         │`,
		`│     \   \         /   /  .──────────────────────────────────────────────────────────.     │`,
		`│      '-._'-----'_.-'     │ > [★] INTELLIGENT RESPONSE │ SYSTEM STATUS: ONLINE       │     │`,
		`│        /         \       │ > [⚡] HIGH PERFORMANCE    │ VERSION      : 1.0.0        │     │`,
		`│       /  .-----.  \      │ > [🔒] SECURE & PRIVATE    │ UPTIME       : 24/7         │     │`,
		`│      /  /   AI  \  \     │ > [✦] ALWAYS ONLINE        │ MODE         : ACTIVE       │     │`,
		`│     |  |   [=]   |  |    '──────────────────────────────────────────────────────────'     │`,
		`│     |  |         |  |                                                                     │`,
		`│                          > RAMS BOT READY TO ASSIST YOU... █                              │`,
		`╰───────────────────────────────────────────────────────────────────────────────────────────╯`,
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
		`╭─USER@SYSTEM:~─────────────────────────────────────────────────────────────────────────────╮`,
		`│             _                                                                             │`,
		`│            (○)                                                                            │`,
		`│             │                                                                             │`,
		`│         .───┴───.        ___    _   __  __  ___   ___   ___  _____                        │`,
		`│       .-'       '-.     | _ \  /_\ |  \/  |/ __| | _ ) / _ \|_   _|                       │`,
		`│      /   .-----.   \    |   / / _ \| |\/| |\__ \ | _ \| (_) | | |                         │`,
		`│     /   /  o o  \   \   |_|_\/_/ \_\_|  |_||___/ |___/ \___/  |_|                         │`,
		`│    |   |   \ = /   |                                                                      │`,
		`│    |---|    '-'    |---| ───────────────────────────────────────────────────────────────  │`,
		`│    |   |           |   | > AI ASSISTANT • SMART • FAST • RELIABLE                         │`,
		`│     \   \         /   /  .──────────────────────────────────────────────────────────.     │`,
		`│      '-._'-----'_.-'     │ > [★] INTELLIGENT RESPONSE │ SYSTEM STATUS: ONLINE       │     │`,
		`│        /         \       │ > [⚡] HIGH PERFORMANCE    │ VERSION      : 1.0.0        │     │`,
		`│       /  .-----.  \      │ > [🔒] SECURE & PRIVATE    │ UPTIME       : 24/7         │     │`,
		`│      /  /   AI  \  \     │ > [✦] ALWAYS ONLINE        │ MODE         : ACTIVE       │     │`,
		`│     |  |   [=]   |  |    '──────────────────────────────────────────────────────────'     │`,
		`│     |  |         |  |                                                                     │`,
		`│                          > RAMS BOT READY TO ASSIST YOU... █                              │`,
		`╰───────────────────────────────────────────────────────────────────────────────────────────╯`,
	}

	eyeOpenLine := `│     /   /  o o  \   \   |_|_\/_/ \_\_|  |_||___/ |___/ \___/  |_|                         │`
	eyeClosedLine := `│     /   /  - -  \   \   |_|_\/_/ \_\_|  |_||___/ |___/ \___/  |_|                         │`

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

	// Tampilkan layout pembuka lalu verifikasi autentikasi ERP
	clearScreen()
	showStaticLayout()
	authenticateERP(scanner)

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
	config := make(map[string]string)
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
			config[key] = val
		}
	}
	return config
}

func saveEnvMap(config map[string]string) error {
	var lines []string
	for k, v := range config {
		lines = append(lines, fmt.Sprintf("%s=%s", k, v))
	}
	content := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(getEnvFilePath(), []byte(content), 0644)
}

type LoginResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Message     string `json:"message"`
}

func loginToERP(apiURL, email, password string) (*LoginResponse, error) {
	url := fmt.Sprintf("%s/auth/login", strings.TrimSuffix(apiURL, "/"))
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %s", string(body))
	}

	if resp.StatusCode != http.StatusOK {
		if loginResp.Message != "" {
			return nil, fmt.Errorf(loginResp.Message)
		}
		return nil, fmt.Errorf("HTTP status %d", resp.StatusCode)
	}

	return &loginResp, nil
}

func authenticateERP(scanner *bufio.Scanner) {
	apiURL := os.Getenv("ERP_API_URL")
	if apiURL == "" {
		apiURL = "http://erp.sidomulyo.test/api"
		os.Setenv("ERP_API_URL", apiURL)
	}

	email := os.Getenv("ERP_EMAIL")
	password := os.Getenv("ERP_PASS")

	if email != "" && password != "" {
		fmt.Printf("🔄 Menghubungkan ke ERP (%s) sebagai %s...\n", apiURL, email)
		resp, err := loginToERP(apiURL, email, password)
		if err == nil {
			fmt.Printf("✅ Berhasil login ke ERP! Selamat datang.\n")
			os.Setenv("ERP_TOKEN", resp.AccessToken)
			time.Sleep(1 * time.Second)
			return
		}
		fmt.Printf("⚠️ Auto-login ERP gagal: %v\n", err)
	}

	for {
		fmt.Println("\n🔐 SILAKAN LOGIN ERP SIDOMULYO:")
		fmt.Printf("👉 URL API ERP [%s]: ", apiURL)
		if !scanner.Scan() {
			os.Exit(1)
		}
		inputURL := strings.TrimSpace(scanner.Text())
		if inputURL != "" {
			apiURL = inputURL
			os.Setenv("ERP_API_URL", apiURL)
		}

		fmt.Print("👉 Email       : ")
		if !scanner.Scan() {
			os.Exit(1)
		}
		inputEmail := strings.TrimSpace(scanner.Text())

		fmt.Print("👉 Password    : ")
		if !scanner.Scan() {
			os.Exit(1)
		}
		inputPassword := strings.TrimSpace(scanner.Text())

		if inputEmail == "" || inputPassword == "" {
			fmt.Println("❌ Email dan password tidak boleh kosong!")
			continue
		}

		fmt.Printf("🔄 Memverifikasi kredensial ke %s...\n", apiURL)
		resp, err := loginToERP(apiURL, inputEmail, inputPassword)
		if err != nil {
			fmt.Printf("❌ Login gagal: %v\n", err)
			continue
		}

		fmt.Printf("✅ Login berhasil! Selamat datang.\n")
		os.Setenv("ERP_EMAIL", inputEmail)
		os.Setenv("ERP_PASS", inputPassword)
		os.Setenv("ERP_TOKEN", resp.AccessToken)

		// Simpan kredensial baru ke .env
		config := loadEnvMap()
		config["BCA_USER"] = os.Getenv("BCA_USER")
		config["BCA_PASS"] = os.Getenv("BCA_PASS")
		config["ERP_API_URL"] = apiURL
		config["ERP_EMAIL"] = inputEmail
		config["ERP_PASS"] = inputPassword
		_ = saveEnvMap(config)

		time.Sleep(1 * time.Second)
		break
	}
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
	
	fmt.Println("\n🛠️  PENGATURAN API ERP SIDOMULYO:")
	
	erpURL := currentConfig["ERP_API_URL"]
	if erpURL == "" {
		erpURL = "http://erp.sidomulyo.test/api"
	}
	fmt.Printf("3. Masukkan URL API ERP [%s]: ", erpURL)
	scanner.Scan()
	inputURL := strings.TrimSpace(scanner.Text())
	if inputURL == "" {
		inputURL = erpURL
	}
	
	fmt.Printf("4. Masukkan Email ERP [%s]: ", currentConfig["ERP_EMAIL"])
	scanner.Scan()
	erpEmail := strings.TrimSpace(scanner.Text())
	if erpEmail == "" {
		erpEmail = currentConfig["ERP_EMAIL"]
	}
	
	fmt.Printf("5. Masukkan Password ERP [%s]: ", currentConfig["ERP_PASS"])
	scanner.Scan()
	erpPass := strings.TrimSpace(scanner.Text())
	if erpPass == "" {
		erpPass = currentConfig["ERP_PASS"]
	}
	
	currentConfig["BCA_USER"] = user
	currentConfig["BCA_PASS"] = pass
	currentConfig["ERP_API_URL"] = inputURL
	currentConfig["ERP_EMAIL"] = erpEmail
	currentConfig["ERP_PASS"] = erpPass
	
	err := saveEnvMap(currentConfig)
	if err != nil {
		fmt.Printf("\n❌ Gagal menyimpan pengaturan: %v\n", err)
	} else {
		fmt.Println("\n✅ Pengaturan kredensial berhasil disimpan ke file .env!")
		os.Setenv("BCA_USER", user)
		os.Setenv("BCA_PASS", pass)
		os.Setenv("ERP_API_URL", inputURL)
		os.Setenv("ERP_EMAIL", erpEmail)
		os.Setenv("ERP_PASS", erpPass)
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

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
	"golang.org/x/term"
)

func getEnvFilePath() string {
	if _, err := os.Stat(".env"); err == nil {
		return ".env"
	}
	projectEnv := "/Users/sidomulyo/Projects-SM/bca-cli-bot/.env"
	if _, err := os.Stat(projectEnv); err == nil {
		return projectEnv
	}
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

	// Jalankan kedipan saat masuk menu utama
	playIntroBlinkAnimation()

	var email, password string
	var err error
	for {
		email, password, err = promptERPLogin(scanner)
		if err != nil {
			fmt.Printf("❌ Autentikasi ERP Gagal: %v. Silakan coba lagi.\n", err)
			continue
		}
		break
	}
	fmt.Println("\n🔒 Login ERP Berhasil! Akses diberikan.")
	time.Sleep(1 * time.Second)

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
			runBotFlow(scanner, email, password)
		case "2", "/setting":
			runSettingFlow(scanner)
		case "3", "/scheduler":
			runSchedulerFlow(scanner, email, password)
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

func saveEnv(user, pass string) error {
	config := loadEnvMap()
	config["BCA_USER"] = user
	config["BCA_PASS"] = pass

	var lines []string
	keysOrder := []string{"BCA_USER", "BCA_PASS", "ERP_API_URL", "BCA_BOT_TOKEN", "ERP_EMAIL", "ERP_PASSWORD"}
	for _, k := range keysOrder {
		if val, exists := config[k]; exists {
			lines = append(lines, fmt.Sprintf("%s=%s", k, val))
		}
	}

	for k, val := range config {
		isDefaultKey := false
		for _, dk := range keysOrder {
			if k == dk {
				isDefaultKey = true
				break
			}
		}
		if !isDefaultKey {
			lines = append(lines, fmt.Sprintf("%s=%s", k, val))
		}
	}

	return os.WriteFile(getEnvFilePath(), []byte(strings.Join(lines, "\n")+"\n"), 0644)
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

func runBotFlow(scanner *bufio.Scanner, email, password string) {
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

	fmt.Println("\n🔑 Menghubungkan ke backend ERP...")
	accessToken, err := loginToERP(email, password)
	if err != nil {
		fmt.Printf("❌ Gagal login ke ERP: %v\n", err)
		return
	}

	fmt.Println("\n👤 Menjalankan bot untuk User ID:", user)
	fmt.Println("⏳ Mohon tunggu, bot sedang membuka browser...")

	defaultStart, defaultEnd := bot.DefaultDateRange()
	
	mutations, err := bot.RunBCA(user, pass, defaultStart, defaultEnd)
	if err != nil {
		fmt.Printf("\n❌ Terjadi kesalahan saat menjalankan bot: %v\n", err)
	} else {
		fmt.Println("\n✅ Penarikan data mutasi selesai!")
		sendMutationsToBackend(mutations, accessToken)
	}
}

func runSchedulerFlow(scanner *bufio.Scanner, email, password string) {
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
	triggerBotExecution(email, password)

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
			triggerBotExecution(email, password)
		case <-stopChan:
			fmt.Println("\n🛑 Mode penjadwalan otomatis dinonaktifkan.")
			return
		}
	}
}

func triggerBotExecution(email, password string) {
	user := os.Getenv("BCA_USER")
	pass := os.Getenv("BCA_PASS")
	if user == "" || pass == "" {
		config := loadEnvMap()
		user = config["BCA_USER"]
		pass = config["BCA_PASS"]
	}

	if user == "" || pass == "" {
		fmt.Println("❌ Gagal: Kredensial login KlikBCA belum diatur.")
		return
	}

	fmt.Println("🔑 Menghubungkan ke backend ERP...")
	accessToken, err := loginToERP(email, password)
	if err != nil {
		fmt.Printf("❌ Gagal login ke ERP saat jadwal otomatis: %v\n", err)
		return
	}

	defaultStart, defaultEnd := bot.DefaultDateRange()
	mutations, err := bot.RunBCA(user, pass, defaultStart, defaultEnd)
	if err != nil {
		fmt.Printf("❌ Gagal menjalankan bot terjadwal: %v\n", err)
	} else {
		fmt.Println("✅ Penjadwalan bot berhasil menyelesaikan tugas.")
		sendMutationsToBackend(mutations, accessToken)
	}
	fmt.Printf("⏳ Menunggu jadwal berikutnya...\n")
}

func sendMutationsToBackend(mutations []bot.Mutation, accessToken string) {
	apiUrl := os.Getenv("ERP_API_URL")
	botToken := os.Getenv("BCA_BOT_TOKEN")

	if apiUrl == "" {
		fmt.Println("⚠️  Peringatan: ERP_API_URL belum diatur di .env. Data mutasi tidak dikirim ke backend.")
		return
	}

	if botToken == "" {
		fmt.Println("⚠️  Peringatan: BCA_BOT_TOKEN belum diatur di .env. Data mutasi tidak dikirim ke backend.")
		return
	}

	type MutationPayload struct {
		Date        string `json:"date"`
		Description string `json:"description"`
		Amount      string `json:"amount"`
		Type        string `json:"type"`
	}

	type RequestPayload struct {
		Mutations []MutationPayload `json:"mutations"`
	}

	var reqMutations []MutationPayload
	for _, m := range mutations {
		reqMutations = append(reqMutations, MutationPayload{
			Date:        m.Date,
			Description: m.Description,
			Amount:      m.Amount,
			Type:        m.Type,
		})
	}

	payload := RequestPayload{Mutations: reqMutations}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("❌ Gagal membuat payload JSON: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Printf("❌ Gagal membuat request http: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Bot-Token", botToken)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Gagal mengirim mutasi ke backend: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		fmt.Println("✅ Berhasil mengirimkan data mutasi ke backend ERP!")
	} else {
		fmt.Printf("❌ Backend ERP mengembalikan status error: %s\n", resp.Status)
	}
}

type LoginResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

func loginToERP(email, password string) (string, error) {
	config := loadEnvMap()
	apiUrl := config["ERP_API_URL"]
	if apiUrl == "" {
		apiUrl = "http://127.0.0.1/api/pos/bank-mutations"
	}

	loginUrl := strings.Replace(apiUrl, "/api/pos/bank-mutations", "/api/auth/login", 1)

	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("gagal marshal login payload: %v", err)
	}

	req, err := http.NewRequest("POST", loginUrl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", fmt.Errorf("gagal membuat request login: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("gagal menghubungi backend ERP: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("gagal membaca response login: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Message string `json:"message"`
		}
		json.Unmarshal(bodyBytes, &errResp)
		msg := errResp.Message
		if msg == "" {
			msg = resp.Status
		}
		return "", fmt.Errorf("login ditolak oleh ERP: %s", msg)
	}

	var loginResp LoginResponse
	err = json.Unmarshal(bodyBytes, &loginResp)
	if err != nil {
		return "", fmt.Errorf("gagal parsing response login: %v", err)
	}

	if !loginResp.Success || loginResp.AccessToken == "" {
		return "", fmt.Errorf("login gagal: %s", loginResp.Message)
	}

	return loginResp.AccessToken, nil
}

func promptERPLogin(scanner *bufio.Scanner) (string, string, error) {
	fmt.Println("\n==============================================")
	fmt.Println("🔑   SILAKAN LOGIN KE ERP SIDOMULYO   🔑")
	fmt.Println("==============================================")
	fmt.Print("Masukkan Email ERP   : ")
	if !scanner.Scan() {
		return "", "", fmt.Errorf("batal")
	}
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan Password ERP: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", "", fmt.Errorf("gagal membaca password: %v", err)
	}
	fmt.Println() // Print newline after user presses Enter
	password := strings.TrimSpace(string(bytePassword))

	if email == "" || password == "" {
		return "", "", fmt.Errorf("email dan password tidak boleh kosong")
	}

	_, err = loginToERP(email, password)
	if err != nil {
		return "", "", err
	}

	return email, password, nil
}

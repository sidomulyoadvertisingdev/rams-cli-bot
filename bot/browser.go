package bot

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// RunBCA menjalankan seluruh alur: login KlikBCA → navigasi mutasi → scrape tabel.
// startDate dan endDate dalam format "ddmmyyyy" (contoh: "01072026").
func RunBCA(user, pass, startDate, endDate string) ([]Mutation, error) {
	// ── 1. Launch browser ──────────────────────────────────────────────────
	headless := os.Getenv("BCA_HEADLESS") == "true"
	u := launcher.New().
		Headless(headless).
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect()
	defer browser.MustClose()

	page, err := browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		return nil, fmt.Errorf("gagal membuka browser: %w", err)
	}

	// Daftarkan event handler untuk auto-dismiss dialogs (alerts)
	go page.EachEvent(func(e *proto.PageJavascriptDialogOpening) {
		log.Printf("⚠️ Dialog alert muncul: %s (Otomatis ditutup)", e.Message)
		_ = proto.PageHandleJavaScriptDialog{
			Accept:     true,
			PromptText: "",
		}.Call(page)
	})()

	log.Println("🌐 Membuka halaman login KlikBCA...")
	if err := page.Navigate("https://ibank.klikbca.com/login.jsp"); err != nil {
		return nil, fmt.Errorf("gagal buka halaman login KlikBCA: %w", err)
	}

	// Tunggu load maksimal 30 detik, lanjut meski timeout
	if err := page.WaitLoad(); err != nil {
		log.Println("⚠️  WaitLoad timeout, lanjut coba login...")
	}
	time.Sleep(4 * time.Second)

	// ── 2. Isi form login langsung di page (tidak ada frameset) ───────────────
	log.Println("⌨️  Memasukkan user ID...")
	if err := fillInput(page, "#txt_user_id, #user_id", user); err != nil {
		return nil, fmt.Errorf("gagal mengisi user ID: %w", err)
	}

	log.Println("⌨️  Memasukkan password...")
	if err := fillInput(page, "#txt_pswd, #pswd", pass); err != nil {
		return nil, fmt.Errorf("gagal mengisi password: %w", err)
	}

	log.Println("🔐 Mengklik tombol Login...")
	loginBtn, err := page.Element(`[name="value(Submit)"], input[type="submit"], #btnSubmit`)
	if err != nil {
		return nil, fmt.Errorf("tombol login tidak ditemukan: %w", err)
	}
	loginBtn.MustClick()

	// ── 3. Lacak tab aktif baru pasca-login ─────────────────────────────────
	log.Println("⏳ Menunggu halaman dashboard setelah login (±12 detik)...")
	time.Sleep(12 * time.Second)

	pages, err := browser.Pages()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil list pages: %w", err)
	}
	log.Printf("📱 Jumlah halaman terbuka: %d\n", len(pages))
	for _, p := range pages {
		info, _ := p.Info()
		if info != nil {
			log.Printf("  - URL: %s | Title: %s\n", info.URL, info.Title)
		}
	}

	// Cari page dashboard KlikBCA aktif (authentication.do)
	var activePage *rod.Page
	for _, p := range pages {
		info, err := p.Info()
		if err == nil && info != nil && strings.Contains(info.URL, "authentication.do") {
			activePage = p
			break
		}
	}

	if activePage == nil {
		// Cari page mana saja yang masih valid
		for _, p := range pages {
			_, err := p.Info()
			if err == nil {
				activePage = p
				break
			}
		}
	}

	if activePage == nil {
		// Fallback terakhir ke page utama
		activePage = page
	}

	page = activePage

	// Daftarkan event handler untuk auto-dismiss dialogs (alerts) pada page aktif
	go page.EachEvent(func(e *proto.PageJavascriptDialogOpening) {
		log.Printf("⚠️ Dialog alert muncul: %s (Otomatis ditutup)", e.Message)
		_ = proto.PageHandleJavaScriptDialog{
			Accept:     true,
			PromptText: "",
		}.Call(page)
	})()

	currentURL := page.MustInfo().URL
	log.Println("🌍 URL sekarang:", currentURL)

	// Periksa apakah login berhasil
	if strings.Contains(strings.ToLower(currentURL), "login") {
		loginHTML, _ := page.HTML()
		os.WriteFile("debug_login_fail.html", []byte(loginHTML), 0644)
		return nil, fmt.Errorf("❌ Login gagal — URL masih di halaman login\n" +
			"   Kemungkinan: password salah, CAPTCHA, atau perlu verifikasi OTP/KeyBCA")
	}

	// Login sukses, daftarkan defer logout agar sesi bank selalu ditutup jika terjadi error setelah ini
	loggedIn := true
	defer func() {
		if loggedIn {
			log.Println("🚪 [Cleanup] Mengirim perintah Logout (Navigasi langsung)...")
			_ = page.Navigate("https://ibank.klikbca.com/authentication.do?value(actions)=logout")
			log.Println("🔒 [Cleanup] Sesi berhasil di-logout")
			time.Sleep(1500 * time.Millisecond)
		}
	}()

	// Tunggu sampai ada frame yang termuat di page dashboard
	log.Println("⏳ Menunggu frame dashboard dimuat...")
	for i := 0; i < 20; i++ {
		if len(getFrameElements(page)) > 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// ── 4. Dapatkan frame menu (sisi kiri) ──────────────────────────────────
	menuFrame, err := getFrameByName(page, "menu")
	if err != nil {
		pageHTML, _ := page.HTML()
		os.WriteFile("debug_page.html", []byte(pageHTML), 0644)
		return nil, fmt.Errorf("gagal mendapatkan frame menu: %w", err)
	}

	// ── 5. Dapatkan frame main / atm (sisi kanan) ──────────────────────────
	mainFrame, err := getFrameByName(page, "atm")
	if err != nil {
		pageHTML, _ := page.HTML()
		os.WriteFile("debug_page.html", []byte(pageHTML), 0644)
		return nil, fmt.Errorf("gagal mendapatkan frame main: %w", err)
	}

	// ── 6. Navigasi menu secara bertahap menggunakan klik frame ───────────
	log.Println("📂 Menavigasi ke halaman Mutasi Rekening...")

	// 6.2 Klik 'Informasi Rekening' di frame menu
	log.Println("📂 Mengklik menu 'Informasi Rekening'...")
	if err := clickMenuLink(menuFrame, `/informasi rekening|account information/i`); err != nil {
		menuHTML, _ := menuFrame.HTML()
		os.WriteFile("debug_menu.html", []byte(menuHTML), 0644)
		pageHTML, _ := page.HTML()
		os.WriteFile("debug_page.html", []byte(pageHTML), 0644)
		return nil, fmt.Errorf("gagal klik menu Informasi Rekening: %w", err)
	}
	time.Sleep(1 * time.Second)

	// 6.3 Klik 'Mutasi Rekening' di frame menu
	log.Println("📂 Mengklik menu 'Mutasi Rekening'...")
	if err := clickMenuLink(menuFrame, `/mutasi rekening|account statement/i`); err != nil {
		menuHTML, _ := menuFrame.HTML()
		os.WriteFile("debug_menu.html", []byte(menuHTML), 0644)
		pageHTML, _ := page.HTML()
		os.WriteFile("debug_page.html", []byte(pageHTML), 0644)
		return nil, fmt.Errorf("gagal klik menu Mutasi Rekening: %w", err)
	}

	// 6.4 Tunggu formulir mutasi dimuat di frame main
	log.Println("⏳ Menunggu form mutasi dimuat di frame main...")
	var startDtEl *rod.Element
	for i := 0; i < 20; i++ {
		startDtEl, err = mainFrame.Element(`[name="value(startDt)"], [name="startDt"], [name="fromDay"]`)
		if err == nil && startDtEl != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if startDtEl == nil {
		return nil, fmt.Errorf("form mutasi rekening tidak muncul di frame main")
	}

	// ── 7. Debug: dump HTML halaman mutasi ──────────────────────────────
	mutasiURL2 := mainFrame.MustInfo().URL
	log.Println("🌐 URL halaman mutasi:", mutasiURL2)

	// Dump HTML ke file untuk debug
	mutasiHTML, _ := mainFrame.HTML()
	os.WriteFile("debug_mutasi.html", []byte(mutasiHTML), 0644)

	// Submit form di dalam mainFrame menggunakan tanggal default (hari ini)
	log.Println("🔍 Klik tombol Lihat Mutasi Rekening (menggunakan tanggal default/hari ini)...")
	mainFrame.MustEval(`() => {
		// Pastikan opsi Mutasi Harian terpilih (value=1)
		let r1 = document.querySelector('input[name="value(r1)"][value="1"]') ||
		         document.querySelector('input[name="r1"][value="1"]');
		if (r1) r1.checked = true;

		let btn = document.querySelector('input[name="value(submit1)"]') ||
		          document.querySelector('input[name="value(Submit)"]') ||
		          document.querySelector('input[type="submit"]') ||
		          document.querySelector('input[type="Submit"]') ||
		          document.querySelector('input[value*="Lihat"]') ||
		          document.querySelector('input[value*="Tampilkan"]') ||
		          document.querySelector('button[type="submit"]') ||
		          document.querySelector('input[name="Submit"]');
		if (btn) { btn.click(); }
		else { console.error('Submit button tidak ditemukan'); }
	}`)

	log.Println("⏳ Menunggu hasil mutasi...")
	time.Sleep(3 * time.Second) // basic sleep to let the frame reload

	// Wait for the table to appear in the main frame (up to 7.5 seconds)
	var rows rod.Elements
	for i := 0; i < 15; i++ {
		rows, err = mainFrame.Elements("table tr")
		if err == nil && len(rows) > 5 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// ── 9. Scrape tabel mutasi ──────────────────────────────────────────────
	log.Println("📊 Scraping tabel mutasi...")

	var mutations []Mutation

	// Scrape dari mainFrame
	mutations = append(mutations, scrapeTable(mainFrame)...)

	// Hapus duplikat berdasarkan (date+description+amount)
	mutations = deduplicate(mutations)

	log.Printf("✅ Ditemukan %d transaksi\n", len(mutations))

	// Tampilkan hasil mutasi ke konsol terlebih dahulu sebelum logout/tutup browser
	fmt.Println()
	fmt.Printf("📋 MUTASI REKENING BCA — %s s/d %s\n",
		FormatDateDisplay(startDate), FormatDateDisplay(endDate))
	fmt.Println()
	PrintTable(mutations)
	fmt.Println()

	// ── 10. Logout ──────────────────────────────────────────────────────────
	loggedIn = false // Matikan flag agar defer cleanup tidak melakukan logout ganda
	log.Println("🚪 Mengirim perintah Logout (Navigasi langsung)...")
	_ = page.Navigate("https://ibank.klikbca.com/authentication.do?value(actions)=logout")
	log.Println("🔒 Sesi berhasil di-logout")
	time.Sleep(2 * time.Second)

	return mutations, nil
}

// getFrameByName mengambil frame/iframe berdasarkan name attribute (case-insensitive).
func getFrameByName(page *rod.Page, name string) (*rod.Page, error) {
	var foundNames []string
	for i := 0; i < 10; i++ {
		frames := getFrameElements(page)
		foundNames = nil
		for _, frameEl := range frames {
			fName, err := frameEl.Attribute("name")
			if err == nil && fName != nil {
				foundNames = append(foundNames, *fName)
				if strings.EqualFold(*fName, name) {
					return frameEl.Frame()
				}
			} else {
				fID, err := frameEl.Attribute("id")
				if err == nil && fID != nil {
					foundNames = append(foundNames, "id:"+*fID)
				} else {
					foundNames = append(foundNames, "(no name/id)")
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil, fmt.Errorf("frame %q tidak ditemukan. Frame yang terdeteksi: %v", name, foundNames)
}

// clickMenuLink mencari dan mengklik link menu berdasarkan pola teks regex.
func clickMenuLink(menuFrame *rod.Page, pattern string) error {
	var el *rod.Element
	var err error
	for i := 0; i < 15; i++ {
		el, err = menuFrame.ElementR("a", pattern)
		if err == nil && el != nil {
			break
		}
		el, err = menuFrame.ElementR("td", pattern)
		if err == nil && el != nil {
			break
		}
		el, err = menuFrame.ElementR("*", pattern)
		if err == nil && el != nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if el == nil {
		return fmt.Errorf("link menu dengan pola %q tidak ditemukan", pattern)
	}
	el.MustScrollIntoView()
	el.MustClick()
	return nil
}

// ── Helpers ──────────────────────────────────────────────────────────────────

// findFrameWithElement mencari frame/iframe yang mengandung selector tertentu.
// Mengembalikan *rod.Page yang merepresentasikan frame tersebut.
func findFrameWithElement(page *rod.Page, selector string) (*rod.Page, error) {
	// Coba di page utama dulu
	el, err := page.Element(selector)
	if err == nil && el != nil {
		return page, nil
	}

	// Cari di setiap frame element
	for _, frameEl := range getFrameElements(page) {
		fp, err := frameEl.Frame()
		if err != nil {
			continue
		}
		time.Sleep(500 * time.Millisecond)
		el, err := fp.Element(selector)
		if err == nil && el != nil {
			return fp, nil
		}
	}

	return nil, fmt.Errorf("selector %q tidak ditemukan di page maupun frame", selector)
}

// getFrameElements mengambil semua element <frame> dan <iframe> dari sebuah page.
func getFrameElements(page *rod.Page) rod.Elements {
	var els rod.Elements
	// frame (frameset)
	if frames, err := page.Elements("frame"); err == nil {
		els = append(els, frames...)
	}
	// iframe
	if iframes, err := page.Elements("iframe"); err == nil {
		els = append(els, iframes...)
	}
	return els
}

// fillInput mengisi input field dengan value tertentu menggunakan JS direct evaluation untuk menghindari stuck/delay.
func fillInput(page *rod.Page, selector, value string) error {
	tPage := page.Timeout(15 * time.Second)
	_, err := tPage.Element(selector)
	if err != nil {
		// Dump HTML untuk analisa kenapa gagal
		htmlVal, _ := page.HTML()
		os.WriteFile("debug_fill_input_fail.html", []byte(htmlVal), 0644)
		return fmt.Errorf("element %q tidak ditemukan/timeout: %w", selector, err)
	}
	
	_, errEval := tPage.Eval(`(sel, val) => {
		let el = document.querySelector(sel);
		if (el) {
			el.value = val;
			el.dispatchEvent(new Event('input', { bubbles: true }));
			el.dispatchEvent(new Event('change', { bubbles: true }));
			el.focus();
		}
	}`, selector, value)
	if errEval != nil {
		htmlVal, _ := page.HTML()
		os.WriteFile("debug_fill_input_fail.html", []byte(htmlVal), 0644)
		return fmt.Errorf("gagal mengevaluasi JS input untuk %q: %w", selector, errEval)
	}
	return nil
}

// scrapeTable mengambil data mutasi dari tabel HTML di sebuah page/frame.
func scrapeTable(page *rod.Page) []Mutation {
	var mutations []Mutation

	rows, err := page.Elements("table tr")
	if err != nil || len(rows) == 0 {
		return mutations
	}

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		cells, err := row.Elements("td")
		if err != nil || len(cells) < 5 {
			continue
		}

		date := strings.TrimSpace(cells[0].MustText())
		// Skip baris header atau baris kosong
		if date == "" || strings.EqualFold(date, "tanggal") || strings.EqualFold(date, "date") {
			continue
		}
		// Skip baris yang jelas bukan tanggal (misalnya footer)
		if len(date) < 6 {
			continue
		}

		mutations = append(mutations, Mutation{
			Date:        date,
			Description: strings.TrimSpace(cells[1].MustText()),
			Amount:      strings.TrimSpace(cells[2].MustText()),
			Type:        strings.TrimSpace(cells[3].MustText()),
			Balance:     strings.TrimSpace(cells[4].MustText()),
		})
	}

	return mutations
}

// deduplicate menghapus baris duplikat berdasarkan kombinasi Date+Description+Amount.
func deduplicate(mutations []Mutation) []Mutation {
	seen := make(map[string]bool)
	var result []Mutation
	for _, m := range mutations {
		key := m.Date + "|" + m.Description + "|" + m.Amount
		if !seen[key] {
			seen[key] = true
			result = append(result, m)
		}
	}
	return result
}

// DefaultDateRange mengembalikan range 7 hari terakhir dalam format "ddmmyyyy".
func DefaultDateRange() (start, end string) {
	now := time.Now()
	weekAgo := now.AddDate(0, 0, -7)
	format := func(t time.Time) string {
		return fmt.Sprintf("%02d%02d%04d", t.Day(), int(t.Month()), t.Year())
	}
	return format(weekAgo), format(now)
}

// GetEnvOrDefault membaca env variable, kembalikan fallback jika kosong.
func GetEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// FormatDateDisplay mengubah "ddmmyyyy" → "dd/mm/yyyy"
func FormatDateDisplay(d string) string {
	if len(d) != 8 {
		return d
	}
	return d[:2] + "/" + d[2:4] + "/" + d[4:]
}

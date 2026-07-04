package bot

import (
	"fmt"
	"strings"
)

// Mutation merepresentasikan satu baris transaksi mutasi rekening BCA.
type Mutation struct {
	Date        string
	Description string
	Amount      string
	Type        string // "DB" (debit) atau "CR" (credit)
	Balance     string
}

// PrintTable mencetak daftar mutasi ke stdout dalam format tabel ASCII.
func PrintTable(mutations []Mutation) {
	if len(mutations) == 0 {
		fmt.Println("⚠️  Tidak ada mutasi ditemukan untuk periode tersebut.")
		return
	}

	// Header
	sep := strings.Repeat("─", 110)
	fmt.Println(sep)
	fmt.Printf("%-12s  %-45s  %-18s  %-4s  %-18s\n",
		"TANGGAL", "KETERANGAN", "NOMINAL", "TYPE", "SALDO")
	fmt.Println(sep)

	for _, m := range mutations {
		keterangan := m.Description
		if len(keterangan) > 44 {
			keterangan = keterangan[:41] + "..."
		}

		typeLabel := m.Type
		switch strings.TrimSpace(strings.ToUpper(m.Type)) {
		case "DB":
			typeLabel = "🔴 DB"
		case "CR":
			typeLabel = "🟢 CR"
		}

		fmt.Printf("%-12s  %-45s  %-18s  %-6s  %-18s\n",
			m.Date, keterangan, m.Amount, typeLabel, m.Balance)
	}

	fmt.Println(sep)
	fmt.Printf("Total transaksi: %d\n", len(mutations))
}

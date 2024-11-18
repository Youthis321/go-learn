package main

import (
	"fmt"
	"time"
)

func main() {
	// ===== If Expression =====
	fmt.Println("=== If Expression ===")
	nilai := 75
	if nilai >= 80 {
		fmt.Println("Nilai Anda A")
	} else if nilai >= 70 {
		fmt.Println("Nilai Anda B")
	} else {
		fmt.Println("Nilai Anda C")
	}

	// ===== If dengan Short Statement =====
	fmt.Println("\n=== If dengan Short Statement ===")
	// Statement sebelum kondisi, dipisahkan dengan semicolon
	if skor := 85; skor >= 80 {
		fmt.Printf("Selamat! Nilai %d mendapatkan peringkat A\n", skor)
	} else {
		fmt.Printf("Nilai %d belum mencapai peringkat A\n", skor)
	}
	// Note: variabel skor hanya tersedia dalam blok if-else

	// ===== Switch Expression =====
	fmt.Println("\n=== Switch Expression ===")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari kerja")
		// fallthrough akan melanjutkan ke case berikutnya
		fallthrough
	case "Selasa":
		fmt.Println("Meeting mingguan")
	case "Sabtu", "Minggu": // multiple cases
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari biasa")
	}

	// ===== Switch tanpa Expression =====
	fmt.Println("\n=== Switch tanpa Expression ===")
	jam := 15
	switch {
	case jam < 12:
		fmt.Println("Selamat Pagi")
	case jam < 17:
		fmt.Println("Selamat Siang")
	default:
		fmt.Println("Selamat Malam")
	}

	// ===== Switch dengan Short Statement =====
	fmt.Println("\n=== Switch dengan Short Statement ===")
	switch waktu := time.Now(); {
	case waktu.Hour() < 12:
		fmt.Println("Sekarang masih pagi")
	case waktu.Hour() < 17:
		fmt.Println("Sekarang siang")
	default:
		fmt.Println("Sekarang malam")
	}

	// ===== For Loop Basic =====
	fmt.Println("\n=== For Loop Basic ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iterasi ke-%d\n", i)
	}

	// ===== For sebagai While =====
	fmt.Println("\n=== For sebagai While ===")
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// ===== For Range =====
	fmt.Println("\n=== For Range ===")
	fruits := []string{"apel", "jeruk", "mangga"}
	for index, fruit := range fruits {
		fmt.Printf("index: %d, buah: %s\n", index, fruit)
	}

	// ===== For dengan Break =====
	fmt.Println("\n=== For dengan Break ===")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at i =", i)
			break
		}
		fmt.Printf("Nilai i: %d\n", i)
	}

	// ===== For dengan Continue =====
	fmt.Println("\n=== For dengan Continue ===")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("Skipping i =", i)
			continue
		}
		fmt.Printf("Nilai i: %d\n", i)
	}

	// ===== Nested Loop dengan Label =====
	fmt.Println("\n=== Nested Loop dengan Label ===")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking outer loop")
				break outerLoop
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// ===== Contoh Praktis: Game Tebak Angka =====
	fmt.Println("\n=== Game Tebak Angka ===")
	targetAngka := 42
	maksPercobaan := 5

	fmt.Println("Selamat datang di Game Tebak Angka!")
	fmt.Printf("Tebak angka antara 1-100 (Kesempatan: %d kali)\n", maksPercobaan)

	for percobaan := 1; percobaan <= maksPercobaan; percobaan++ {
		var tebakan int
		fmt.Printf("\nPercobaan ke-%d\nMasukkan tebakan Anda: ", percobaan)
		// Dalam contoh ini, kita akan menggunakan angka hardcoded
		// Biasanya Anda akan menggunakan fmt.Scan(&tebakan)
		if percobaan == 1 {
			tebakan = 50
		} else if percobaan == 2 {
			tebakan = 45
		} else {
			tebakan = 42
		}
		
		fmt.Printf("Anda menebak: %d\n", tebakan)

		if tebakan == targetAngka {
			fmt.Printf("\nSelamat! Anda berhasil menebak angka %d\n", targetAngka)
			fmt.Printf("Anda berhasil dalam %d percobaan\n", percobaan)
			break
		} else if percobaan == maksPercobaan {
			fmt.Printf("\nMaaf, Anda kehabisan kesempatan. Angka yang benar adalah %d\n", targetAngka)
			continue
		} else if tebakan < targetAngka {
			fmt.Println("Tebakan terlalu kecil")
			continue
		} else {
			fmt.Println("Tebakan terlalu besar")
			continue
		}
	}
}

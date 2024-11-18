package main

import "fmt"

func main() {
	// Data login default
	// const defaultUsername = "user123"
	// const defaultPassword = "user000"

	// constant multiple
	const (
		defaultUsername = "user123"
		defaultPassword = "user000"
	)

	// Variabel untuk menyimpan input login
	var inputUsername, inputPassword string

	for {
		fmt.Println("\nAnda Mau Apa?")
		fmt.Println("1. Login")
		fmt.Println("2. Lihat Data Login")
		fmt.Println("3. Keluar")

		// Pilihan menu
		var pilihan int
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&pilihan)

		// Switch untuk pemilihan menu
		switch pilihan {
		case 1:
			// Login
			fmt.Print("Masukan Username: ")
			fmt.Scanln(&inputUsername)
			fmt.Print("Masukan Password: ")
			fmt.Scanln(&inputPassword)

			// Validasi login
			if inputUsername == defaultUsername && inputPassword == defaultPassword {
				fmt.Println("Login Berhasil")
			} else {
				fmt.Println("Login Gagal: Username atau Password salah")
			}

		case 2:
			// Lihat Data Login
			fmt.Println("Data Login:")
			fmt.Println("Username: ", defaultUsername)
			fmt.Println("Password: Tidak ditampilkan demi keamanan")

		case 3:
			// Keluar
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return

		default:
			fmt.Println("Pilihan tidak tersedia, coba lagi.")
		}
	}
}

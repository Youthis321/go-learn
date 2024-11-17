package main

import "fmt"

func main() {
	var (
		username = "user123"
		password = "user000"
	)

	for {
		fmt.Println("\n Anda Mau Apa? ")
		fmt.Println("1. Login")
		fmt.Println("2. Lihat Data Login")
		fmt.Println("3. Keluar")

		// Pilihan menu
		var pilihan int
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&pilihan)

		// Switch untuk pemilihan menu
		switch pilihan {
		case 1:
			// Login
			fmt.Print("Masukan Username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukan Password: ")
			fmt.Scanln(&password)

			// validasi login
			if username == "user123" && password == "user000" {
				fmt.Println("Login Berhasil")
				return
			} else if username != "user123" && password != "user000" {
				fmt.Println("Login Gagal")
				continue
			}

		case 2:
			// Lihat Data Login
			fmt.Println("Data Login")
			fmt.Println("Username: ", username)
			fmt.Println("Password: ", password)

		case 3:
			// Keluar
			fmt.Println("Terima Kasih")
			return

		default:
			fmt.Println("Pilihan Tidak Tersedia")
		}

	}
}

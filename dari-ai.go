package main

import "fmt"

func main() {
	// Array untuk menyimpan divisi
	var divisi = [3]string{"Divisi IT", "Divisi Marketing", "Divisi Pengorderan"}

	// Struct untuk menyimpan data registrasi
	type TodoRegister struct {
		Nama      string
		NIK       int
		Email     string
		Password  string
		Divisi    string
	}

	// Slice untuk menyimpan daftar todo register
	var todoRegisters []TodoRegister

	// Perulangan untuk menu todoRegisters
	for {
		fmt.Println("\nMenu Todo Register")
		fmt.Println("1. Register")
		fmt.Println("2. Lihat Data Register")
		fmt.Println("3. Keluar")

		// Pilihan menu
		var choice int
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&choice)

		// Switch untuk pemilihan menu
		switch choice {
		case 1:
			// Input data register
			var nama, email, password, rememberPassword string
			var nik int

			fmt.Print("Masukkan Nama Lengkap: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan NIK: ")
			fmt.Scanln(&nik)
			fmt.Print("Masukkan Email: ")
			fmt.Scanln(&email)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&password)
			fmt.Print("Ulangi Password: ")
			fmt.Scanln(&rememberPassword)

			// Validasi password
			if password != rememberPassword {
				fmt.Println("Password tidak cocok. Registrasi gagal.")
				continue
			}

			// Pemilihan divisi
			var divisiIndex int
			fmt.Println("\nPilih Divisi:")
			for i, div := range divisi {
				fmt.Printf("%d. %s\n", i+1, div)
			}
			fmt.Print("Masukan nomor divisi (1-3): ")
			fmt.Scanln(&divisiIndex)

			// Validasi divisi
			if divisiIndex < 1 || divisiIndex > len(divisi) {
				fmt.Println("Pilihan divisi tidak valid. Registrasi gagal.")
				continue
			}

			// Simpan data ke slice
			todoRegisters = append(todoRegisters, TodoRegister{
				Nama:     nama,
				NIK:      nik,
				Email:    email,
				Password: password,
				Divisi:   divisi[divisiIndex-1],
			})
			fmt.Println("Data register berhasil disimpan!")

		case 2:
			// Lihat data register
			if len(todoRegisters) == 0 {
				fmt.Println("Belum ada data register.")
			} else {
				fmt.Println("\nData Register:")
				for i, register := range todoRegisters {
					fmt.Printf("%d. Nama: %s, NIK: %d, Email: %s, Divisi: %s\n",
						i+1, register.Nama, register.NIK, register.Email, register.Divisi)
				}
			}

		case 3:
			// Keluar
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

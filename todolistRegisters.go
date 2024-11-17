package main

import "fmt"

func main() {
	// Array untuk menyimpan divisi
	var divisi = [3]string{"Divisi IT", "Divisi Marketing", "Divisi Pengorderan"}

	// slice untuk menyimpan daftar todo register
	// data berupa string dan integer
	// var todoRegisters []string []int ganti menggunakan struct
	// Masalah: Deklarasi tidak valid. Slice tidak bisa memiliki dua tipe data berbeda. Slice hanya boleh berupa satu tipe data (misalnya []string atau []int).
	type TodoRegister struct {
		Nama     string
		NIK      int
		Email    string
		Password string
		Divisi   string
	}

	// slice untuk menyimpan daftar todo register
	var todoRegisters []TodoRegister

	// perulangan untuk menu todoRegisters
	for {
		fmt.Println("\nMenu Todo Register")
		fmt.Println("1. Register")
		fmt.Println("2. Lihat Data Register")
		fmt.Println("3. Keluar")

		// Pilihan menu
		var choice int
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&choice)

		// switch untuk pemilihan menu
		switch choice {
		case 1:
			// input data register
			// fmt.Print("Masukkan Nama Lengkap: ") // conoth nama DR. Sutisna M.Kom
			// var nama string
			// fmt.Print("Masukkan NIK: ")
			// var nik int
			// fmt.Print("Masukkan Email: ")
			// var email string
			// fmt.Print("Masukkan Password: ")
			// var password string
			// fmt.Print("Masukkan Ulangi Password: ")
			// var remember_password string
			// fmt.Scanln(&nama, &nik, &email, &password, &remember_password)
			// Masalah: Scanln hanya menerima satu variabel dalam satu panggilan.

			// perbaikan
			// input data register
			var nama, email, password, rememberPassword string
			var nik int

			// inputannya
			fmt.Print("Masukkan Nama Lengkap: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan NIK: ")
			fmt.Scanln(&nik)
			fmt.Print("Masukkan Email: ")
			fmt.Scanln(&email)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&password)
			fmt.Print("Masukkan Ulangi Password: ")
			fmt.Scanln(&rememberPassword)
			// fmt.Scanln(&nama, &nik, &email, &password, &remember

			// validasi password
			if password != rememberPassword {
				fmt.Println("Password tidak sesuai, Registrasi gagal.")
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

			// validasi divisi
			if divisiIndex < 1 || divisiIndex > len(divisi) {
				fmt.Println("Pilihan divisi tidak valid, Register gagal.")
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
			fmt.Println("Data Register berhasil disimpan!")

		case 2:
			// lihat data register
			if len(todoRegisters) == 0 {
				fmt.Println("Belum ada data register")
			} else {
				fmt.Println("\nData Register:")
				for i, register := range todoRegisters {
					fmt.Printf("%d. Nama: %s, NIK: %d, Email: %s, Divisi: %s\n",
						i+1, register.Nama, register.NIK, register.Email, register.Divisi)
				}
			}

		case 3:
			// keluar
			fmt.Println("Terima kasih telah menggunakan aplikasi ini")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi")
		}
	}
}

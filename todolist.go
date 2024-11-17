package main

import (
	"fmt"
)

func main() {
	// Array untuk menyimpan prioritas tugas tetap (ukuran tetap, tidak bisa diubah)
	var priorities = [3]string{"Low", "Medium", "High"}

	// Slice untuk menyimpan daftar Todo (ukuran fleksibel)
	var todos []string

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambahkan Todo")
		fmt.Println("2. Lihat Todos")
		fmt.Println("3. Keluar")

		// Pilihan menu
		var choice int
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Tambahkan tugas baru
			var task string
			fmt.Print("Masukkan nama tugas: ")
			fmt.Scan(&task)

			var priorityIndex int
			fmt.Println("Pilih prioritas:")
			for i, priority := range priorities {
				fmt.Printf("%d. %s\n", i+1, priority)
			}
			fmt.Print("Masukkan nomor prioritas (1-3): ")
			fmt.Scan(&priorityIndex)

			// Validasi prioritas
			if priorityIndex < 1 || priorityIndex > len(priorities) {
				fmt.Println("Prioritas tidak valid!")
			} else {
				// Tambahkan tugas ke slice
				todos = append(todos, fmt.Sprintf("%s [%s]", task, priorities[priorityIndex-1]))
				fmt.Println("Tugas berhasil ditambahkan!")
			}

		case 2:
			// Lihat semua tugas
			if len(todos) == 0 {
				fmt.Println("Tidak ada tugas.")
			} else {
				fmt.Println("\nDaftar Todo:")
				for i, todo := range todos {
					fmt.Printf("%d. %s\n", i+1, todo)
				}
			}

		case 3:
			// Keluar dari aplikasi
			fmt.Println("Terima kasih telah menggunakan TodoList App!")
			return

		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}

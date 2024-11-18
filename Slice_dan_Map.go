package main

import (
	"fmt"
)

func main() {
	// Demonstrasi Slice
	fmt.Println("=== Demonstrasi Slice ===")
	
	// Membuat slice dengan make
	numbers := make([]int, 5)
	fmt.Printf("1. Slice kosong dengan panjang 5: %v\n", numbers)
	
	// Mengisi nilai ke slice
	for i := 0; i < 5; i++ {
		numbers[i] = (i + 1) * 10
	}
	fmt.Printf("2. Slice setelah diisi: %v\n", numbers)
	
	// Menambah elemen dengan append
	numbers = append(numbers, 60, 70)
	fmt.Printf("3. Slice setelah append: %v\n", numbers)
	
	// Slice dari slice yang ada
	subset := numbers[2:5] // mengambil elemen index 2,3,4
	fmt.Printf("4. Subset slice (index 2-4): %v\n", subset)
	
	// Mengubah nilai pada subset akan mempengaruhi slice asli
	subset[0] = 999
	fmt.Printf("5. Slice asli setelah mengubah subset: %v\n", numbers)
	
	// Demonstrasi Map
	fmt.Println("\n=== Demonstrasi Map ===")
	
	// Membuat map dengan make
	studentScores := make(map[string]int)
	
	// Menambah data ke map
	studentScores["Budi"] = 85
	studentScores["Ani"] = 90
	studentScores["Citra"] = 88
	fmt.Printf("1. Map setelah diisi: %v\n", studentScores)
	
	// Mengakses nilai dari map
	aniScore := studentScores["Ani"]
	fmt.Printf("2. Nilai Ani: %d\n", aniScore)
	
	// Mengecek keberadaan key dalam map
	score, exists := studentScores["Dedi"]
	if exists {
		fmt.Printf("3. Nilai Dedi: %d\n", score)
	} else {
		fmt.Println("3. Dedi tidak ada dalam daftar")
	}
	
	// Menghapus item dari map
	delete(studentScores, "Budi")
	fmt.Printf("4. Map setelah menghapus Budi: %v\n", studentScores)
	
	// Iterasi map menggunakan range
	fmt.Println("\n5. Daftar semua siswa dan nilai:")
	for name, score := range studentScores {
		fmt.Printf("   %s: %d\n", name, score)
	}
	
	// Contoh penggunaan slice dan map bersama
	fmt.Println("\n=== Kombinasi Slice dan Map ===")
	
	// Membuat slice of maps
	students := []map[string]interface{}{
		{"nama": "Eko", "nilai": 85, "hobi": []string{"membaca", "coding"}},
		{"nama": "Sari", "nilai": 90, "hobi": []string{"musik", "menari"}},
		{"nama": "Rudi", "nilai": 88, "hobi": []string{"olahraga", "gaming"}},
	}
	
	// Mengakses dan menampilkan data
	fmt.Println("Data Siswa:")
	for _, student := range students {
		nama := student["nama"].(string)
		nilai := student["nilai"].(int)
		hobi := student["hobi"].([]string)
		
		fmt.Printf("Nama: %s\n", nama)
		fmt.Printf("Nilai: %d\n", nilai)
		fmt.Printf("Hobi: %v\n", hobi)
		fmt.Println("---")
	}
}

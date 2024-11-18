Saya akan membantu Anda memahami penggunaan array slice dan map dalam Go dengan membuat sebuah contoh aplikasi sederhana.



Saya telah membuat contoh kode yang mendemonstrasikan penggunaan slice dan map dalam Go. Mari saya jelaskan bagian-bagian pentingnya:

1. Demonstrasi Slice:
- Membuat slice kosong dengan `make([]int, 5)`
- Mengisi nilai ke dalam slice menggunakan loop
- Menambah elemen dengan `append()`
- Membuat subset dari slice dengan slicing `numbers[2:5]`
- Menunjukkan bahwa perubahan pada subset mempengaruhi slice asli

2. Demonstrasi Map:
- Membuat map dengan `make(map[string]int)`
- Menambah dan mengakses data dalam map
- Mengecek keberadaan key dengan cara yang aman
- Menghapus item dengan `delete()`
- Iterasi map menggunakan `range`

3. Kombinasi Slice dan Map:
- Membuat slice yang berisi map
- Menggunakan interface{} untuk menyimpan berbagai tipe data
- Mengakses dan menampilkan data nested

Output program ini akan menunjukkan:
- Cara kerja slice: pembuatan, pengisian, append, dan slicing
- Operasi dasar map: pembuatan, pengisian, pengecekan, dan penghapusan
- Bagaimana mengkombinasikan slice dan map untuk data yang lebih kompleks

Beberapa poin penting untuk diingat:
1. Slice adalah reference type, artinya ketika Anda membuat subset atau meneruskan slice ke fungsi, mereka berbagi data yang sama
2. Map juga reference type dan tidak perlu dipassing sebagai pointer untuk dimodifikasi
3. Gunakan `make()` untuk membuat slice dan map untuk menghindari nil pointer
4. Selalu gunakan comma-ok idiom (`value, exists := map[key]`) ketika mengakses map untuk menghindari panic

Apakah ada bagian tertentu yang ingin Anda pahami lebih dalam?

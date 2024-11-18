Saya akan membantu Anda memahami berbagai jenis kontrol flow dalam Go dengan membuat contoh aplikasi yang mendemonstrasikan semuanya.

Mari saya jelaskan setiap bagian dari kode tersebut:

1. **If Expression**
```go
if nilai >= 80 {
    fmt.Println("Nilai Anda A")
} else if nilai >= 70 {
    fmt.Println("Nilai Anda B")
} else {
    fmt.Println("Nilai Anda C")
}
```
- Ini adalah bentuk dasar if-else
- Kondisi tidak perlu menggunakan tanda kurung ()
- Kurung kurawal {} wajib digunakan

2. **If dengan Short Statement**
```go
if skor := 85; skor >= 80 {
    fmt.Printf("Selamat! Nilai %d mendapatkan peringkat A\n", skor)
}
```
- Memungkinkan deklarasi variabel sebelum kondisi
- Variabel hanya berlaku dalam scope if-else
- Berguna untuk membatasi scope variabel

3. **Switch Expression**
```go
switch hari {
case "Senin":
    fmt.Println("Hari kerja")
    fallthrough
case "Selasa":
    fmt.Println("Meeting mingguan")
case "Sabtu", "Minggu":
    fmt.Println("Hari libur")
default:
    fmt.Println("Hari biasa")
}
```
- Tidak perlu break (otomatis break)
- Bisa menggunakan fallthrough untuk lanjut ke case berikutnya
- Bisa multiple values dalam satu case

4. **Switch dengan Short Statement**
```go
switch waktu := time.Now(); {
case waktu.Hour() < 12:
    fmt.Println("Sekarang masih pagi")
```
- Bisa mendeklarasikan variabel sebelum switch
- Bisa tanpa expression (seperti if-else)

5. **For Loop**
```go
// Basic for
for i := 0; i < 5; i++ {
    fmt.Printf("Iterasi ke-%d\n", i)
}

// For sebagai while
for counter < 3 {
    fmt.Printf("Counter: %d\n", counter)
    counter++
}

// For range
for index, fruit := range fruits {
    fmt.Printf("index: %d, buah: %s\n", index, fruit)
}
```
- Satu-satunya struktur loop di Go
- Bisa digunakan seperti while
- For range untuk iterasi slice, array, map, dll

6. **Break dan Continue**
```go
// Break
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
}

// Continue
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
}
```
- Break untuk menghentikan loop
- Continue untuk melanjutkan ke iterasi berikutnya

7. **Label dengan Break**
```go
outerLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break outerLoop
            }
        }
    }
```
- Label berguna untuk break dari nested loop
- Membuat kode lebih jelas dalam kasus nested loop

8. **Contoh Praktis**
- Implementasi game tebak angka yang menggunakan semua konsep di atas
- Menunjukkan penggunaan break dan continue dalam konteks nyata
- Menggunakan if-else untuk logika permainan
- Menggunakan for loop untuk mengontrol jumlah percobaan

Poin-poin penting:
1. Go tidak memiliki operator ternary (?:)
2. Switch di Go lebih fleksibel dibanding bahasa lain
3. For adalah satu-satunya loop construct di Go
4. Break dan continue bisa digunakan dengan label
5. Short statement dalam if dan switch membuat kode lebih rapi

Apakah ada bagian tertentu yang ingin Anda pahami lebih dalam?

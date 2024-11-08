// CASE 1: Troubleshooting
// identifikasi dan koreksi kesalahan pada program di bawah

// EXPECTED OUTPUT:
// [0 1 2 4 5 6]
// 6
// Umur Imanuel tahun ini: 21
// Siswa baru, Siska berumur: 19
// [11 13 17 19 23 29]

package main

import "fmt"

func main() {
	// Part 1: Mengapa terdapat error pada baris-baris kode berikut?
	// Perbaiki baris-baris berikut agar program dapat dijalankan
	myArr := []uint{0, 1, 2, 4, 5, 6} // Mengganti dari array[5] menjadi slice[]

	fmt.Println(myArr)
	fmt.Println(myArr[len(myArr)-1]) // Diubah untuk mengakses elemen terakhir dengan benar
	// -- End of Part 1 --

	// Part 2: Output program tidak sesuai dengan EXPECETED OUTPUT
	// Jangan lakukan perubahan langsung pada stdAge di awal.
	// Coba gunakan key untuk melakukan perubahan dan menambahkan
	// item ke map.
	stdAge := map[string]uint{
		"Agus":     19,
		"Irfan":    20,
		"Johannes": 19,
		"Imanuel":  20,
	}

	stdAge["Imanuel"] = 21 // Memperbarui umur Imanuel ke 21
	stdAge["Siska"] = 19   // Menambahkan entry baru untuk Siska

	fmt.Println("Umur Imanuel tahun ini:", stdAge["Imanuel"])
	fmt.Println("Siswa baru, Siska berumur:", stdAge["Siska"])
	// -- End of Part 2 --

	// Part 3: Baris-baris kode berikut seharusnya akan menampilkan
	// bilangan prima antara 10 hingga 30.
	// Identifikasi letak kesalahan logika pada program dan
	// koreksi agar menampilkan output yang sesuai.
	var primes []int
	for i := 10; i < 30; i++ {
		if i == 2 || i == 3 || i == 5 || i == 7 || (i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0) {
			primes = append(primes, i) // Perbaikan untuk hanya menambahkan bilangan prima
		}
	}

	fmt.Println(primes)
}

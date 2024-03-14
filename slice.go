package main

import "fmt"

func main(){


	// Membuat Slice
	// Karena slice adalah referensi dari array, maka kita definisikan dulu arraynya (sebenarnya bisa juga tidak)
	months := [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember"}

	fmt.Println("Array Asli : ", months)

	// Kita 'pasangkan' slice di array tersebut agar bisa lebih fleksibel
	// format penulisan : nama variabel-:=-referensi array-urutan index pada array dari ke sampai
	monthsSlice := months[3:9]
	fmt.Println("Slice dari index 3 - 8 : ",monthsSlice)

	// Cap atau Capacity adalah kapasitas data yang tersedia dari index pertama slice sampai bawah
	// Dalam hal ini, monthsSlice membuat slice dari array month dimulai index ke 3 (April)
	// Maka dari April - Desember (index 3 - 11) ada 9 data (kapasitas slice)
	fmt.Println(cap(monthsSlice))

	// Kita bisa membuat slice tanpa harus mendefinisikan array
	// Ini artinya kita Go membuat array secara otomatis dan mengaitkan slice dengan array tersebut
	devices := []string{"Laptop", "Smartphone", "Printer"}

	fmt.Println(devices)
}
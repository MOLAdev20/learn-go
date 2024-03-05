package main

import "fmt"

func main() {

	// Mendeklarasikan variabel disertai inisialisasi tipe data
	var name string
	name = "Sabiilul Hikam"
	fmt.Println(name)

	// Mendeklarasikan variabel tanpa inisialisai tipe data
	var alamat = "Tangerang"
	fmt.Println(alamat)

	// Bisa juga sembari diinisialisasi tipe data.
	// Cara ini berguna untuk tipe data integer
	var umur int8 = 20
	fmt.Println(umur)

	// Mendeklarasikan variabel dengan cara singkat (tanpa keyword var)
	// Kelemahan: Kita tidak bisa mengatur tipe datanya.
	// Variabel pendapatan otomatis diset ke int64 (sesuai jenis arsitektur device)
	pendapatan := 4000000
	fmt.Println(pendapatan)

	// Mendeklarasikan variabel secara bersamaan
	// Kita bisa menginisialisasi tipe data juga pada teknik ini
	var (
		hobi   string = "Koding"
		kuliah bool   = true
	)

	fmt.Println("Hobi =", hobi)
	fmt.Println("Apakah ", name, " kuliah?", kuliah)
}

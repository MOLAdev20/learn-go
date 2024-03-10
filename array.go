package main

import "fmt"

func main(){

	// Membuat Array dengan berbagai macam pola

	// CARA 1
	// nama variabel-kemungkinan jumlah data-tipe data
	var karyawan [4]string

	// Inisialisasi isiannya
	karyawan[0] = "Erik"
	karyawan[1] = "Jason"
	karyawan[2] = "Budi"
	karyawan[3] = "Kumar"

	fmt.Println("Daftar Karyawan :",karyawan)

	// CARA 2
	// nama variabel-kemungkinan jumlah data-tipe data-{data/anggota array}
	var prima = [3]int8{2, 3, 5}
	fmt.Println("Bilangan Prima :",prima)

	// CARA 3
	// tanpa mendefinisikan jumlah data
	var infinity = [...]int8{1,2,3,4,5}

	fmt.Println("Nilai :",infinity)

	fmt.Println("Panjang elemen array infinity : ",len(infinity))
}
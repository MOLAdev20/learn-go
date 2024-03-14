package main

import "fmt"

func main(){
	
	// Membuat if di Go sama seperti membuat if di bahasa lain
	// Bedanya, go tidak memerlukan dalam/tutup kurung di perkondisiannya

	username := "admin"

	if username == "admin" {
		fmt.Println("Username benar")
	}

	// Kita bisa menginisialisasi variabel langsung sebaris dengan if
	// Dengan catatan kita hanya bisa mengakses variabel hanya dalam lingkup if dan kondisinya saja
	if password := "4dmin"; password == "4dmin"{
		fmt.Println("Silahkan Masuk")
	}

	if x, y := 12, 18; x < y {
		fmt.Println("Nilai X adalah ",x," X lebih kecil dari ", y)
	}

	// If else dan if else if
	maxUmur := 20
	umur := 25

	if maxUmur > umur {
		fmt.Println("Bocil dilarang masuk")
	}else{
		fmt.Println("Kamu dewasa")
	}

	nilai := 81

	if nilai <= 60 {
		fmt.Println("Nilai Kamu C")
	}else if nilai <= 80 {
		fmt.Println("Nilai Kamu B")
	}else if nilai <= 90 {
		fmt.Println("Nilai Kamu A-")
	}else{
		fmt.Println("Nilai Kamu A+")
	}
}
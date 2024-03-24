package main

import "fmt"

func main(){
	// Function utama adalah main (kode program dieksekusi pertama kali di sini)
	// Setiap function yang kita buat, kalo mau dieksekusi, wajib hukumnya diletakan di main func ini
	Sapa()
	SayHi("Sabil")
	var todayQuotes string = GetKataKata("Never Stop Learning Because Tech Never Stop Growing")
	fmt.Println(todayQuotes)
}

// Function biasa
func Sapa(){
	fmt.Println("Hello World")
}

// Function yang menerima parameter
// Kita wajib mendefinisikan tipe data si paramaternya
func SayHi(name string){
	fmt.Println("Hai "+name)
}


// Function yang menjalankan proses sekaligus mengembalikan sebuah nilai
func GetKataKata(quotes string)string{
	return "Kata Kata Hari Ini: "+quotes
}
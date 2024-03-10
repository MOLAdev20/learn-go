package main

import "fmt"

func main(){
	
	// Konversi tipe data dari integer kapasitas rendah ke kapasitas besar
	var nilai int32 = 250000
	var nilaiMeningkat int64 = int64(nilai)
	fmt.Println(nilaiMeningkat)

	// Konversi tipe data integer ke string (Mendapatkan satu huruf tertentu dari sebuah kata)
	var nama = "Sabilul Hikam"
	var getChar = nama[2] // Mendapatkan huruf 'b' dari nama Sabilul
	fmt.Println(getChar) // Hasilnya pasti 98 (integer)

	// Kenapa demikian? karena kita hanya mendapatkan nilai byte nya saja.
	// Sedangkan byte itu adalah integer (int8)
	
	// Konversi byte ke string
	var byteToChar = string(getChar)
	fmt.Println(byteToChar)
}
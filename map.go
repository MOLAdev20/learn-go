package main

import "fmt"

func main(){
	
	// Membuat MAP dengan cara panjang
	var profile map[string]string = map[string]string{
		"nama" : "Sabilul Hikam",
		"umur" : "20 Tahun",
		"status" : "Belajar Go-Lang",
	}

	fmt.Println("Deklarasi Panjang : ",profile)

	// Cara simple
	profile2 := map[string]string{
		"nama" : "Budi Wicaksono",
		"umur" : "25 Tahun",
		"status" : "Belajar C++",
	}

	fmt.Println("Deklarasi Singkat :",profile2)

	// Mengakses data pada map
	fmt.Println(profile["nama"])

	// Menambahkan data pada map
	profile["pekerjaan"] = "Software Engineer"
	fmt.Println(profile)
}
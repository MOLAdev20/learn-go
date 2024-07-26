package main

import (
	"coba-go/helper"
	"fmt"
	// "time"
)

func main(){

	var tanggalLahir string = "24-05-2003"
	fmt.Println("Kamu lahir ditahun",tanggalLahir,". Usiamu saat ini :",helper.SumAge("24-05-2003"),"Tahun")

	fmt.Println(helper.Jumlahkan(12, 12))

	// today := time.Now()

	// fmt.Println(today.Year())
}
package helper

import (
	"time"
	"strings"
	"strconv"
)

func SayHello(words string) string {
	return "Hai, "+words
}

func convertString(word string) int {
	result, err := strconv.Atoi(word)

	if(err != nil){
		panic("error, bukan string")
		return 0;
	}

	return result
}

func SumAge(tglLahir string) int {

	parseTglLahir := strings.Split(tglLahir, "-")

	tahunLahir := convertString(parseTglLahir[2])
	// bulanLahir := convertString(parseTglLahir[1])
	// hariLahir := convertString(parseTglLahir[0])
	
	today := time.Now()

	usia := today.Year() - tahunLahir

	return usia
}

func Jumlahkan(a int, b int) int{
	hasil := ProsesJumlah(a, b)

	return hasil
}
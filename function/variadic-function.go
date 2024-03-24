package main

import "fmt"

// Variadic function adalah sebuah function yang seakan bisa menerima parameter tak hingga (infinity)
// Tapi data parameter tersebut harus sejenis

func hitungSemua(nomor ...int) int{

	total := 0;
	for i := 0; i < len(nomor); i++{
		total = nomor[i] + total
	}

	return total
}

func main(){
	fmt.Println(hitungSemua(10, 10, 10))
}
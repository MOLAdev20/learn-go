package main

import "fmt"

type Filter =  func(string) string

func SayHello(nama string, filter Filter) {
	fmt.Println("Hello ", filter(nama))
}

func BadWordFilter(karakter string) string {
	if(karakter == "Anjing"){
		return "..."
	}else{
		return karakter
	}
}

func main(){
	SayHello("Anjing", BadWordFilter)
}
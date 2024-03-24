package main

import "fmt"

func SayGoodBye(name string) string {
    return name;
}

func main(){

    var name string = SayGoodBye("Budi");

    fmt.Println("Selamat Tinggal ", name)

}
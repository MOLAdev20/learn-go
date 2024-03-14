package main

import "fmt"

func main(){

	animals := [5]string{"Kucing", "Sapi", "Kelinci", "Kuda", "Kerbau"}

	animSlice := animals[2:]

	animSlice[0] = "Burung"

	fmt.Println(animals)
	fmt.Println(animSlice)

}